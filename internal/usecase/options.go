package usecase

import (
	"context"

	"github.com/mobqom/questions/internal/domain"
	"github.com/mobqom/questions/internal/dto"
)

type OptionsUseCase interface {
	AddOption(ctx context.Context, option dto.AddOptionDto) (domain.Options, error)
	FindByQuestionID(ctx context.Context, questionID uint) ([]domain.Options, error)
	FindAll(ctx context.Context) ([]domain.Options, error)
}

type optionsUseCase struct {
	repo domain.OptionsRepository
}

func NewOptionsUseCase(repo domain.OptionsRepository) OptionsUseCase {
	return &optionsUseCase{repo: repo}
}

func (u *optionsUseCase) AddOption(ctx context.Context, option dto.AddOptionDto) (domain.Options, error) {
	o := domain.Options{
		Content:    option.Content,
		QuestionID: option.QuestionID,
	}

	if err := u.repo.Create(ctx, &o); err != nil {
		return domain.Options{}, err
	}
	return o, nil
}

func (u *optionsUseCase) FindByQuestionID(ctx context.Context, questionID uint) ([]domain.Options, error) {
	return u.repo.FindByQuestionID(ctx, questionID)
}

func (u *optionsUseCase) FindAll(ctx context.Context) ([]domain.Options, error) {
	return u.repo.FindAll(ctx)
}
