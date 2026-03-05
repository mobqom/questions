package dto

type AddQuestionDto struct {
	Content string `json:"content" validate:"required"`
	Game    string `json:"game" validate:"required"`
}
