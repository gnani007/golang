package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age: ", age)

	adultYears := getAdultYears(age)

	fmt.Println("Adult Years: ", adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
