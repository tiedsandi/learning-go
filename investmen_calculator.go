package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amout: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatedFV := fmt.Sprintf("Future Value: %v\n", futureValue)
	formatedRFV := fmt.Sprintf("Future Value (adjust of Inflation): %.1f", futureRealValue)
	// outputs information
	// fmt.Println("Future Value: ", futureValue)
	// 	fmt.Printf(`Future Value: %v
	// Future Value (adjust of Inflation): %.2f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjust of Inflation): ",futureRealValue)

	fmt.Print(formatedFV, formatedRFV)
}
