package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
	"zen_api/internal/dto"
	"zen_api/internal/repository"
)

type IndicatorService interface {
	CreateIndicator(indicator dto.Indicator) (*int64, error)
	GetIndicators() ([]datastruct.Indicator, error)
	DeleteIndicator(indicatorID int64, userID int64) error
	UpdateIndicator(indicator dto.Indicator) (*datastruct.Indicator, error)
}

type indicatorService struct {
	dao repository.DAO
}

func NewIndicatorService(dao repository.DAO) IndicatorService {
	return &indicatorService{dao: dao}
}

func (i *indicatorService) CreateIndicator(indicator dto.Indicator) (*int64, error) {
	user, err := i.dao.NewUserQuery().GetUser(indicator.UserID)
	if err != nil {
		return nil, err
	}
	if user.Role == datastruct.ADMIN {
		id, err := i.dao.NewIndicatorQuery().CreateIndicator(indicator.Title, indicator.Description)
		if err != nil {
			return nil, err
		}
		return id, nil
	}
	return nil, status.Errorf(codes.PermissionDenied, "you have no access")
}

func (i *indicatorService) GetIndicators() ([]datastruct.Indicator, error) {
	indicators, err := i.dao.NewIndicatorQuery().GetIndicators()
	if err != nil {
		return nil, err
	}
	return indicators, nil
}

func (i *indicatorService) DeleteIndicator(indicatorID int64, userID int64) error {
	user, err := i.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}

	if user.Role == datastruct.ADMIN {
		err = i.dao.NewIndicatorQuery().DeleteIndicator(indicatorID)
		if err != nil {
			return err
		}
		return nil
	} else {
		return status.Errorf(codes.PermissionDenied, "you have no access")
	}
}

func (i *indicatorService) UpdateIndicator(indicator dto.Indicator) (*datastruct.Indicator, error) {
	user, err := i.dao.NewUserQuery().GetUser(indicator.UserID)
	if err != nil {
		return nil, err
	}

	if user.Role == datastruct.ADMIN {
		updateIndicator, err := i.dao.NewIndicatorQuery().UpdateIndicator(datastruct.Indicator{ID: indicator.ID,
			Title: indicator.Title, Description: indicator.Description})
		if err != nil {
			return nil, err
		}
		return updateIndicator, nil
	} else {
		return nil, status.Errorf(codes.PermissionDenied, "you have no access")
	}
}
