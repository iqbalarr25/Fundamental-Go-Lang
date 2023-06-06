package main

import "fmt"

func main() {
	fullName, age, isMarried := getProfile()
	fmt.Println(fullName, age, isMarried)
}

// function with return named multiple value
func getProfile() (fullName string, age int, isMarried bool) {
	fullName = "Iqbal Arrafi"
	age = 22
	isMarried = true

	return
}
