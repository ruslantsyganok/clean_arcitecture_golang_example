package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) CreateQuestion(ctx context.Context, req *desc.CreateQuestionRequest) (*desc.CreateQuestionResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	id, err := m.questionService.CreateQuestion(dto.Question{
		UserID:      userID,
		IndicatorID: req.GetIndicatorId(),
		Title:       req.GetTitle()})
	if err != nil {
		return nil, err
	}
	return &desc.CreateQuestionResponse{Id: *id}, nil
}
