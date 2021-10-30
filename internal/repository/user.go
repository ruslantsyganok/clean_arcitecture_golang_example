package repository

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
	"zen_api/internal/dto"
)

type UserQuery interface {
	CreateUser(user datastruct.Person) (*int64, error)
	GetUser(id int64) (*datastruct.Person, error)
	DeleteUser(userID int64) error
	UpdateUser(person dto.Person) (*datastruct.Person, error)
	GetUserPasswordByEmail(email string) (*string, error)
	GetEmailByUserID(id int64) (string, error)
	GetEmailCode(id int64) (*int64, error)
	UpdateEmailCode(userID, code int64) error
	VerifiedTrueEmailCodeZero(id int64) error
	GetUserIdByEmail(email string) (*int64, error)
}

type userQuery struct{}

func (u *userQuery) GetUserIdByEmail(email string) (*int64, error) {
	qb := pgQb().
		Select("id").
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"email": email})

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot get user id %v", err)
	}
	return &id, nil
}

func (u *userQuery) CreateUser(user datastruct.Person) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.PersonTableName).
		Columns("first_name", "email", "password", "last_name",
			"role", "verified", "email_code", "balance", "phone_number").
		Values(user.FirstName, user.Email, user.Password, user.LastName,
			user.Role, user.Verified, user.EmailCode, user.Balance, user.PhoneNumber).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (u *userQuery) GetUserPasswordByEmail(email string) (*string, error) {
	qb := pgQb().
		Select("password").
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"email": email})

	var password string
	err := qb.QueryRow().Scan(&password)
	if err != nil {
		return nil, fmt.Errorf("email and password don't match %v", err)
	}
	return &password, nil
}

func (u *userQuery) GetUser(id int64) (*datastruct.Person, error) {
	qb := pgQb().Select("id", "first_name", "last_name", "email",
		"role", "verified", "balance", "phone_number", "email_code").
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"id": id})

	user := datastruct.Person{}
	err := qb.QueryRow().Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email,
		&user.Role, &user.Verified, &user.Balance, &user.PhoneNumber, &user.EmailCode)
	if err != nil {
		return &datastruct.Person{}, status.Errorf(codes.NotFound, "cannot scan user: %v", err)
	}

	return &user, nil
}

func (u *userQuery) DeleteUser(userID int64) error {
	qb := pgQb().
		Delete(datastruct.PersonTableName).
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"id": userID})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (u *userQuery) UpdateUser(person dto.Person) (*datastruct.Person, error) {
	qb := pgQb().Update(datastruct.PersonTableName).SetMap(map[string]interface{}{
		"first_name":   person.FirstName,
		"last_name":    person.LastName,
		"email":        person.Email,
		"phone_number": person.PhoneNumber,
	}).Where(squirrel.Eq{"id": person.ID}).Suffix("RETURNING id, first_name, last_name, email, phone_number")

	var updatedPerson datastruct.Person
	err := qb.QueryRow().Scan(&updatedPerson.ID, &updatedPerson.FirstName, &updatedPerson.LastName, &updatedPerson.Email, &updatedPerson.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &updatedPerson, nil
}

func (u *userQuery) GetEmailByUserID(id int64) (string, error) {
	qb := pgQb().
		Select("email").
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"id": id})

	var email string
	err := qb.QueryRow().Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func (u *userQuery) GetEmailCode(id int64) (*int64, error) {
	qb := pgQb().
		Select("email_code").
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"id": id})

	var code int64
	err := qb.QueryRow().Scan(&code)
	if err != nil {
		return nil, fmt.Errorf("cannot get email code %v", err)
	}

	return &code, nil
}

func (u *userQuery) UpdateEmailCode(userID, code int64) error {
	qb := pgQb().
		Update(datastruct.PersonTableName).
		Set("email_code", code).
		Where(squirrel.Eq{"id": userID})

	_, err := qb.Exec()
	if err != nil {
		return fmt.Errorf("cannot update email code %v", err)
	}
	return nil
}

func (u *userQuery) VerifiedTrueEmailCodeZero(id int64) error {
	qb := pgQb().
		Update(datastruct.PersonTableName).
		SetMap(map[string]interface{}{
			"verified":   true,
			"email_code": 0,
		}).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return fmt.Errorf("cannot change verified to true %v", err)
	}
	return nil
}
