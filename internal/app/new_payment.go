package app

import (
	"context"

	desc "zen_api/pkg"
)

func (m *MicroserviceServer) NewPayment(ctx context.Context, req *desc.PaymentRequest) (*desc.PaymentResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err = m.paymentService.NewTransaction(userID, req.GetCourseID())
	if err != nil {
		return nil, err
	}

	return &desc.PaymentResponse{}, nil
}
