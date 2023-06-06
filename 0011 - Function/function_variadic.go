package main

import "fmt"

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 10))

	//insert slice to variadic function
	slice := []int{10, 10, 10, 10}
	fmt.Println(sumAll(slice...))
}

// variadic function (parameter varags)
func sumAll(numbers ...int) (total int) {
	total = 0
	for _, number := range numbers {
		total += number
	}
	return
}
