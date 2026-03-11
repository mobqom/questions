package repository

import (
	"context"
	"time"

	"github.com/mobqom/questions/internal/domain"
	"gorm.io/gorm"
)

type OptionModel struct {
	gorm.Model
	Content    string `gorm:"not null"`
	QuestionID uint   `gorm:"not null"`
}

func (OptionModel) TableName() string {
	return "options"
}

type optionsRepository struct {
	db *gorm.DB
}

func NewOptionsRepository(db *gorm.DB) domain.OptionRepository {
	return &optionsRepository{db: db}
}

func (r *optionsRepository) Create(ctx context.Context, option *domain.Option) error {
	model := OptionModel{
		Content:    option.Content,
		QuestionID: option.QuestionID,
	}
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	option.ID = model.ID
	option.CreatedAt = model.CreatedAt
	option.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *optionsRepository) FindByQuestionID(ctx context.Context, questionID uint) ([]domain.Option, error) {
	var models []OptionModel
	if err := r.db.WithContext(ctx).Where("question_id = ?", questionID).Find(&models).Error; err != nil {
		return nil, err
	}
	options := make([]domain.Option, len(models))
	for i, m := range models {
		options[i] = m.ToDomain()
	}
	return options, nil
}

func (r *optionsRepository) FindAll(ctx context.Context) ([]domain.Option, error) {
	var models []OptionModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	options := make([]domain.Option, len(models))
	for i, m := range models {
		options[i] = m.ToDomain()
	}
	return options, nil
}

func (m OptionModel) ToDomain() domain.Option {
	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	return domain.Option{
		ID:         m.ID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  deletedAt,
		Content:    m.Content,
		QuestionID: m.QuestionID,
	}
}
