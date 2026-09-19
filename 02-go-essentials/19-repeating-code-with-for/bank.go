package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.0

	for i := 0; i < 200; i++ {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Account Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Please enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Account Balance is: ", accountBalance)
		} else if choice == 2 {
			var depositMoney float64
			fmt.Print("Enter the deposit money: ")
			fmt.Scan(&depositMoney)

			if depositMoney == 0 {
				fmt.Println("Deposit money must be greater than 0")
			}
			accountBalance += depositMoney
			fmt.Println("Your account is updated!. New balance is: ", accountBalance)
		} else if choice == 3 {
			var withdrawalAmount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount > accountBalance {
				fmt.Println("Withdrawal amount should not greater than you have")
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Your account balance updated!. new balance is: ", accountBalance)
		} else {
			fmt.Println("Goodbye!")
		}
	}
}
