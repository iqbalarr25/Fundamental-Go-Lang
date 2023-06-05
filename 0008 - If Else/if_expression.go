package main

import "fmt"

func main() {
	name := "Udin"

	if name == "Iqbal" {
		fmt.Println("Hello Iqbal")
	} else if name == "David" {
		fmt.Println("Hello David")
	} else {
		fmt.Println("Hello", name)
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah cukup")
	}
}
