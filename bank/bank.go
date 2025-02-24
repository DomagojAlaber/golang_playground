package main

import (
	"fmt"

	"examples.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		panic("Can't continue. Sorry.")
	}
	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Your balance is ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var despositAmount float64
			fmt.Scan(&despositAmount)

			if despositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += despositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdrawal : ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. Not enough money.")
				continue
			}

			accountBalance -= withdrawalAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount:", accountBalance)
		default:
			fmt.Println("Thanks for choosing $tonk$ bank.")
			return
		}
	}
}
