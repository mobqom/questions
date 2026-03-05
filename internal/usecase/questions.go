package usecase

import (
	"context"

	"github.com/mobqom/questions/internal/domain"
	"github.com/mobqom/questions/internal/dto"
)

type QuestionUseCase interface {
	FindAll(ctx context.Context) ([]domain.Question, error)
	AddQuestion(ctx context.Context, question dto.AddQuestionDto) (domain.Question, error)
}

type questionUseCase struct {
	repo domain.QuestionRepository
}

func NewQuestionUseCase(repo domain.QuestionRepository) QuestionUseCase {
	return &questionUseCase{repo: repo}
}

func (u *questionUseCase) FindAll(ctx context.Context) ([]domain.Question, error) {
	return u.repo.FindAll(ctx)
}

func (u *questionUseCase) AddQuestion(ctx context.Context, question dto.AddQuestionDto) (domain.Question, error) {
	q := domain.Question{
		Content: question.Content,
		Game:    question.Game,
	}

	if err := u.repo.Create(ctx, &q); err != nil {
		return domain.Question{}, err
	}
	return q, nil
}
