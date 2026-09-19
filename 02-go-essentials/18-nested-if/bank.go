package main

import "fmt"

func main() {

	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int

	fmt.Println("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Account balance: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Enter the deposit money: ")
		var depostMoney float64
		fmt.Scan(&depostMoney)

		if depostMoney <= 0 {
			fmt.Println("Deposit money must be greater than zero")
			return
		}
		accountBalance += float64(depostMoney)
		fmt.Print("Account updated!. New balance is: ", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Enter withdraw amount: ")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount <= 0 || withdrawAmount > accountBalance {
			fmt.Println("Withdrawal amount must be greater than zero and let than the account balance")
			return
		}
		accountBalance -= withdrawAmount
		fmt.Print("Account updated!. Your balance is: ", accountBalance)
	} else {
		fmt.Println("GoodBye!")
	}

}
