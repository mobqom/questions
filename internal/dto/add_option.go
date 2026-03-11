package dto

type AddOptionDto struct {
	Content    string `json:"content" validate:"required"`
	QuestionID uint   `json:"question_id" validate:"required"`
}
