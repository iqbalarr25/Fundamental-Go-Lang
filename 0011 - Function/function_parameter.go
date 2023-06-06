package main

import "fmt"

func main() {
	//insert function to variable
	hello := sayHelloTo
	hello(getFullName())

	sayHelloTo("Iqbal", "Naufal")
}

// define function with parameters
func sayHelloTo(from string, to string) {
	fmt.Println(from, "say hello to", to)
}

// function with return multiple value
func getFullName() (string, string) {
	return "Iqbal", "Arrafi"
}
