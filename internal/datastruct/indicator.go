package datastruct

const IndicatorTableName = "indicator"

type Indicator struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}
