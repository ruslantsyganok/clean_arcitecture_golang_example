package datastruct

import "time"

const ReviewTableName = "review"

type Review struct {
	ID        int64      `db:"id"`
	UserID    int64      `db:"user_id"`
	CourseID  int64      `db:"course_id"`
	Feedback  string     `db:"feedback"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
