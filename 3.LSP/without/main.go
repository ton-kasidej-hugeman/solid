package main

import "fmt"

// PaymentMethod interface
type PaymentMethod interface {
	ProcessPayment(amount float64) string
	RefundPayment(amount float64) string
}

// CreditCard struct
type CreditCard struct{}

func (cc CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing credit card payment of $%.2f", amount)
}

func (cc CreditCard) RefundPayment(amount float64) string {
	return fmt.Sprintf("Refunding credit card payment of $%.2f", amount)
}

// PayPal struct
type PayPal struct{}

func (pp PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing PayPal payment of $%.2f", amount)
}

// PayPal does not support refunds, but we are forced to implement the method
func (pp PayPal) RefundPayment(amount float64) string {
	return "PayPal does not support refunds"
}

// ProcessPayment function that takes a PaymentMethod
func ProcessPayment(pm PaymentMethod, amount float64) {
	fmt.Println(pm.ProcessPayment(amount))
	fmt.Println(pm.RefundPayment(amount))
}

func main() {
	cc := CreditCard{}
	pp := PayPal{}

	ProcessPayment(cc, 100.0)
	ProcessPayment(pp, 50.0)
}
