package main

import "fmt"

func main() {
	// var revenue, expensess float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expensess: ")
	// fmt.Scan(&expensess)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	revenue := input("Revenue : ")
	expensess := input("Expensess : ")
	taxRate := input("Tax Rate : ")

	// ebt := getEBT(revenue, expensess)
	// profit := getProfit(ebt, taxRate)
	// ratio := getRatio(ebt, profit)

	ebt, profit, ratio := profitCalculator(revenue, expensess, taxRate)

	// fmt.Println(ebt)
	// fmt.Println(profit)
	// fmt.Println(ratio)

	formatedEBT := fmt.Sprintf("EBY Value: %.2f\n", ebt)
	fmt.Print(formatedEBT)
	fmt.Printf("Profit Value: %.f\nRatio Value: %.2f", profit, ratio)

}

func input(infoText string) (userInput float64) {
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return
}

func getEBT(revenue, expensess float64) (ebt float64) {
	ebt = revenue - expensess
	return
}

func getProfit(ebt, taxRate float64) float64 {
	profit := ebt * (1 - taxRate/100)
	return profit
}

func getRatio(ebt, profit float64) float64 {
	var ratio = ebt / profit
	return ratio
}

func profitCalculator(revenue, expensess, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expensess
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
