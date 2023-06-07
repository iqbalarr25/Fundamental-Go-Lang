package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func main() {
	var iqbal Customer
	iqbal.name = "Iqbal Arrafi"
	iqbal.address = "Bandung"
	iqbal.age = 22

	arrafi := Customer{
		name:    "Arrafi",
		address: "Bandung",
		age:     22,
	}

	azza := Customer{"Azza", "Bandung", 22}
	fmt.Println(iqbal, arrafi, azza)
}
