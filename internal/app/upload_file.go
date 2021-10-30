package app

import (
	"context"

	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UploadFile(ctx context.Context, req *desc.UploadFileRequest) (*desc.UploadFileResponse, error) {
	filePath, err := m.fileUploaderService.UploadFile(req.GetFile())
	if err != nil {
		return nil, err
	}
	return &desc.UploadFileResponse{FilePath: filePath}, nil
}
