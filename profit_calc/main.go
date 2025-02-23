package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64 = 25

	// fmt.Print("Enter your revenue: ")
	outputText("Enter your revenue: ")
	fmt.Scan(&revenue)

	// fmt.Print("Enter your expenses: ")
	outputText("Enter your expenses: ")
	fmt.Scan(&expenses)

	// fmt.Print("Enter the taxrate: ")
	outputText("Enter the taxrate: ")
	fmt.Scan(&taxRate)

	ebt := (revenue - expenses)
	profit := ebt / (1 - taxRate/100)
	ratio := ebt / profit

	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print("Formatted Ratio: " + formattedRatio)

	fmt.Printf("EBT %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
}

func outputText(text string) {
	fmt.Print(text)
}
