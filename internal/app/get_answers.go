package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetAnswers(ctx context.Context, empty *emptypb.Empty) (*desc.GetAnswersResponse, error) {
	answers, err := m.answerService.GetAnswers()
	if err != nil {
		return nil, err
	}
	answersResp := make([]*desc.GetAnswersResponse_Answer, 0, len(answers))
	for _, answer := range answers {
		answersResp = append(answersResp, &desc.GetAnswersResponse_Answer{
			Id:         answer.ID,
			QuestionId: answer.QuestionID,
			Answer:     answer.Answer,
			Score:      answer.Score})
	}
	return &desc.GetAnswersResponse{Answers: answersResp}, nil
}
