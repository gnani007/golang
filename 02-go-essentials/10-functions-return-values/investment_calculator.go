package main

import (
	"fmt"
	"math"
)

const inflationRatio = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRatio := 5.5

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter expected return: ")
	fmt.Scan(&expectedReturnRatio)

	outputText("Enter years: ")
	fmt.Scan(&years)

	futureValue, realFutureValue := calculationFutureValue(investmentAmount, expectedReturnRatio, years)

	formattedFV := fmt.Sprintf("Future value is %.f1\n", futureValue)
	formattedRV := fmt.Sprintf("Future real value (adjusted for inflation) is %.f1\n", realFutureValue)

	fmt.Print(formattedFV, formattedRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculationFutureValue(investmentAmount, expectedReturnRatio, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRatio/100, years)
	rfv := fv / math.Pow(1+inflationRatio/100, years)

	return fv, rfv
}
