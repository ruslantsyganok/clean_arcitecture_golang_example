package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UpdateIndicator(ctx context.Context, req *desc.UpdateIndicatorRequest) (*desc.UpdateIndicatorResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	indicator, err := m.indicatorService.UpdateIndicator(dto.Indicator{
		UserID:      userID,
		ID:          req.GetId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription()})
	if err != nil {
		return nil, err
	}
	return &desc.UpdateIndicatorResponse{Id: indicator.ID, Title: indicator.Title,
		Description: indicator.Description}, nil
}
