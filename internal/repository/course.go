package repository

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
)

type CourseQuery interface {
	GetCourse(id int64) (*datastruct.Course, error)
	CreateCourse(course datastruct.Course) (*int64, error)
	UpdateCourse(course datastruct.Course) (*datastruct.Course, error)
	DeleteCourse(id int64) error
}

type courseQuery struct{}

func (c *courseQuery) GetCourse(id int64) (*datastruct.Course, error) {
	qb := pgQb().
		Select("id", "user_id", "title", "description", "price").
		From(datastruct.CourseTableName).
		Where(squirrel.Eq{"id": id})

	var course datastruct.Course
	err := qb.QueryRow().Scan(&course.ID, &course.UserID, &course.Title, &course.Description, &course.Price)
	if err != nil {
		return &datastruct.Course{}, status.Errorf(codes.NotFound, "cannot find course with id %v", id)
	}

	return &course, nil
}

func (c *courseQuery) CreateCourse(course datastruct.Course) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.CourseTableName).
		Columns("user_id", "title", "price", "description").
		Values(course.UserID, course.Title, course.Price, course.Description).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create course %v", err)
	}
	return &id, nil
}

func (c *courseQuery) UpdateCourse(course datastruct.Course) (*datastruct.Course, error) {
	qb := pgQb().Update(datastruct.CourseTableName).SetMap(map[string]interface{}{
		"title":       course.Title,
		"description": course.Description,
		"price":       course.Price,
	}).Where(squirrel.Eq{"id": course.ID}).Suffix("RETURNING id, user_id, title, description, price")

	var updatedCourse datastruct.Course
	err := qb.QueryRow().Scan(&updatedCourse.ID, &updatedCourse.UserID,
		&updatedCourse.Title, &updatedCourse.Description, &updatedCourse.Price)
	if err != nil {
		return nil, fmt.Errorf("cannot update the course %v", err)
	}
	return &updatedCourse, nil
}

func (c *courseQuery) DeleteCourse(id int64) error {
	qb := pgQb().
		Delete(datastruct.CourseTableName).
		From(datastruct.CourseTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
