package repository

import (
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"zen_api/internal/datastruct"
)

type ReviewQuery interface {
	GetReviews(courseID int64) ([]datastruct.Review, error)
	CreateReview(review datastruct.Review) (*int64, error)
	UpdateReview(review datastruct.Review) (*int64, error)
	GetReviewIdByCourseIdUserId(courseID, userID int64) (*int64, error)
	DeleteReview(id int64) error
}

type reviewQuery struct{}

func (r *reviewQuery) GetReviews(courseID int64) ([]datastruct.Review, error) {
	qb := pgQb().
		Select("*").
		From(datastruct.ReviewTableName).
		Where(squirrel.Eq{"course_id": courseID})

	var reviews []datastruct.Review
	rows, err := qb.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var review datastruct.Review
		err = rows.Scan(&review.ID, &review.UserID, &review.CourseID, &review.Feedback, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("cannot get reviews %v", err)
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (r *reviewQuery) CreateReview(review datastruct.Review) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.ReviewTableName).
		Columns("user_id", "course_id", "feedback").
		Values(review.UserID, review.CourseID, review.Feedback).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot create a review %v", err)
	}
	return &id, nil
}

func (r *reviewQuery) UpdateReview(review datastruct.Review) (*int64, error) {
	qb := pgQb().
		Update(datastruct.ReviewTableName).
		SetMap(map[string]interface{}{
			"feedback":   review.Feedback,
			"updated_at": time.Now(),
		}).Where(squirrel.Eq{"id": review.ID}).Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot update a review %v", err)
	}
	return &id, nil
}

func (r *reviewQuery) GetReviewIdByCourseIdUserId(courseID, userID int64) (*int64, error) {
	qb := pgQb().
		Select("id").
		From(datastruct.ReviewTableName).
		Where(squirrel.And{squirrel.Eq{"course_id": courseID}, squirrel.Eq{"user_id": userID}})

	var id int64

	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot get review by user id and course id %v", err)
	}

	return &id, nil
}

func (r *reviewQuery) DeleteReview(id int64) error {
	qb := pgQb().
		Delete(datastruct.ReviewTableName).
		From(datastruct.ReviewTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
