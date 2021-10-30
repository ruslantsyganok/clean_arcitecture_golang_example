package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) DeleteReview(ctx context.Context, req *desc.DeleteReviewRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err = m.reviewService.DeleteReview(dto.DeleteReview{
		ID:       req.GetId(),
		UserID:   userID,
		ReviewID: req.GetId(),
		CourseID: req.GetCourseId()})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
