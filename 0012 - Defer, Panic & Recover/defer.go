package main

import "fmt"

func main() {
	runApplication(0)
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println(result)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}
