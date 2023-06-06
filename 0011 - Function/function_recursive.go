package main

import "fmt"

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(factorialRecursive(10))
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
