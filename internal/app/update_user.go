package app

import (
	"context"

	"zen_api/internal/dto"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*desc.UpdateUserResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	updatedUser, err := m.userService.UpdateUser(dto.Person{
		ID:          userID,
		Email:       req.GetEmail(),
		FirstName:   req.GetFirstName(),
		LastName:    req.GetLastName(),
		PhoneNumber: req.GetPhoneNumber()})
	if err != nil {
		return nil, err
	}

	return &desc.UpdateUserResponse{Id: updatedUser.ID, FirstName: updatedUser.FirstName,
		LastName: updatedUser.LastName, Email: updatedUser.Email, PhoneNumber: updatedUser.PhoneNumber}, nil
}
