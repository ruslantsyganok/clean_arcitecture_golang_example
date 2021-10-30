package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) DeleteCourseSection(ctx context.Context, req *desc.DeleteCourseSectionRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	section, err := m.sectionService.GetSectionByID(req.GetSectionId())
	if err != nil {
		return nil, err
	}

	err = m.sectionService.DeleteCourseSection(req.GetSectionId(), userID)
	if err != nil {
		return nil, err
	}

	err = m.fileUploaderService.DeleteFile(section.FilePath)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
