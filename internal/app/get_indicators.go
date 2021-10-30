package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetIndicators(ctx context.Context, empty *emptypb.Empty) (*desc.GetIndicatorsResponse, error) {
	indicators, err := m.indicatorService.GetIndicators()
	if err != nil {
		return nil, err
	}
	var responseIndicator []*desc.Indicator
	for _, indicator := range indicators {
		responseIndicator = append(responseIndicator, &desc.Indicator{Id: indicator.ID,
			Title: indicator.Title, Description: indicator.Description})
	}
	return &desc.GetIndicatorsResponse{Indicators: responseIndicator}, nil
}
