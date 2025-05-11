// The Factory pattern is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

// Key Components
// Product Interface: Defines the interface of objects the factory creates

// Concrete Products: Implement the Product interface

// Factory: Creates objects without exposing the instantiation logic

package main

import (
	"factory-pattern/enum"
	"factory-pattern/factory"
	"fmt"
)

func main() {
	factory := factory.NewFactory()
	ccProcessor, _ := factory.GetPaymentProcessorFacFactory(enum.CreditCard)
	fmt.Println(ccProcessor.Pay(1000))

}
