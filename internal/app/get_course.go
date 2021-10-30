package app

import (
	"context"

	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetCourse(ctx context.Context, req *desc.GetCourseRequest) (*desc.GetCourseResponse, error) {
	courseID := req.GetId()

	course, err := m.courseService.GetCourse(courseID)
	if err != nil {
		return nil, err
	}

	return &desc.GetCourseResponse{Id: course.ID, UserId: course.UserID,
		Title: course.Title, Description: course.Description,
		Price: course.Price}, nil
}
