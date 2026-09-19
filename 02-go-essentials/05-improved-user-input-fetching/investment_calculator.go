package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRatio = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter investment amount: ")
	fmt.Scanln(&investmentAmount)

	fmt.Print("Enter expectedReturnRate: ")
	fmt.Scanln(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scanln(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realValue := futureValue / math.Pow(1+inflationRatio/100, years)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Real value", realValue)

}
