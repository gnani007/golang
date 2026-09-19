package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRatio = 2.5
	var expectedReturnRatio = 5.5
	var investmentAmount float64
	var years float64

	outputText("Enter Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter expected return: ")
	fmt.Scan(&expectedReturnRatio)

	outputText("Enter years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRatio/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRatio/100, years)

	formattedFV := fmt.Sprintf("Future return value: %.f1\n", futureValue)
	formattedRV := fmt.Sprintf("Future real value (adjusted for inflation) %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRV)

}

func outputText(text string) {
	fmt.Print(text)
}
