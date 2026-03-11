package domain

import (
	"context"
	"time"
)

type Question struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitzero"`
	Content   string     `json:"content"`
	Game      string     `json:"game"`
	Options   []Option   `json:"options,omitzero"`
}

type QuestionRepository interface {
	FindAll(ctx context.Context) ([]Question, error)
	Create(ctx context.Context, question *Question) error
	FindRandomQuestion(ctx context.Context, gameId string) (*Question, error)
	FindByGameId(ctx context.Context, gameId string) ([]Question, error)
}
