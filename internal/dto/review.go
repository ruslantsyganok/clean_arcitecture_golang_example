package dto

import "time"

type UpsertReview struct {
	ID       int64
	UserID   int64
	CourseID int64
	Feedback string
}

type GetReview struct {
	ID        int64
	FirstName string
	LastName  string
	Feedback  string
	Date      time.Time
}

type DeleteReview struct {
	ID       int64
	UserID   int64
	ReviewID int64
	CourseID int64
}
