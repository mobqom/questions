package grpc_controller

import (
	"context"

	"github.com/mobqom/questions/internal/dto"
	"github.com/mobqom/questions/internal/usecase"
	optionsv1 "github.com/mobqom/questions/proto/v1/option"
)

type OptionsServer struct {
	optionsv1.UnimplementedOptionsServiceServer
	useCase usecase.OptionsUseCase
}

func NewOptionsServer(uc usecase.OptionsUseCase) *OptionsServer {
	return &OptionsServer{useCase: uc}
}

func (s *OptionsServer) AddOption(ctx context.Context, req *optionsv1.AddOptionRequest) (*optionsv1.AddOptionResponse, error) {
	o, err := s.useCase.AddOption(ctx, dto.AddOptionDto{
		Content:    req.Content,
		QuestionID: uint(req.QuestionId),
	})
	if err != nil {
		return nil, err
	}

	return &optionsv1.AddOptionResponse{
		Option: &optionsv1.Option{
			Id:         uint32(o.ID),
			Content:    o.Content,
			QuestionId: uint32(o.QuestionID),
		},
	}, nil
}

func (s *OptionsServer) FindOptionsByQuestionId(ctx context.Context, req *optionsv1.FindOptionsByQuestionIdRequest) (*optionsv1.FindOptionsByQuestionIdResponse, error) {
	options, err := s.useCase.FindByQuestionID(ctx, uint(req.QuestionId))
	if err != nil {
		return nil, err
	}

	res := &optionsv1.FindOptionsByQuestionIdResponse{
		Options: make([]*optionsv1.Option, 0, len(options)),
	}

	for _, o := range options {
		res.Options = append(res.Options, &optionsv1.Option{
			Id:         uint32(o.ID),
			Content:    o.Content,
			QuestionId: uint32(o.QuestionID),
		})
	}

	return res, nil
}
