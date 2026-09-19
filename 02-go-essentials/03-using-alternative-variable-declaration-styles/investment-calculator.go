package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 10000
	expectedReturnRate := 5.7
	years := 10

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Print(futureValue)
}
