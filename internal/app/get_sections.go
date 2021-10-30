package app

import (
	"context"

	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetCourseSections(ctx context.Context, req *desc.GetCourseSectionsRequest) (*desc.GetCourseSectionsResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	courseID := req.GetCourseId()

	sections, err := m.sectionService.GetCourseSections(userID, courseID)
	if err != nil {
		return nil, err
	}

	var requestSections []*desc.CourseSection

	for _, section := range sections {
		requestSections = append(requestSections, &desc.CourseSection{Id: section.ID,
			Title: section.Title, Description: section.Description, FilePath: section.FilePath})
	}

	return &desc.GetCourseSectionsResponse{CourseId: courseID, Sections: requestSections}, nil
}
