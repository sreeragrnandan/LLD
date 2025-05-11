package products

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64) string
}

type CardPaymentProcessor struct{}

func (c *CardPaymentProcessor) Pay(amount float64) string {
	return fmt.Sprintf("Amount Payed using Card $%.2f", amount)
}

type BankTransferProcessor struct{}

func (c *BankTransferProcessor) Pay(amount float64) string {
	return fmt.Sprintf("Amount Payed using Card $%.2f", amount)
}

type PayPalProcessor struct{}

func (c *PayPalProcessor) Pay(amount float64) string {
	return fmt.Sprintf("Amount Payed using Card $%.2f", amount)
}
