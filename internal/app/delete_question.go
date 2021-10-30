package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) DeleteQuestion(ctx context.Context, req *desc.DeleteQuestionRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err = m.questionService.DeleteQuestion(userID, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
