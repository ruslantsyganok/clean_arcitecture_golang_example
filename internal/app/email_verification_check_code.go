package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) EmailVerificationCheckCode(ctx context.Context, req *desc.EmailVerificationCheckCodeRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err = m.emailService.CheckCode(userID, req.GetCode())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
