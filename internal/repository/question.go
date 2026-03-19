package repository

import (
	"context"
	"time"

	"github.com/mobqom/questions/internal/domain"
	"gorm.io/gorm"
)

type QuestionModel struct {
	gorm.Model
	Content string        `gorm:"not null"`
	Game    string        `gorm:"type:varchar(255);not null;index"`
	Type    string        `gorm:"type:varchar(255);not null;index;default:'simple'"`
	Options []OptionModel `gorm:"foreignKey:QuestionID"`
}

func (QuestionModel) TableName() string {
	return "questions"
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) domain.QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) FindAll(ctx context.Context) ([]domain.Question, error) {
	var models []QuestionModel
	if err := r.db.WithContext(ctx).Preload("Options").Find(&models).Error; err != nil {
		return nil, err
	}

	questions := make([]domain.Question, len(models))
	for i, m := range models {
		questions[i] = m.ToDomain()
	}
	return questions, nil
}

func (r *questionRepository) Create(ctx context.Context, question *domain.Question) error {
	model := QuestionModel{
		Content: question.Content,
		Game:    question.Game,
	}
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	question.ID = model.ID
	question.CreatedAt = model.CreatedAt
	question.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *questionRepository) FindRandomQuestionListByGameId(ctx context.Context, gameId string, qType string, count int) ([]domain.Question, error) {
	var model []QuestionModel
	var d []domain.Question
	if err := r.db.WithContext(ctx).Preload("Options").Where("game = ? AND type = ?", gameId, qType).Order("RANDOM()").Limit(count).Find(&model).Error; err != nil {
		return nil, err
	}
	for _, m := range model {
		d = append(d, m.ToDomain())
	}
	return d, nil
}

func (r *questionRepository) FindByGameId(ctx context.Context, gameId string) ([]domain.Question, error) {
	var models []QuestionModel
	if err := r.db.WithContext(ctx).Preload("Options").Where("game = ?", gameId).Find(&models).Error; err != nil {
		return nil, err
	}

	questions := make([]domain.Question, len(models))
	for i, m := range models {
		questions[i] = m.ToDomain()
	}
	return questions, nil
}

func (m QuestionModel) ToDomain() domain.Question {
	options := make([]domain.Option, len(m.Options))
	for i, o := range m.Options {
		options[i] = o.ToDomain()
	}

	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	return domain.Question{
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAt,
		Content:   m.Content,
		Game:      m.Game,
		Options:   options,
	}
}
