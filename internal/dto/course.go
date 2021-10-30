package dto

type Course struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Price       int64
	Sections    []Section
}
