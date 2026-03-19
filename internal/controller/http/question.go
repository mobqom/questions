package httpController

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mobqom/questions/internal/dto"
	"github.com/mobqom/questions/internal/usecase"
)

type QuestionController struct {
	uc       usecase.QuestionUseCase
	validate *validator.Validate
}

func NewQuestionController(uc usecase.QuestionUseCase, validate *validator.Validate) *QuestionController {
	return &QuestionController{uc: uc, validate: validate}
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

	if err := c.validate.Struct(question); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

// FindRandomQuestionListByGameId godoc
// @Summary Получить случайный вопрос
// @Description Возвращает случайный вопрос из базы данных
// @Tags questions
// @Produce json
// @Param gameId query string true "ID игры"
// @Success 200 {object} domain.Question
// @Router /questions/random [get]
func (c *QuestionController) FindRandomQuestionListByGameId(w http.ResponseWriter, r *http.Request) {
	params := dto.QuestionQueryParams{
		GameID: r.URL.Query().Get("gameId"),
		Type:   r.URL.Query().Get("type"),
		Count:  r.URL.Query().Get("count"),
	}

	if err := c.validate.Struct(params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Если в UseCase нужен именно int, конвертируем после валидации
	count, _ := strconv.Atoi(params.Count)
	_ = count // Использование в зависимости от логики

	// В данном случае FindRandomQuestion принимает только gameId
	q, err := c.uc.FindRandomQuestionListByGameId(r.Context(), params.GameID, params.Type, count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

// FindByGameId godoc
// @Summary Получить вопросы по ID игры
// @Description Возвращает список всех вопросов для конкретной игры
// @Tags questions
// @Produce json
// @Param gameId query string true "ID игры"
// @Success 200 {array} domain.Question
// @Router /questions/find-by-game [get]
func (c *QuestionController) FindByGameId(w http.ResponseWriter, r *http.Request) {
	params := dto.QuestionQueryParams{
		GameID: r.URL.Query().Get("gameId"),
	}

	if err := c.validate.Struct(params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	questions, err := c.uc.FindByGameId(r.Context(), params.GameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}
