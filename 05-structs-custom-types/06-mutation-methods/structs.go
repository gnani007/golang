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

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserInput("Enter first name: ")
	userLastName := getUserInput("Enter last name: ")
	userBirthDate := getUserInput("Enter birthDate (mm/dd/yyyy): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createAt:  time.Now(),
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}
