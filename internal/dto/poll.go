package dto

type PollQuestion struct {
	ID          int64
	IndicatorID int64
	Title       string
	Answers     []PollAnswer
}

type PollAnswer struct {
	ID         int64
	QuestionID int64
	Answer     string
	Score      int64
}
