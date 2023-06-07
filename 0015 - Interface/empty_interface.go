package main

import "fmt"

func main() {
	var data = ups(false, "Berhasil")
	fmt.Println(data)
}

func ups(isError bool, any interface{}) interface{} {
	if isError {
		panic("ERROR")
	} else {
		return any
	}
}
