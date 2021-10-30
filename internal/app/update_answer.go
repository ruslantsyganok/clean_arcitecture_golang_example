package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UpdateAnswer(ctx context.Context, req *desc.UpdateAnswerRequest) (*desc.UpdateAnswerResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	answer, err := m.answerService.UpdateAnswer(dto.Answer{
		UserID: userID,
		ID:     req.GetId(),
		Answer: req.GetAnswer(),
		Score:  req.GetScore(),
	})
	if err != nil {
		return nil, err
	}
	return &desc.UpdateAnswerResponse{
		Id:         answer.ID,
		QuestionId: answer.QuestionID,
		Answer:     answer.Answer,
		Score:      req.GetScore()}, nil
}
