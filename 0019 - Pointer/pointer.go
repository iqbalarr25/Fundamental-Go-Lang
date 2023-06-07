package main

import "fmt"

/**
sending variable value is dupicating value to another variable/function/method (pass by value)
*/

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = address2

	address2.city = "Bandung"
	address3.city = "Padang"

	//add * to bulk change with new variable
	*address2 = Address{"Trenggalek", "Jawa Timur", "Indonesia"}
	*address3 = Address{"Padang", "Sumatera Barat", "Indonesia"}

	//new pointer with default value (pointer to default value)
	address4 := new(Address)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
}
