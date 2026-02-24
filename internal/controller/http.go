package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mobqom/questions/internal/dto"
	"github.com/mobqom/questions/internal/usecase"
	"gorm.io/gorm"
)

// FindAll godoc
// @Summary Получить все вопросы
// @Description Возвращает список всех вопросов из базы данных
// @Tags questions
// @Produce json
// @Success 200 {array} domain.Question
// @Router /find-all [get]
func FindAll(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questions := usecase.GetQuestions(db)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
	}
}

// AddQuestion godoc
// @Summary Добавить новый вопрос
// @Description Создает новый вопрос в базе данных
// @Tags questions
// @Accept json
// @Produce json
// @Param question body dto.AddQuestionDto true "Объект вопроса"
// @Success 201 {object} domain.Question
// @Failure 400 {string} string "Invalid request body"
// @Router /add-question [post]
func AddQuestion(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var question dto.AddQuestionDto
		if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if question.Content == "" || question.Game == "" {
			http.Error(w, "Content and Game are required", http.StatusBadRequest)
			return
		}

		q := usecase.AddQuestion(db, question)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(q); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error encoding JSON: %v", err)
		}
	}
}

func HttpController(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	r.Get("/find-all", FindAll(db))
	r.Post("/add-question", AddQuestion(db))
	return r
}
