package types

// Money represented by cents
type Money int64

// Currency code
type Currency string

// Currency code values
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN payment card number
type PAN string

// Card info
type Card struct {
	ID 			int
	PAN 		PAN
	Balance 	Money
	MinBalance 	Money
	Currency 	Currency
	Color 		string
	Name 		string
	Active 		bool
}

// Category of payment
type Category string

// Status of payment
type Status string

// Status values
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment info
type Payment struct {
	ID 			int
	Amount 		Money
	Category 	Category
	Status		Status
}

// PaymentSource - source info
type PaymentSource struct {
	Type		string
	Number		string
	Balance		Money
}