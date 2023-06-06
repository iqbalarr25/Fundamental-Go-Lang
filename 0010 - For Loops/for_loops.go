package main

import "fmt"

func main() {
	count := 0
	for count < 5 {
		fmt.Println("Perulangan ke:", count)
		count++
	}

	for i := 0; i < count; i++ {
		fmt.Println("Perulangan ke:", i)
	}

	fruits := [...]string{
		"Apple",
		"Orange",
		"Manggo",
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for index, fruit := range fruits {
		fmt.Println("index", index, "=", fruit)
	}

	person := map[string]string{
		"name":   "Iqbal",
		"age":    "22",
		"status": "Kawin",
	}

	for key, value := range person {
		fmt.Println("key", key, "=", value)
	}
}
