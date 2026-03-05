package http_controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mobqom/questions/internal/dto"
	"github.com/mobqom/questions/internal/usecase"
)

type QuestionController struct {
	uc usecase.QuestionUseCase
}

func NewQuestionController(uc usecase.QuestionUseCase) *QuestionController {
	return &QuestionController{uc: uc}
}

// FindAll godoc
// @Summary Получить все вопросы
// @Description Возвращает список всех вопросов из базы данных
// @Tags questions
// @Produce json
// @Success 200 {array} domain.Question
// @Router /questions/find-all [get]
func (c *QuestionController) FindAll(w http.ResponseWriter, r *http.Request) {
	questions, err := c.uc.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
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
// @Router /questions/add-question [post]
func (c *QuestionController) AddQuestion(w http.ResponseWriter, r *http.Request) {
	var question dto.AddQuestionDto
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if question.Content == "" || question.Game == "" {
		http.Error(w, "Content and Game are required", http.StatusBadRequest)
		return
	}

	q, err := c.uc.AddQuestion(r.Context(), question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(q); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}
