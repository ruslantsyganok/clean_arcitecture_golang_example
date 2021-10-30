package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) DeleteAnswer(ctx context.Context, req *desc.DeleteAnswerRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err = m.answerService.DeleteAnswer(userID, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
