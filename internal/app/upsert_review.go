package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UpsertReview(ctx context.Context, req *desc.UpsertReviewRequest) (*desc.UpsertReviewResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	review := dto.UpsertReview{
		UserID:   userID,
		CourseID: req.GetCourseId(),
		Feedback: req.GetFeedback()}

	id, err := m.reviewService.UpsertReview(review)
	if err != nil {
		return nil, err
	}

	return &desc.UpsertReviewResponse{Id: *id}, nil
}
