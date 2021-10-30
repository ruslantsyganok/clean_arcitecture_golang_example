package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UpdateQuestion(ctx context.Context, req *desc.UpdateQuestionRequest) (*desc.UpdateQuestionResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	question, err := m.questionService.UpdateQuestion(dto.Question{
		UserID: userID,
		ID:     req.GetId(),
		Title:  req.GetTitle()})
	if err != nil {
		return nil, err
	}
	return &desc.UpdateQuestionResponse{Id: question.ID, Title: question.Title}, nil
}
