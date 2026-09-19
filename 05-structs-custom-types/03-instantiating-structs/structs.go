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
	userFirstName := getUserInput("Enter first name: ")
	userLastName := getUserInput("Enter last name: ")
	userBirthdate := getUserInput("Enter birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(firstName, lastName, birthDate)

}

func outputUserDetails(firstName, lastName, birthDate string) {
	fmt.Println(firstName, lastName, birthDate)
}

func getUserInput(prompt string) string {
	var userInput string
	fmt.Println(prompt)
	fmt.Scan(&userInput)
	return userInput
}
