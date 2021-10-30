package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) CreateCourse(ctx context.Context, req *desc.CreateCourseRequest) (*desc.CreateCourseResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	course := dto.Course{
		UserID:      userID,
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice()}

	courseID, err := m.courseService.CreateCourse(course)
	if err != nil {
		return nil, err
	}

	return &desc.CreateCourseResponse{Id: *courseID}, nil
}
