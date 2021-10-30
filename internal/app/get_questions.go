package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "zen_api/pkg"
)

func (m *MicroserviceServer) GetQuestions(ctx context.Context, empty *emptypb.Empty) (*desc.GetQuestionsResponse, error) {
	questions, err := m.questionService.GetQuestions()
	if err != nil {
		return nil, err
	}
	questionsResp := make([]*desc.GetQuestionsResponse_Question, 0, len(questions))
	for _, question := range questions {
		questionsResp = append(questionsResp, &desc.GetQuestionsResponse_Question{
			Id:          question.ID,
			IndicatorId: question.IndicatorID,
			Title:       question.Title})
	}
	return &desc.GetQuestionsResponse{Questions: questionsResp}, nil
}
