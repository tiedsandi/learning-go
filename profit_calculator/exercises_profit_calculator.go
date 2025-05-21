package main

import "fmt"

func main() {
	var revenue, expensess int
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expensess: ")
	fmt.Scan(&expensess)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expensess
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := float64(ebt) / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
