package main

import (
	"errors"
	"fmt"
)

// ValidateOrder Validate input
func ValidateOrder(orderID int, price float64, quantity int) error {
	if orderID <= 0 || price <= 0 || quantity <= 0 {
		return errors.New("invalid input")
	}
	return nil
}

// CalculateTotal Calculate total price
func CalculateTotal(price float64, quantity int) float64 {
	return price * float64(quantity)
}

// FormatOutput Format output
func FormatOutput(orderID int, total float64) string {
	return fmt.Sprintf("Order ID: %d, Total: $%.2f", orderID, total)
}

// ProcessOrder now coordinates the separate functions
func ProcessOrder(orderID int, price float64, quantity int) (string, error) {
	// Validate input
	if err := ValidateOrder(orderID, price, quantity); err != nil {
		return "", err
	}

	// Calculate total price
	total := CalculateTotal(price, quantity)

	// Format output
	output := FormatOutput(orderID, total)
	return output, nil
}

func main() {
	result, err := ProcessOrder(1, 100.0, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
