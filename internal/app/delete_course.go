package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) DeleteCourse(ctx context.Context, req *desc.DeleteCourseRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err = m.courseService.DeleteCourse(req.GetCourseId(), userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
