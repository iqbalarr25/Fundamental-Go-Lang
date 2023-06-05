package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Iqbal"
	names[1] = "Arrafi"
	names[2] = "Sugeng"

	fmt.Println(names[0], names[1], names[2])

	var numbers = [3]int{
		25,
		11,
		1,
	}

	fmt.Println(numbers[0], numbers[1], numbers[2])

	numbers[2] = 101

	//len count length array, not sum of data
	fmt.Println(len(numbers), numbers[2])

	//array with undeclared array length
	var fruits = [...]string{
		"apple",
	}

	fmt.Println(len(fruits), fruits)
}
