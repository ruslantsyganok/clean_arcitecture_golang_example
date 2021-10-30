package service

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
	"zen_api/internal/dto"
	"zen_api/internal/repository"
)

type ReviewService interface {
	GetReviews(courseID int64) ([]dto.GetReview, error)
	UpsertReview(review dto.UpsertReview) (*int64, error)
	DeleteReview(review dto.DeleteReview) error
}

type reviewService struct {
	dao repository.DAO
}

func NewReviewService(dao repository.DAO) ReviewService {
	return &reviewService{dao: dao}
}

func (r *reviewService) GetReviews(courseID int64) ([]dto.GetReview, error) {
	reviews, err := r.dao.NewReviewQuery().GetReviews(courseID)
	if err != nil {
		return nil, err
	}
	reviewsDto := make([]dto.GetReview, 0, len(reviews))
	for _, review := range reviews {
		user, err := r.dao.NewUserQuery().GetUser(review.UserID)
		if err != nil {
			return nil, err
		}
		reviewsDto = append(reviewsDto, dto.GetReview{
			ID:        review.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Feedback:  review.Feedback,
			Date:      *review.CreatedAt})
	}
	return reviewsDto, nil
}

func (r *reviewService) UpsertReview(review dto.UpsertReview) (*int64, error) {
	transaction, err := r.dao.NewTransactionQuery().GetTransactionByUserCourseIDs(review.UserID, review.CourseID)
	if err != nil {
		log.Println("user haven't paid the course")
	}

	course, err := r.dao.NewCourseQuery().GetCourse(review.CourseID)
	if err != nil {
		return nil, err
	}

	reviewID, err := r.dao.NewReviewQuery().GetReviewIdByCourseIdUserId(review.CourseID, review.UserID)
	if err != nil {
		log.Println("review doesn't exist yet")
	}

	if transaction.Status == datastruct.PAID || course.Price == 0 {
		if reviewID != nil {
			reviewID, err = r.dao.NewReviewQuery().UpdateReview(datastruct.Review{
				ID:       *reviewID,
				UserID:   review.UserID,
				CourseID: review.CourseID,
				Feedback: review.Feedback,
			})
			if err != nil {
				return nil, err
			}
			return reviewID, nil
		}
		reviewID, err = r.dao.NewReviewQuery().CreateReview(datastruct.Review{
			UserID:   review.UserID,
			CourseID: review.CourseID,
			Feedback: review.Feedback,
		})
		if err != nil {
			return nil, err
		}
		return reviewID, nil
	}
	return nil, status.Errorf(codes.PermissionDenied, "you cannot create review for the course")
}

func (r *reviewService) DeleteReview(review dto.DeleteReview) error {
	user, err := r.dao.NewUserQuery().GetUser(review.UserID)
	if err != nil {
		return err
	}

	id, err := r.dao.NewReviewQuery().GetReviewIdByCourseIdUserId(review.CourseID, review.UserID)
	if err != nil {
		return err
	}

	if user.Role == datastruct.ADMIN || *id == review.ID {
		err = r.dao.NewReviewQuery().DeleteReview(review.ReviewID)
		if err != nil {
			return err
		}
		return nil
	}
	return status.Errorf(codes.PermissionDenied, "you have no access")
}
