package main

import "fmt"

func main() {
	var a = 10
	var b = 10

	var plusResult = a + b
	var minusResult = a - b
	var timeResult = a * b
	var divideResult = a / b
	var modulusResult = a % b

	fmt.Println(plusResult, minusResult, timeResult, divideResult, modulusResult)

	//augmented assignment
	plusResult += a
	minusResult -= a
	timeResult *= a
	divideResult /= a
	modulusResult %= a

	fmt.Println(plusResult, minusResult, timeResult, divideResult, modulusResult)

	//unary operator
	plusResult++
	minusResult--

	fmt.Println(plusResult, minusResult, timeResult, divideResult, modulusResult)
}
