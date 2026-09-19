package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	editAdultToYears(agePointer)

	fmt.Println("Age: ", age)

}

func editAdultToYears(agePointer *int) {
	*agePointer = *agePointer - 18
}
