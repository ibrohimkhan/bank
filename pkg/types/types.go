package types

// Money represented by cents
type Money int64

// Currency represents code of currency
type Currency string

// Currency code
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN payment card number
type PAN string

// Card represents payment info
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

// Category of things used in payment
type Category string

// Payment представляет информацию о платеже
type Payment struct {
	ID 			int
	Amount 		Money
	Category 	Category
}

// PaymentSource - source info
type PaymentSource struct {
	Type		string
	Number		string
	Balance		Money
}