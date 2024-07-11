package main

import (
	"errors"
	"fmt"
)

// ProcessOrder function does everything
func ProcessOrder(orderID int, price float64, quantity int) (string, error) {
	// Validate input
	if orderID <= 0 || price <= 0 || quantity <= 0 {
		return "", errors.New("invalid input")
	}

	// Calculate total price
	total := price * float64(quantity)

	// Format output
	output := fmt.Sprintf("Order ID: %d, Total: $%.2f", orderID, total)
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
