package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Enter the investment amount: ")
	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter the expected return rate: ")
	outputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Enter the years: ")
	outputText("Enter the years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}
