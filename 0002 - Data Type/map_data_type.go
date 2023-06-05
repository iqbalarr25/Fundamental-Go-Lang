package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name": "Iqbal",
		"age":  "22",
	}

	person["name"] = "Iqbal Arrafi"
	person["status"] = "Jomblo"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["status"])

	//delete value with key map
	delete(person, "status")

	fmt.Println(person)

	//create new map
	book := map[string]string{}
	data := make(map[string]string)

	fmt.Println(book)
	fmt.Println(data)
}
