package domain

import (
	"context"

	"gorm.io/gorm"
)

type Options struct {
	gorm.Model
	Content    string   `gorm:"not null" json:"content"`
	QuestionID uint     `gorm:"not null" json:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID" json:"-"`
}

type OptionsRepository interface {
	Create(ctx context.Context, option *Options) error
	FindByQuestionID(ctx context.Context, questionID uint) ([]Options, error)
	FindAll(ctx context.Context) ([]Options, error)
}
