package repository

import (
	"github.com/Masterminds/squirrel"
	"zen_api/internal/datastruct"
)

type QuestionQuery interface {
	GetQuestions() ([]datastruct.Question, error)
	CreateQuestion(indicatorID int64, title string) (*int64, error)
	UpdateQuestion(id int64, title string) (*datastruct.Question, error)
	DeleteQuestion(id int64) error
}

type questionQuery struct{}

func (q *questionQuery) GetQuestions() ([]datastruct.Question, error) {
	qb := pgQb().
		Select("id", "indicator_id", "title").
		From(datastruct.QuestionTableName)

	rows, err := qb.Query()
	if err != nil {
		return nil, err
	}
	var questions []datastruct.Question
	var question datastruct.Question
	for rows.Next() {
		err = rows.Scan(&question.ID, &question.IndicatorID, &question.Title)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (q *questionQuery) CreateQuestion(indicatorID int64, title string) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.QuestionTableName).
		Columns("indicator_id", "title").
		Values(indicatorID, title).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (q *questionQuery) UpdateQuestion(id int64, title string) (*datastruct.Question, error) {
	qb := pgQb().Update(datastruct.QuestionTableName).SetMap(map[string]interface{}{
		"title": title,
	}).Where(squirrel.Eq{"id": id}).Suffix("RETURNING id, title")

	var question datastruct.Question
	err := qb.QueryRow().Scan(&question.ID, &question.Title)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (q *questionQuery) DeleteQuestion(id int64) error {
	qb := pgQb().
		Delete(datastruct.QuestionTableName).
		From(datastruct.QuestionTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
