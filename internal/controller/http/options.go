package http_controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mobqom/questions/internal/dto"
	"github.com/mobqom/questions/internal/usecase"
)

type OptionsController struct {
	uc usecase.OptionsUseCase
}

func NewOptionsController(uc usecase.OptionsUseCase) *OptionsController {
	return &OptionsController{uc: uc}
}

// FindAll godoc
// @Summary Получить все варианты ответов
// @Description Возвращает список всех вариантов ответов из базы данных
// @Tags options
// @Produce json
// @Success 200 {array} domain.Options
// @Router /options/find-all [get]
func (c *OptionsController) FindAll(w http.ResponseWriter, r *http.Request) {
	options, err := c.uc.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

// FindByQuestionID godoc
// @Summary Получить варианты ответов для вопроса
// @Description Возвращает список всех вариантов ответов для конкретного вопроса
// @Tags options
// @Produce json
// @Param questionId path int true "ID вопроса"
// @Success 200 {array} domain.Options
// @Router /options/{questionId} [get]
func (c *OptionsController) FindByQuestionID(w http.ResponseWriter, r *http.Request) {
	questionIDStr := chi.URLParam(r, "questionId")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	options, err := c.uc.FindByQuestionID(r.Context(), uint(questionID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

// AddOption godoc
// @Summary Добавить новый вариант ответа
// @Description Создает новый вариант ответа для вопроса
// @Tags options
// @Accept json
// @Produce json
// @Param option body dto.AddOptionDto true "Объект варианта ответа"
// @Success 201 {object} domain.Options
// @Failure 400 {string} string "Invalid request body"
// @Router /options/add-option [post]
func (c *OptionsController) AddOption(w http.ResponseWriter, r *http.Request) {
	var option dto.AddOptionDto
	if err := json.NewDecoder(r.Body).Decode(&option); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if option.Content == "" || option.QuestionID == 0 {
		http.Error(w, "Content and QuestionID are required", http.StatusBadRequest)
		return
	}

	o, err := c.uc.AddOption(r.Context(), option)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(o); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}
