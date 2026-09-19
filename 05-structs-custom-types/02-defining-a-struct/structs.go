package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Enter your firstName: ")
	lastName := getUserData("Enter your lastName: ")
	birthDate := getUserData("Enter your birthDate (dd/mm/yyyy): ")

	showUserInfo(firstName, lastName, birthDate)
}

func showUserInfo(firstName, lastName, birthDate string) {
	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(prompt string) string {
	var userInput string
	fmt.Print(prompt)
	fmt.Scan(&userInput)
	return userInput
}
