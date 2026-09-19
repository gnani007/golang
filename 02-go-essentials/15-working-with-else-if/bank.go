package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice:")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Account Balance: ", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Enter deposit amount: ")
		fmt.Scan(&depositAmount)
		accountBalance += float64(depositAmount)
		fmt.Println("Account updated!. Your Account Balance is: ", accountBalance)
	}

	fmt.Println("Your Choice: ", choice)

}
