package domain

import (
	"context"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Content string    `gorm:"not null" json:"content"`
	Game    string    `gorm:"type:varchar(255);not null;index" json:"game"`
	Options []Options `gorm:"foreignKey:QuestionID" json:"options,omitzero"`
}

type QuestionRepository interface {
	FindAll(ctx context.Context) ([]Question, error)
	Create(ctx context.Context, question *Question) error
	FindRandomQuestion(ctx context.Context, gameId string) (*Question, error)
	FindByGameId(ctx context.Context, gameId string) ([]Question, error)
}
