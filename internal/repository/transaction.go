package repository

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"zen_api/internal/datastruct"
)

type TransactionQuery interface {
	GetTransactionByUserCourseIDs(userID, courseID int64) (*datastruct.Transaction, error)
}

type transactionQuery struct{}

func (t *transactionQuery) GetTransactionByUserCourseIDs(userID, courseID int64) (*datastruct.Transaction, error) {
	qb := pgQb().
		Select("*").
		From(datastruct.TransactionTableName).
		Where(squirrel.Eq{"user_id": userID, "course_id": courseID})

	transaction := datastruct.Transaction{}

	err := qb.QueryRow().Scan(&transaction.ID, &transaction.UserID, &transaction.CourseID, &transaction.Status)
	if err != nil {
		return &datastruct.Transaction{}, fmt.Errorf("cannot get a transaction %v", err)
	}

	return &transaction, nil
}

func GetAllTransactions() {

}
