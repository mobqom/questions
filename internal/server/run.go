package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/mobqom/questions/config"
	grpcController "github.com/mobqom/questions/internal/controller/grpc"
	httpController "github.com/mobqom/questions/internal/controller/http"
	"github.com/mobqom/questions/internal/db"
	"github.com/mobqom/questions/internal/repository"
	"github.com/mobqom/questions/internal/usecase"
	"github.com/mobqom/questions/migrations"
	optionsv1 "github.com/mobqom/questions/proto/v1/option"
	questionv1 "github.com/mobqom/questions/proto/v1/question"
	"google.golang.org/grpc/reflection"

	_ "github.com/mobqom/questions/docs"
)

func Run(cfg *config.AppConfig) {
	var dbConn *gorm.DB
	var err error

	// Retry connection to database
	for i := range 10 {
		dbConn, err = db.Connection(cfg)
		if err == nil {
			break
		}
		log.Printf("failed to connect to database (attempt %d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("could not connect to database after 10 attempts: %v", err)
	}
	migrations.Init(dbConn)

	validate := validator.New()

	// Initialize layers
	repo := repository.NewQuestionRepository(dbConn)
	uc := usecase.NewQuestionUseCase(repo)
	questionCtrl := httpController.NewQuestionController(uc, validate)

	repoOpt := repository.NewOptionsRepository(dbConn)
	ucOpt := usecase.NewOptionsUseCase(repoOpt)
	optionsCtrl := httpController.NewOptionsController(ucOpt, validate)

	r := chi.NewRouter()

	// Health checks
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	r.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if dbConn != nil {
			sqlDB, err := dbConn.DB()
			if err == nil {
				if err := sqlDB.Ping(); err == nil {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("ok"))
					return
				}
			}
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})

	// API routes
	r.Route("/api/questions", func(r chi.Router) {
		httpController.RegisterRoutes(r, questionCtrl, optionsCtrl)
		// Swagger UI
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("/api/questions/swagger/doc.json"),
		))
	})

	// gRPC Server
	grpcSrv := grpc.NewServer()
	questionv1.RegisterQuestionServiceServer(grpcSrv, grpcController.NewQuestionServer(uc))
	optionsv1.RegisterOptionsServiceServer(grpcSrv, grpcController.NewOptionsServer(ucOpt))
	reflection.Register(grpcSrv)

	addr := fmt.Sprintf(":%s", cfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	grpcAddr := fmt.Sprintf(":%s", cfg.GrpcPort)
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		grpcSrv.GracefulStop()
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	log.Printf("HTTP Server starting on %s", addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server failed to start: %v", err)
		}
	}()

	log.Printf("gRPC Server starting on %s", grpcAddr)
	go func() {
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed to start: %v", err)
		}
	}()

	// Wait for server context to be stopped
	<-serverCtx.Done()
	log.Println("Server gracefully stopped")
}
