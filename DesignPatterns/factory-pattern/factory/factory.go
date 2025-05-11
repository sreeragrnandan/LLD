package factory

import (
	"factory-pattern/enum"
	"factory-pattern/products"
	"fmt"
)

type PaymentProcessorFactory struct{}

func (f *PaymentProcessorFactory) GetPaymentProcessorFacFactory(processor enum.Processor) (products.PaymentProcessor, error) {
	switch processor {
	case enum.CreditCard:
		return &products.CardPaymentProcessor{}, nil
	case enum.BankTransfer:
		return &products.BankTransferProcessor{}, nil
	case enum.PayPal:
		return &products.PayPalProcessor{}, nil
	default:
		return nil, fmt.Errorf("Invalid processor")
	}
}

func NewFactory() *PaymentProcessorFactory {
	return &PaymentProcessorFactory{}
}
