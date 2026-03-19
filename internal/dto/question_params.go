package dto

type QuestionQueryParams struct {
	GameID string `validate:"required"`
	Type   string `validate:"required"`
	Count  string `validate:"required,numeric,min=1,max=3"`
}
