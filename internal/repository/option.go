package repository

import (
	"context"

	"github.com/mobqom/questions/internal/domain"
	"gorm.io/gorm"
)

type optionsRepository struct {
	db *gorm.DB
}

func NewOptionsRepository(db *gorm.DB) domain.OptionsRepository {
	return &optionsRepository{db: db}
}

func (r *optionsRepository) Create(ctx context.Context, option *domain.Options) error {
	return r.db.WithContext(ctx).Create(option).Error
}

func (r *optionsRepository) FindByQuestionID(ctx context.Context, questionID uint) ([]domain.Options, error) {
	var options []domain.Options
	if err := r.db.WithContext(ctx).Where("question_id = ?", questionID).Find(&options).Error; err != nil {
		return nil, err
	}
	return options, nil
}

func (r *optionsRepository) FindAll(ctx context.Context) ([]domain.Options, error) {
	var options []domain.Options
	if err := r.db.WithContext(ctx).Find(&options).Error; err != nil {
		return nil, err
	}
	return options, nil
}
