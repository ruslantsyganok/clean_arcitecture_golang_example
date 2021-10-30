package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) DeleteFile(ctx context.Context, req *desc.DeleteFileRequest) (*emptypb.Empty, error) {
	err := m.fileUploaderService.DeleteFile(req.GetFilePath())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
