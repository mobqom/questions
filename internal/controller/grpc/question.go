package grpc_controller

import (
	"context"

	"github.com/mobqom/questions/internal/dto"
	"github.com/mobqom/questions/internal/usecase"
	questionv1 "github.com/mobqom/questions/proto/v1/question"
)

type QuestionServer struct {
	questionv1.UnimplementedQuestionServiceServer
	useCase usecase.QuestionUseCase
}

func NewQuestionServer(uc usecase.QuestionUseCase) *QuestionServer {
	return &QuestionServer{useCase: uc}
}

func (s *QuestionServer) FindAll(ctx context.Context, req *questionv1.FindAllRequest) (*questionv1.FindAllResponse, error) {
	questions, err := s.useCase.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &questionv1.FindAllResponse{
		Questions: make([]*questionv1.Question, 0, len(questions)),
	}

	for _, q := range questions {
		res.Questions = append(res.Questions, &questionv1.Question{
			Id:      uint32(q.ID),
			Content: q.Content,
			Game:    q.Game,
		})
	}

	return res, nil
}

func (s *QuestionServer) AddQuestion(ctx context.Context, req *questionv1.AddQuestionRequest) (*questionv1.AddQuestionResponse, error) {
	q, err := s.useCase.AddQuestion(ctx, dto.AddQuestionDto{
		Content: req.Content,
		Game:    req.Game,
	})
	if err != nil {
		return nil, err
	}

	return &questionv1.AddQuestionResponse{
		Question: &questionv1.Question{
			Id:      uint32(q.ID),
			Content: q.Content,
			Game:    q.Game,
		},
	}, nil
}

func (s *QuestionServer) FindRandomQuestion(ctx context.Context, req *questionv1.FindRandomQuestionRequest) (*questionv1.FindRandomQuestionResponse, error) {
	q, err := s.useCase.FindRandomQuestion(ctx, req.GameId)
	if err != nil {
		return nil, err
	}

	if q == nil {
		return &questionv1.FindRandomQuestionResponse{}, nil
	}

	return &questionv1.FindRandomQuestionResponse{
		Question: &questionv1.Question{
			Id:      uint32(q.ID),
			Content: q.Content,
			Game:    q.Game,
		},
	}, nil
}

func (s *QuestionServer) FindByGameId(ctx context.Context, req *questionv1.FindByGameIdRequest) (*questionv1.FindByGameIdResponse, error) {
	questions, err := s.useCase.FindByGameId(ctx, req.GameId)
	if err != nil {
		return nil, err
	}

	res := &questionv1.FindByGameIdResponse{
		Questions: make([]*questionv1.Question, 0, len(questions)),
	}

	for _, q := range questions {
		res.Questions = append(res.Questions, &questionv1.Question{
			Id:      uint32(q.ID),
			Content: q.Content,
			Game:    q.Game,
		})
	}

	return res, nil
}
