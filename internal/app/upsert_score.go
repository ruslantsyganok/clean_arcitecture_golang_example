package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UpsertScore(ctx context.Context, req *desc.UpsertScoreRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	score := make([]dto.UpsertIndicatorScore, 0, len(req.GetIndicatorScore()))
	for _, item := range req.GetIndicatorScore() {
		score = append(score, dto.UpsertIndicatorScore{
			UserID:      userID,
			Title:       item.GetTitle(),
			IndicatorID: item.GetIndicatorId(),
			Score:       item.GetScore()})
	}
	err = m.scoreService.UpsertScore(score)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
