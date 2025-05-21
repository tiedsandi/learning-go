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

	// fmt.Print("Investment Amout: ")
	outputText("Investment Amout: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatedFV := fmt.Sprintf("Future Value: %v\n", futureValue)
	formatedRFV := fmt.Sprintf("Future Value (adjust of Inflation): %.1f", futureRealValue)
	// outputs information
	// fmt.Println("Future Value: ", futureValue)
	// 	fmt.Printf(`Future Value: %v
	// Future Value (adjust of Inflation): %.2f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjust of Inflation): ",futureRealValue)

	fmt.Print(formatedFV, formatedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

// func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)

// 	return
// }
