package datastruct

const AnswerTableName = "answer"

type Answer struct {
	ID         int64  `db:"id"`
	QuestionID int64  `db:"question_id"`
	Answer     string `db:"answer"`
	Score      int64  `db:"score"`
}
