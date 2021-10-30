package service

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/gomail.v2"
	"zen_api/internal/repository"
)

type EmailService interface {
	VerifyEmail(userID int64) error
	CheckCode(userID int64, code int64) error
}

type emailService struct {
	dao repository.DAO
}

func NewEmailVerificationService(dao repository.DAO) EmailService {
	return &emailService{dao: dao}
}

func (e *emailService) VerifyEmail(userID int64) error {
	email, err := e.dao.NewUserQuery().GetEmailByUserID(userID)
	if err != nil {
		return err
	}

	user, err := e.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}

	if user.Verified {
		return fmt.Errorf("user already verified")
	}

	code := e.generateCode()
	err = e.sendEmail(code, email)
	if err != nil {
		return err
	}

	err = e.dao.NewUserQuery().UpdateEmailCode(userID, code)
	if err != nil {
		return err
	}
	return nil
}

func (e *emailService) CheckCode(userID int64, code int64) error {
	user, err := e.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}

	if int64(user.EmailCode) == code {
		err = e.dao.NewUserQuery().VerifiedTrueEmailCodeZero(userID)
		if err != nil {
			return err
		}
		return nil
	}
	return status.Errorf(codes.NotFound, "code isn't valid")
}

func (e *emailService) sendEmail(code int64, email string) error {
	emailFromLogin := "ruslan.test.test.test@gmail.com"
	emailFromPassword := "Q12345678!"

	m := gomail.NewMessage()
	m.SetHeader("To", email)
	m.SetHeader("From", emailFromLogin)
	m.SetHeader("Subject", "")
	m.SetBody("text/plain", "Ваш код регистрации: "+strconv.Itoa(int(code)))

	d := gomail.NewDialer("smtp.gmail.com", 587, emailFromLogin, emailFromPassword)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("cannot send an enail: %v", err)
	}
	return nil
}

func (e *emailService) generateCode() int64 {
	rand.Seed(time.Now().UnixNano())
	max := 9999
	min := 1000
	code := min + rand.Intn(max-min)
	return int64(code)
}
