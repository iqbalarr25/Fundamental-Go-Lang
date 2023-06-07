package main

import (
	"errors"
	"fmt"
)

func main() {
	result, error := pembagian(10, 0)
	if error == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", error.Error())
	}

}

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}
