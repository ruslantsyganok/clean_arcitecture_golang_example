package app

import (
	"context"

	"zen_api/internal/datastruct"
	//"zen_api/internal/datastruct"
	//"zen_api/internal/service"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) CreateUser(ctx context.Context, req *desc.SignUpRequest) (*desc.SignUpResponse, error) {
	user := datastruct.Person{
		FirstName:   req.GetFirstName(),
		LastName:    req.GetLastName(),
		Email:       req.GetEmail(),
		Password:    req.GetPassword(),
		PhoneNumber: req.GetPhoneNumber(),
		Role:        "user",
	}
	id, err := m.authService.SignUp(user)
	if err != nil {
		return nil, err
	}

	return &desc.SignUpResponse{Id: *id}, nil
}
