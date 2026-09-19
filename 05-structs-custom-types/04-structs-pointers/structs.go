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
	userFistName := getUserInput("Enter first name: ")
	userLastName := getUserInput("Enter last name: ")
	userBirthDate := getUserInput("Enter birthDate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFistName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)

}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserInput(prompt string) string {
	var userInput string
	fmt.Println(prompt)
	fmt.Scan(&userInput)
	return userInput
}
