package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
	"zen_api/internal/dto"
	"zen_api/internal/repository"
)

type QuestionService interface {
	GetQuestions() ([]datastruct.Question, error)
	CreateQuestion(question dto.Question) (*int64, error)
	UpdateQuestion(question dto.Question) (*datastruct.Question, error)
	DeleteQuestion(userID int64, id int64) error
}

type questionService struct {
	dao repository.DAO
}

func NewQuestionService(dao repository.DAO) QuestionService {
	return &questionService{dao: dao}
}

func (q *questionService) GetQuestions() ([]datastruct.Question, error) {
	questions, err := q.dao.NewQuestionQuery().GetQuestions()
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (q *questionService) CreateQuestion(question dto.Question) (*int64, error) {
	user, err := q.dao.NewUserQuery().GetUser(question.UserID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		if user.Role == datastruct.ADMIN {
			id, err := q.dao.NewQuestionQuery().CreateQuestion(question.IndicatorID, question.Title)
			if err != nil {
				return nil, err
			}
			return id, nil
		}
		return nil, status.Errorf(codes.PermissionDenied, "you have no access")
	}
	return nil, err
}

func (q *questionService) UpdateQuestion(question dto.Question) (*datastruct.Question, error) {
	user, err := q.dao.NewUserQuery().GetUser(question.UserID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		if user.Role == datastruct.ADMIN {
			updatedQuestion, err := q.dao.NewQuestionQuery().UpdateQuestion(question.ID, question.Title)
			if err != nil {
				return nil, err
			}
			return updatedQuestion, nil
		}
		return nil, status.Errorf(codes.PermissionDenied, "you have no access")
	}
	return nil, nil
}

func (q *questionService) DeleteQuestion(userID int64, id int64) error {
	user, err := q.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}
	if user != nil {
		if user.Role == datastruct.ADMIN {
			err = q.dao.NewQuestionQuery().DeleteQuestion(id)
			if err != nil {
				return err
			}
			return nil
		}
		return status.Errorf(codes.PermissionDenied, "you have no access")
	}
	return err
}
