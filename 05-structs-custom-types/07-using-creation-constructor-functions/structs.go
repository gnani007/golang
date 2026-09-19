package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createAt  time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) deleteUser() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createAt:  time.Now(),
	}
}

func main() {
	userFistName := getUserInput("Enter your first name: ")
	userLastName := getUserInput("Enter your last name: ")
	getBirthDate := getUserInput("Enter your birth date (YYYY-MM-DD): ")

	var appUser *user

	appUser = newUser(userFistName, userLastName, getBirthDate)

	appUser.outputUserDetails()
	appUser.deleteUser()
	appUser.outputUserDetails()
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}
