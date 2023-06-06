package main

import "fmt"

func main() {
	name := getHello("David")
	fmt.Println(name)
}

//function with return value
func getHello(name string) string {
	return "Hello " + name
}
