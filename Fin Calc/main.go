package main

import (
	"fmt"
)

func getUserTransaction() float64 {
	var transaction float64
	fmt.Print("Укажите транзакцию (доход или расход): ")
	fmt.Scan(&transaction)
	return transaction
}
func calcBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance = balance + value
	}
	return balance
}

func main() {
	transactions := []float64{}
	var balance float64

	for {
		transaction := getUserTransaction()
		if transaction == 0 {
			break

		}
		transactions = append(transactions, transaction)
	}

	balance = calcBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f\n", balance)
}
