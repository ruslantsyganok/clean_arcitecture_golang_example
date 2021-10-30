package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
	"zen_api/internal/dto"
	"zen_api/internal/repository"
)

type CourseService interface {
	GetCourse(courseID int64) (*dto.Course, error)
	CreateCourse(course dto.Course) (*int64, error)
	UpdateCourse(course dto.Course) (*dto.Course, error)
	DeleteCourse(courseID int64, userID int64) error
}

type courseService struct {
	dao repository.DAO
}

func NewCourseService(dao repository.DAO) CourseService {
	return &courseService{dao: dao}
}

func (c *courseService) GetCourse(courseID int64) (*dto.Course, error) {
	course, err := c.dao.NewCourseQuery().GetCourse(courseID)
	if err != nil {
		return nil, err
	}

	fullCourse := dto.Course{ID: course.ID, Title: course.Title, Description: course.Description, Price: course.Price, UserID: course.UserID}

	return &fullCourse, nil
}

func (c *courseService) CreateCourse(course dto.Course) (*int64, error) {
	courseInfo := datastruct.Course{
		UserID:      course.UserID,
		Title:       course.Title,
		Description: course.Description,
		Price:       course.Price}

	courseID, err := c.dao.NewCourseQuery().CreateCourse(courseInfo)
	if err != nil {
		return nil, err
	}

	return courseID, nil
}

func (c *courseService) UpdateCourse(course dto.Course) (*dto.Course, error) {
	// check weather you are an owner or an admin
	user, err := c.dao.NewUserQuery().GetUser(course.UserID)
	if err != nil {
		return nil, err
	}

	// access check
	selectedCourse, err := c.dao.NewCourseQuery().GetCourse(course.ID)
	if err != nil {
		return nil, err
	}

	if selectedCourse.UserID == course.UserID || user.Role == datastruct.ADMIN {
		// change course info
		dbCourse := datastruct.Course{ID: course.ID, Title: course.Title,
			Description: course.Description, Price: course.Price}

		updatedCourse, err := c.dao.NewCourseQuery().UpdateCourse(dbCourse)
		if err != nil {
			return nil, err
		}
		return &dto.Course{ID: updatedCourse.ID, Title: updatedCourse.Title,
			Description: updatedCourse.Description, Price: updatedCourse.Price}, nil
	} else {
		return nil, status.Errorf(codes.PermissionDenied, "you don't have access")
	}
}

func (c *courseService) DeleteCourse(courseID int64, userID int64) error {
	user, err := c.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}

	course, err := c.dao.NewCourseQuery().GetCourse(courseID)
	if err != nil {
		return err
	}

	if course.UserID == userID || user.Role == datastruct.ADMIN {
		err = c.dao.NewCourseQuery().DeleteCourse(courseID)
		if err != nil {
			return err
		}
	}
	return nil
}
