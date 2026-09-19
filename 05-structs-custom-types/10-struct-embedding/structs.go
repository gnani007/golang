package main

import (
	"fmt"
	"example.com/structs/user"
)


func main() {
	userFirstName := getUserInput("Enter your firstName: ")
	userLastName := getUserInput("Enter your last Name: ")
	userBirthDate := getUserInput("Enter your birthdate (mm/dd/yyyy): ")

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("gnanesdct@gmail.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserInput(prompt string) string {
	var userInput string
	fmt.Println(prompt)
	fmt.Scanln(&userInput)
	return userInput
}