package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetScore(ctx context.Context, req *emptypb.Empty) (*desc.GetScoreResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	score, err := m.scoreService.GetScore(userID)
	if err != nil {
		return nil, err
	}
	respScore := make([]*desc.Score, 0, len(score))
	for _, item := range score {
		respScore = append(respScore, &desc.Score{
			Id:          item.ID,
			IndicatorId: item.IndicatorID,
			Score:       item.Score})
	}
	return &desc.GetScoreResponse{Score: respScore}, nil
}
