package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetPoll(ctx context.Context, empty *emptypb.Empty) (*desc.GetPollResponse, error) {
	poll, err := m.pollService.GetPoll()
	if err != nil {
		return nil, err
	}
	var questionResp []*desc.GetPollResponse_Question

	for i, question := range poll {
		var answersResp []*desc.GetPollResponse_Question_Answer
		questionResp = append(questionResp, &desc.GetPollResponse_Question{
			Id:          question.ID,
			IndicatorId: question.IndicatorID,
			Title:       question.Title,
		})
		for _, answer := range question.Answers {
			if answer.QuestionID == question.ID {
				answersResp = append(answersResp, &desc.GetPollResponse_Question_Answer{
					Id:         answer.ID,
					QuestionId: answer.QuestionID,
					Answer:     answer.Answer,
					Score:      answer.Score,
				})
			}
			questionResp[i].Answers = answersResp
		}
	}
	return &desc.GetPollResponse{Questions: questionResp}, nil
}
