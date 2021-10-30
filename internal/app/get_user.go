package app

import (
	"context"
	"log"

	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		log.Println("user isn't authorized")
	}

	user, err := m.userService.GetUser(req.GetId(), userID)
	if err != nil {
		return nil, err
	}

	return &desc.GetUserResponse{
		Id:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        string(user.Role),
		Verified:    user.Verified,
		Balance:     user.Balance}, nil
}
