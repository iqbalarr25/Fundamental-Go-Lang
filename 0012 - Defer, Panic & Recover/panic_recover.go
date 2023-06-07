package main

import "fmt"

func main() {
	runApp(true)
}

func runApp(isError bool) {
	defer endApp()
	if isError {
		panic("ERROR")

	}

	fmt.Println("Aplikasi berjalan")
}

func endApp() {

	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error:", message)
	}
	fmt.Println("Aplikasi berhenti")
}
