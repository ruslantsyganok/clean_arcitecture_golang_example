package service

import (
	"fmt"

	"zen_api/internal/datastruct"
	"zen_api/internal/dto"
	"zen_api/internal/repository"
)

type ScoreService interface {
	GetScore(userID int64) ([]datastruct.Score, error)
	UpsertScore(indicatorScore []dto.UpsertIndicatorScore) error
}

type scoreService struct {
	dao repository.DAO
}

func NewScoreService(dao repository.DAO) ScoreService {
	return &scoreService{dao: dao}
}

func (s *scoreService) GetScore(userID int64) ([]datastruct.Score, error) {
	score, err := s.dao.NewScoreQuery().GetScore(userID)
	if err != nil {
		return nil, err
	}
	return score, nil
}

func (s *scoreService) UpsertScore(indicatorScore []dto.UpsertIndicatorScore) error {
	for _, item := range indicatorScore {
		if item.IndicatorID == 0 {
			indicatorID, err := s.dao.NewIndicatorQuery().GetIndicatorIdByTitle(item.Title)
			if err != nil {
				return err
			}
			err = s.dao.NewScoreQuery().CreateScore(datastruct.Score{
				ID:          item.ScoreID,
				UserID:      item.UserID,
				IndicatorID: *indicatorID,
				Score:       item.Score})
			if err != nil {
				return err
			}
		}
		if item.IndicatorID != 0 {
			err := s.dao.NewScoreQuery().UpdateScore(datastruct.Score{
				UserID:      item.UserID,
				IndicatorID: item.IndicatorID,
				Score:       item.Score})
			if err != nil {
				return err
			}
		}
	}
	return fmt.Errorf("cannot range indicatorScore")
}
