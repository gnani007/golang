package main

import (
	"math"
)

func main() {
	const inflationRatio = 4.5
	var investmentAmount float64 = 10000
	years := 10
	expectedReturnValue := 6.6

	futureValue := investmentAmount * math.Pow(1+expectedReturnValue/100, float64(years))
	realValue := futureValue / math.Pow(1+inflationRatio/100, float64(years))

	println(futureValue)
	print(realValue)

}
