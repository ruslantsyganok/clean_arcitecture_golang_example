package repository

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"zen_api/internal/datastruct"
)

type IndicatorQuery interface {
	CreateIndicator(title, description string) (*int64, error)
	GetIndicatorIdByTitle(title string) (*int64, error)
	GetIndicators() ([]datastruct.Indicator, error)
	UpdateIndicator(indicator datastruct.Indicator) (*datastruct.Indicator, error)
	DeleteIndicator(id int64) error
}

type indicatorQuery struct{}

func (i *indicatorQuery) CreateIndicator(title, description string) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.IndicatorTableName).
		Columns("title", "description").
		Values(title, description).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot create indicator %v", err)
	}
	return &id, nil
}

func (i *indicatorQuery) GetIndicatorIdByTitle(title string) (*int64, error) {
	qb := pgQb().
		Select("id").
		From(datastruct.IndicatorTableName).
		Where(squirrel.Eq{"title": title})

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (i *indicatorQuery) GetIndicators() ([]datastruct.Indicator, error) {
	qb := pgQb().
		Select("id", "title", "description").
		From(datastruct.IndicatorTableName).
		OrderBy("id")

	var indicators []datastruct.Indicator
	var indicator datastruct.Indicator

	rows, err := qb.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&indicator.ID, &indicator.Title, &indicator.Description)
		if err != nil {
			return nil, err
		}
		indicators = append(indicators, indicator)
	}
	return indicators, nil
}

func (i *indicatorQuery) UpdateIndicator(indicator datastruct.Indicator) (*datastruct.Indicator, error) {
	qb := pgQb().Update(datastruct.IndicatorTableName).SetMap(map[string]interface{}{
		"title":       indicator.Title,
		"description": indicator.Description,
	}).Where(squirrel.Eq{"id": indicator.ID}).Suffix("RETURNING *")

	var updatedIndicator datastruct.Indicator
	err := qb.QueryRow().Scan(&updatedIndicator.ID, &updatedIndicator.Title, &updatedIndicator.Description)
	if err != nil {
		return nil, err
	}
	return &updatedIndicator, nil
}

func (i *indicatorQuery) DeleteIndicator(id int64) error {
	qb := pgQb().
		Delete(datastruct.IndicatorTableName).
		From(datastruct.IndicatorTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
