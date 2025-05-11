package enum

type Processor int

const (
	CreditCard Processor = iota
	PayPal
	BankTransfer
)
