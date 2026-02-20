package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// FindAL godoc
// @Summary Проверка работоспособности
// @Description Возвращает простой текст, чтобы проверить, что сервис отвечает
// @Tags health
// @Produce plain
// @Success 200 {string} string "Hello, World!"
// @Router /api/v1/find-all [get]
func HttpController() chi.Router {
	r := chi.NewRouter()
	r.Get("/find-all", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	return r
}
