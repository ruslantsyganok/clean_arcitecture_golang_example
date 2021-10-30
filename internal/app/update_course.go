package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UpdateCourse(ctx context.Context, req *desc.UpdateCourseRequest) (*desc.UpdateCourseResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	course := dto.Course{
		UserID:      userID,
		ID:          req.GetCourseId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice()}

	updatedCourse, err := m.courseService.UpdateCourse(course)
	if err != nil {
		return nil, err
	}

	return &desc.UpdateCourseResponse{CourseId: updatedCourse.ID, Title: updatedCourse.Title,
		Description: updatedCourse.Description, Price: updatedCourse.Price}, nil
}
