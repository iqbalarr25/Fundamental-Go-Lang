package main

import "fmt"

func main() {
	firstName, lastName := getFullName()
	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	firstName, _ = getFullName()
	fmt.Println(firstName)
}

// function with return multiple value
func getFullName() (string, string) {
	return "Iqbal", "Arrafi"
}
