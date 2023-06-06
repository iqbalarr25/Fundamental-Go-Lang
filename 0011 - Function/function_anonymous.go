package main

import "fmt"

func main() {
	//anonymous function
	blacklist := func(name string) string {
		return "kamu di blacklist " + name
	}

	fmt.Println(blacklist("Iqbal"))
}
