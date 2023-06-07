package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func (customer Customer) customerData() {
	fmt.Println("Nama:", customer.name, "Alamat:", customer.address, "Umur", customer.age)
}

func main() {
	iqbal := Customer{"Iqbal Arrafi", "Bandung", 22}
	iqbal.customerData()
}
