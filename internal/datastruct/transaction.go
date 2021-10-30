package datastruct

const TransactionTableName = "transaction"

type Transaction struct {
	ID       int64         `db:"id"`
	UserID   int64         `db:"user_id"`
	CourseID int64         `db:"course_id"`
	Status   PaymentStatus `db:"status"`
}

type PaymentStatus string

const (
	PAID    PaymentStatus = "PAID"
	WAITING PaymentStatus = "WAITING"
)
