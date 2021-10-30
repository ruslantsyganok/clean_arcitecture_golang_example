package datastruct

const ScoreTableName = "score"

type Score struct {
	ID          int64 `db:"id"`
	UserID      int64 `db:"user_id"`
	IndicatorID int64 `db:"indicator_id"`
	Score       int64 `db:"score"`
}
