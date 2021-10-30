package datastruct

const QuestionTableName = "question"

type Question struct {
	ID          int64  `db:"id"`
	IndicatorID int64  `db:"indicator_id"`
	Title       string `db:"title"`
}
