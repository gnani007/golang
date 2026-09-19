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

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realValue := investmentAmount / math.Pow(1+inflationRatio/100, years)

	formattedFV := fmt.Sprintf("Futer value: %.1f\n", futureValue)
	formattedRV := fmt.Sprintf("Real value (adjusted for inflationRatio): %.1f\n", realValue)

	// outputs information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation): %.1f", futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)

	fmt.Print(formattedFV, formattedRV)

}
