package main

import "fmt"

// PaymentProcessor interface for processing payments
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// RefundProcessor interface for processing refunds
type RefundProcessor interface {
	RefundPayment(amount float64) string
}

// CreditCard struct implements both PaymentProcessor and RefundProcessor
type CreditCard struct{}

func (cc CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing credit card payment of $%.2f", amount)
}

func (cc CreditCard) RefundPayment(amount float64) string {
	return fmt.Sprintf("Refunding credit card payment of $%.2f", amount)
}

// PayPal struct implements only PaymentProcessor
type PayPal struct{}

func (pp PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing PayPal payment of $%.2f", amount)
}

// ProcessPayment function that takes a PaymentProcessor
func ProcessPayment(pm PaymentProcessor, amount float64) {
	fmt.Println(pm.ProcessPayment(amount))
}

// RefundPayment function that takes a RefundProcessor
func RefundPayment(rp RefundProcessor, amount float64) {
	fmt.Println(rp.RefundPayment(amount))
}

func main() {
	cc := CreditCard{}
	pp := PayPal{}

	ProcessPayment(cc, 100.0)
	ProcessPayment(pp, 50.0)
	RefundPayment(cc, 100.0)
	// RefundPayment(pp, 50.0) // This will cause a compile-time error, as intended
}
