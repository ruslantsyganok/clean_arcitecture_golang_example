package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) CreateIndicator(ctx context.Context, req *desc.CreateIndicatorRequest) (*desc.CreateIndicatorResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	indicator := dto.Indicator{
		UserID:      userID,
		Title:       req.GetTitle(),
		Description: req.GetDescription()}
	id, err := m.indicatorService.CreateIndicator(indicator)
	if err != nil {
		return nil, err
	}
	return &desc.CreateIndicatorResponse{Id: *id}, nil
}
