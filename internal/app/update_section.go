package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UpdateCourseSection(ctx context.Context, req *desc.UpdateCourseSectionRequest) (*desc.UpdateCourseSectionResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	filePath, err := m.fileUploaderService.UploadFile(req.GetFile())
	if err != nil {
		return nil, err
	}

	oldCourseSection, err := m.sectionService.GetSectionByID(req.GetId())
	if err != nil {
		return nil, err
	}

	section := dto.Section{
		UserID:      userID,
		ID:          req.GetId(),
		CourseID:    req.GetCourseId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		FilePath:    filePath}

	updatedSection, err := m.sectionService.UpdateCourseSection(section)
	if err != nil {
		m.fileUploaderService.DeleteFile(filePath)
		return nil, err
	}
	err = m.fileUploaderService.DeleteFile(oldCourseSection.FilePath)
	if err != nil {
		return nil, err
	}

	return &desc.UpdateCourseSectionResponse{Id: updatedSection.ID, CourseId: updatedSection.CourseID, Title: updatedSection.Title,
		Description: updatedSection.Description, FileName: updatedSection.FilePath}, nil
}
