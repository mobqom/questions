package usecase

import (
	"context"

	"github.com/mobqom/questions/internal/domain"
	"github.com/mobqom/questions/internal/dto"
)

type OptionsUseCase interface {
	AddOption(ctx context.Context, option dto.AddOptionDto) (domain.Option, error)
	FindByQuestionID(ctx context.Context, questionID uint) ([]domain.Option, error)
	FindAll(ctx context.Context) ([]domain.Option, error)
}

type optionsUseCase struct {
	repo domain.OptionRepository
}

func NewOptionsUseCase(repo domain.OptionRepository) OptionsUseCase {
	return &optionsUseCase{repo: repo}
}

func (u *optionsUseCase) AddOption(ctx context.Context, option dto.AddOptionDto) (domain.Option, error) {
	o := domain.Option{
		Content:    option.Content,
		QuestionID: option.QuestionID,
	}

	if err := u.repo.Create(ctx, &o); err != nil {
		return domain.Option{}, err
	}
	return o, nil
}

func (u *optionsUseCase) FindByQuestionID(ctx context.Context, questionID uint) ([]domain.Option, error) {
	return u.repo.FindByQuestionID(ctx, questionID)
}

func (u *optionsUseCase) FindAll(ctx context.Context) ([]domain.Option, error) {
	return u.repo.FindAll(ctx)
}
