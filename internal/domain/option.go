package domain

import (
	"context"
	"time"
)

type Option struct {
	ID         uint       `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitzero"`
	Content    string     `json:"content"`
	QuestionID uint       `json:"question_id"`
}

type OptionRepository interface {
	Create(ctx context.Context, option *Option) error
	FindByQuestionID(ctx context.Context, questionID uint) ([]Option, error)
	FindAll(ctx context.Context) ([]Option, error)
}
