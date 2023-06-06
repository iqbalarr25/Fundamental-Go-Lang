package main

import "fmt"

func main() {
	name := "Iqbal"

	switch name {
	case "Iqbal":
		fmt.Println("Halo Iqbal")
	case "Arrafi":
		fmt.Println("Halo Arrafi")
	default:
		fmt.Println("Kamu bukan bagian dari diriku")
	}

	switch length := len(name); length > 10 {
	case true:
		fmt.Println("Adalah benar")
	case false:
		fmt.Println("Adalah salah")
	}

	length := len(name) * 2
	switch {
	case length > 10:
		fmt.Println("Lebih dari 10")
	case length < 10:
		fmt.Println("Kurang dari 10")
	default:
		fmt.Println("ini 10")
	}
}
