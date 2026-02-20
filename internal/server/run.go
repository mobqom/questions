package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run() {
	r := chi.NewRouter()

	http.ListenAndServe(":8081", r)
}
