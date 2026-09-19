package main

import "fmt"

func main() {
	firstName := getUserData("Enter your firstName: ")
	lastName := getUserData("Enter your lastName: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(prompt string) string {
	var userInput string

	fmt.Print(prompt)
	fmt.Scan(&userInput)

	return userInput
}
