package repository

import (
	"github.com/Masterminds/squirrel"
	"zen_api/internal/datastruct"
)

type ScoreQuery interface {
	GetScore(userID int64) ([]datastruct.Score, error)
	CreateScore(scoreItem datastruct.Score) error
	UpdateScore(score datastruct.Score) error
	UpsertScore(score []datastruct.Score) error
}

type scoreQuery struct{}

func (s *scoreQuery) GetScore(userID int64) ([]datastruct.Score, error) {
	qb := pgQb().
		Select("id", "indicator_id", "score").
		From(datastruct.ScoreTableName).
		Where(squirrel.Eq{"user_id": userID})
	rows, err := qb.Query()
	if err != nil {
		return nil, err
	}
	var score []datastruct.Score
	var scr datastruct.Score
	for rows.Next() {
		err = rows.Scan(&scr.ID, &scr.IndicatorID, &scr.Score)
		if err != nil {
			return nil, err
		}
		score = append(score, scr)
	}
	return score, nil
}

func (s *scoreQuery) CreateScore(scoreItem datastruct.Score) error {
	qb := pgQb().
		Insert(datastruct.ScoreTableName).
		Columns("user_id", "indicator_id", "score").
		Values(scoreItem.UserID, scoreItem.IndicatorID, scoreItem.Score)

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s *scoreQuery) UpdateScore(score datastruct.Score) error {
	qb := pgQb().Update(datastruct.ScoreTableName).SetMap(map[string]interface{}{
		"score": score.Score,
	}).Where(squirrel.And{squirrel.Eq{"user_id": score.UserID}, squirrel.Eq{"indicator_id": score.IndicatorID}})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s *scoreQuery) UpsertScore(score []datastruct.Score) error {
	qb := pgQb().
		Insert(datastruct.ScoreTableName).
		Columns("id", "user_id", "indicator_id", "score")

	for _, item := range score {
		qb = qb.Values(
			item.ID,
			item.UserID,
			item.IndicatorID,
			item.Score)
	}

	qb = qb.Suffix("ON CONFLICT ON CONSTRAINT score_pkey DO UPDATE SET score=EXCLUDED.score")
	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
