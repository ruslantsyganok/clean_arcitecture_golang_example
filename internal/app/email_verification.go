package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *MicroserviceServer) EmailVerification(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err = m.emailService.VerifyEmail(userID)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
