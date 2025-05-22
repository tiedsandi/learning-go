package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance life")
	}

	var balanceText string = string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error happens")
		fmt.Println(err)
		fmt.Println("-----------------------")

		// return
		panic("Can't continue, sorry.")
	}

	for {
		fmt.Println("Welcome to GO Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				// return
				continue
			}

			accountBalance -= withdrawalAmount

			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Thanks for choosing our bank")
			fmt.Println("Goodbye")
			return
		}
	}
}

// func main() {
// 	var accountBalance = 1000.0

// 	// var choice int
// 	// for choice != 4 {

// 	for {
// 		fmt.Println("Welcome to GO Bank!")
// 		fmt.Println("What do you want to do?")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw money")
// 		fmt.Println("4. Exit")

// 		var choice int
// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		// wantsCheckBalance := choice == 1

// 		if choice == 1 {
// 			//  fmt.Println("Your balance is ", accountBalance)
// 			fmt.Printf("Your balance is %.2f\n", accountBalance)
// 		} else if choice == 2 {
// 			var depositAmount float64
// 			fmt.Print("Your deposit: ")
// 			fmt.Scan(&depositAmount)

// 			if depositAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				// return
// 				continue
// 			}

// 			accountBalance += depositAmount

// 			fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
// 		} else if choice == 3 {
// 			fmt.Print("Withdrawal amount: ")
// 			var withdrawalAmount float64
// 			fmt.Scan(&withdrawalAmount)

// 			if withdrawalAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				// return
// 				continue
// 			}

// 			if withdrawalAmount > accountBalance {
// 				fmt.Println("Invalid amount. You can't withdraw more than you have.")
// 				// return
// 				continue
// 			}

// 			accountBalance -= withdrawalAmount

// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 		} else {
// 			fmt.Println("Goodbye")
// 			// return
// 			break
// 		}

// 		fmt.Println("Your choice:", choice)
// 	}

// 	fmt.Println("Thanks for choosing our bank")
// }
