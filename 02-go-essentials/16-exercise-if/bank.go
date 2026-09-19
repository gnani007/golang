package main

import "fmt"

func main() {
	var accountBalance float64 = 1200.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("what do you wan to do?")
	fmt.Println("1. Account Banance: ")
	fmt.Println("2. Deposit money: ")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Account Babance is: ", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Enter deposit amount:")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Account Updated!. Account balance is: ", depositAmount)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Enter withdraw amount:")
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("Account updated!. New account balance is: ", accountBalance)
	} else {
		fmt.Println("GoodBye!")
	}
}
