package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/mobqom/questions/config"
	"github.com/mobqom/questions/internal/controller"
	"github.com/mobqom/questions/internal/db"
	"github.com/mobqom/questions/migrations"

	_ "github.com/mobqom/questions/docs"
)

func Run(cfg *config.AppConfig) {
	r := chi.NewRouter()
	dbConn, err := db.Connection(cfg)
	if err != nil {
		panic(err)
	}
	migrations.Init(dbConn)

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// API routes
	r.Mount("/api/v1", controller.HttpController())

	http.ListenAndServe(":8081", r)
}
