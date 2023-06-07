package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()

	resultString := result.(string)
	fmt.Println(resultString)

	//assertions using switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown type value")
	}
}
