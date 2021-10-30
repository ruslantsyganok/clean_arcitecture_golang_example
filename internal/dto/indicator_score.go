package dto

type Indicator struct {
	UserID      int64
	ID          int64
	Title       string
	Description string
}

type ScoreReq struct {
	Score []Score
}

type Score struct {
	ID          int64
	IndicatorID int64
	Score       int64
}

type IndicatorScore struct {
	ID          int64
	Title       string
	Description string
	Score       int64
}

type UpsertIndicatorScore struct {
	UserID      int64
	ScoreID     int64
	IndicatorID int64
	Title       string
	Score       int64
}
