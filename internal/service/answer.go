package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
	"zen_api/internal/dto"
	"zen_api/internal/repository"
)

type AnswerService interface {
	GetAnswers() ([]datastruct.Answer, error)
	CreateAnswer(answer dto.Answer) (*int64, error)
	UpdateAnswer(answer dto.Answer) (*datastruct.Answer, error)
	DeleteAnswer(userID int64, id int64) error
}

type answerService struct {
	dao repository.DAO
}

func NewAnswerService(dao repository.DAO) AnswerService {
	return &answerService{dao: dao}
}

func (a *answerService) GetAnswers() ([]datastruct.Answer, error) {
	answers, err := a.dao.NewAnswerQuery().GetAnswers()
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (a *answerService) CreateAnswer(answer dto.Answer) (*int64, error) {
	user, err := a.dao.NewUserQuery().GetUser(answer.UserID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		if user.Role == datastruct.ADMIN {
			id, err := a.dao.NewAnswerQuery().CreateAnswer(datastruct.Answer{
				ID:         answer.ID,
				QuestionID: answer.QuestionID,
				Answer:     answer.Answer,
				Score:      answer.Score,
			})
			if err != nil {
				return nil, err
			}
			return id, nil
		}
		return nil, status.Errorf(codes.PermissionDenied, "you have no access")
	}
	return nil, err
}

func (a *answerService) UpdateAnswer(answer dto.Answer) (*datastruct.Answer, error) {
	user, err := a.dao.NewUserQuery().GetUser(answer.UserID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		if user.Role == datastruct.ADMIN {
			updatedAnswer, err := a.dao.NewAnswerQuery().UpdateAnswer(datastruct.Answer{
				ID:         answer.ID,
				QuestionID: answer.QuestionID,
				Answer:     answer.Answer,
				Score:      answer.Score,
			})
			if err != nil {
				return nil, err
			}
			return updatedAnswer, nil
		}
		return nil, status.Errorf(codes.PermissionDenied, "you have no access")
	}
	return nil, err
}

func (a *answerService) DeleteAnswer(userID int64, id int64) error {
	user, err := a.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}
	if user != nil {
		if user.Role == datastruct.ADMIN {
			err = a.dao.NewAnswerQuery().DeleteAnswer(id)
			if err != nil {
				return err
			}
			return nil
		}
		return status.Errorf(codes.PermissionDenied, "you have no access")
	}
	return err
}
