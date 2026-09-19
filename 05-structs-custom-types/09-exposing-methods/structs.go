package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserInput("Enter your first Name: ")
	userLastName := getUserInput("Enter your last name: ")
	userBirthDate := getUserInput("Enter your bithdate (mm/dd/yyyy): ")

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.DeleteUser()
	appUser.OutputUserDetails()

}

func getUserInput(prompt string) string {
	var userInput string
	fmt.Println(prompt)
	fmt.Scanln(&userInput)
	return userInput
}