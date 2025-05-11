// The Strategy pattern is a behavioral design pattern that enables selecting an algorithm's behavior at runtime.
// It defines a family of algorithms, encapsulates each one, and makes them interchangeable.

// Key Components
// Strategy Interface: Defines the common interface for all concrete strategies
// Concrete Strategies: Implement the algorithm using the Strategy interface
// Context: Maintains a reference to a Strategy object and uses it to execute the algorithm

package main

import "fmt"

// PaymentStrategy is the strategy interface
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardPayments is a concrete strategy
type CreditCardPayments struct{}

func (c *CreditCardPayments) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card", amount)
}

// PayPalPayments is a concrete strategy
type PayPalPayments struct{}

func (c *PayPalPayments) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal", amount)
}

// BankTransferPayment is a concrete strategy
type BankTransferPayment struct{}

func (c *BankTransferPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Bank Transfer", amount)
}

// PaymentContext is the context that uses a payment strategy
type PaymentContext struct {
	Strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategie(strategy PaymentStrategy) {
	p.Strategy = strategy
}

func (p *PaymentContext) ExecuteStrategie(amount float64) string {
	return p.Strategy.Pay(amount)
}

func main() {
	context := PaymentContext{}

	paymentMode := "Card"

	switch paymentMode {
	case "Card":
		context.SetStrategie(&CreditCardPayments{})
	case "PayPal":
		context.SetStrategie(&PayPalPayments{})
	case "Bank":
		context.SetStrategie(&BankTransferPayment{})
	}
	res := context.ExecuteStrategie(10000)

	fmt.Println("res", res)
	return
}
