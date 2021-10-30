package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) CreateCourseSection(ctx context.Context, req *desc.CreateCourseSectionRequest) (*desc.CreateCourseSectionResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	filePath, err := m.fileUploaderService.UploadFile(req.GetFile())
	if err != nil {
		return nil, err
	}

	section := dto.Section{
		UserID:      userID,
		CourseID:    req.GetCourseId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		FilePath:    filePath}

	id, err := m.sectionService.CreateCourseSection(section)
	if err != nil {
		err = m.fileUploaderService.DeleteFile(filePath)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &desc.CreateCourseSectionResponse{Id: *id}, nil
}
