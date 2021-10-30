package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) DeleteIndicator(ctx context.Context, req *desc.DeleteIndicatorRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err = m.indicatorService.DeleteIndicator(req.GetId(), userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
