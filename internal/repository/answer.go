package repository

import (
	"github.com/Masterminds/squirrel"
	"zen_api/internal/datastruct"
)

type AnswerQuery interface {
	GetAnswers() ([]datastruct.Answer, error)
	CreateAnswer(answer datastruct.Answer) (*int64, error)
	UpdateAnswer(answer datastruct.Answer) (*datastruct.Answer, error)
	DeleteAnswer(id int64) error
}

type answerQuery struct{}

func (a *answerQuery) GetAnswers() ([]datastruct.Answer, error) {
	qb := pgQb().
		Select("*").
		From(datastruct.AnswerTableName)

	var answers []datastruct.Answer
	var answer datastruct.Answer
	rows, err := qb.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&answer.ID, &answer.QuestionID, &answer.Answer, &answer.Score)
		if err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}
	return answers, nil
}

func (a *answerQuery) CreateAnswer(answer datastruct.Answer) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.AnswerTableName).
		Columns("question_id", "answer", "score").
		Values(answer.QuestionID, answer.Answer, answer.Score).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (a *answerQuery) UpdateAnswer(answer datastruct.Answer) (*datastruct.Answer, error) {
	qb := pgQb().Update(datastruct.AnswerTableName).SetMap(map[string]interface{}{
		"answer": answer.Answer,
		"score":  answer.Score,
	}).Where(squirrel.Eq{"id": answer.ID}).Suffix("RETURNING id, question_id, answer, score")

	var updatedAnswer datastruct.Answer
	err := qb.QueryRow().Scan(&updatedAnswer.ID, &updatedAnswer.QuestionID, &updatedAnswer.Answer, &updatedAnswer.Score)
	if err != nil {
		return nil, err
	}
	return &updatedAnswer, nil
}

func (a *answerQuery) DeleteAnswer(id int64) error {
	qb := pgQb().
		Delete(datastruct.AnswerTableName).
		From(datastruct.AnswerTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
