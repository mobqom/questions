package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/mobqom/questions/config"
	"github.com/mobqom/questions/internal/controller/http"
	"github.com/mobqom/questions/internal/db"
	"github.com/mobqom/questions/internal/repository"
	"github.com/mobqom/questions/internal/usecase"
	"github.com/mobqom/questions/migrations"

	_ "github.com/mobqom/questions/docs"
)

func Run(cfg *config.AppConfig) {
	dbConn, err := db.Connection(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	migrations.Init(dbConn)

	validate := validator.New()

	// Initialize layers
	repo := repository.NewQuestionRepository(dbConn)
	uc := usecase.NewQuestionUseCase(repo)
	questionCtrl := httpController.NewQuestionController(uc, validate)

	optRepo := repository.NewOptionsRepository(dbConn)
	optUc := usecase.NewOptionsUseCase(optRepo)
	optionsCtrl := httpController.NewOptionsController(optUc, validate)

	r := chi.NewRouter()

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		httpController.RegisterRoutes(r, questionCtrl, optionsCtrl)
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
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
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	log.Printf("Server starting on %s", addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server failed to start: %v", err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	log.Println("Server gracefully stopped")
}
