package service

import (
	"zen_api/internal/dto"
	"zen_api/internal/repository"
)

type PollService interface {
	GetPoll() ([]dto.PollQuestion, error)
}

type pollService struct {
	dao repository.DAO
}

func NewPollService(dao repository.DAO) PollService {
	return &pollService{dao: dao}
}

func (p *pollService) GetPoll() ([]dto.PollQuestion, error) {
	questions, err := p.dao.NewQuestionQuery().GetQuestions()
	if err != nil {
		return nil, err
	}

	answers, err := p.dao.NewAnswerQuery().GetAnswers()
	if err != nil {
		return nil, err
	}

	pollQuestion := make([]dto.PollQuestion, 0, len(questions))
	for i, question := range questions {
		var pollAnswer []dto.PollAnswer
		pollQuestion = append(pollQuestion, dto.PollQuestion{
			ID:          question.ID,
			IndicatorID: question.IndicatorID,
			Title:       question.Title,
		})
		for _, answer := range answers {
			if question.ID == answer.QuestionID {
				pollAnswer = append(pollAnswer, dto.PollAnswer{
					ID:         answer.ID,
					QuestionID: answer.QuestionID,
					Answer:     answer.Answer,
					Score:      answer.Score})
			}
			pollQuestion[i].Answers = pollAnswer
		}
	}
	return pollQuestion, nil
}
