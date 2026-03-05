package repository

import (
	"context"

	"github.com/mobqom/questions/internal/domain"
	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) domain.QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) FindAll(ctx context.Context) ([]domain.Question, error) {
	var questions []domain.Question
	if err := r.db.WithContext(ctx).Preload("Options").Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *questionRepository) Create(ctx context.Context, question *domain.Question) error {
	return r.db.WithContext(ctx).Create(question).Error
}

func (r *questionRepository) FindRandomQuestion(ctx context.Context, gameId uint) (*domain.Question, error) {
	var question domain.Question
	if err := r.db.WithContext(ctx).Preload("Options").Where("game_id = ?", gameId).Order("RANDOM()").First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *questionRepository) FindByGameId(ctx context.Context, gameId uint) ([]domain.Question, error) {
	var questions []domain.Question
	if err := r.db.WithContext(ctx).Preload("Options").Where("game_id = ?", gameId).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
