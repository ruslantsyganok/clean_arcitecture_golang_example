package app

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetReviews(ctx context.Context, req *desc.GetReviewsRequest) (*desc.GetReviewsResponse, error) {
	reviews, err := m.reviewService.GetReviews(req.GetCourseId())
	if err != nil {
		return nil, err
	}

	reviewsResp := make([]*desc.GetReviewsResponse_Reviews, 0, len(reviews))
	for _, review := range reviews {
		reviewsResp = append(reviewsResp, &desc.GetReviewsResponse_Reviews{
			Firstname: review.FirstName,
			Lastname:  review.LastName,
			Feedback:  review.Feedback,
			CreatedAt: timestamppb.New(review.Date),
		})
	}
	return &desc.GetReviewsResponse{Reviews: reviewsResp}, nil
}
