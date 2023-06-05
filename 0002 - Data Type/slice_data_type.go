package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[10:]

	//length of sliced array
	var lengthSlice1 = len(slice1)

	//capacity of sliced array (lowest slice to highest array)
	var capacitySlice1 = cap(slice1)

	fmt.Println(slice1, lengthSlice1, capacitySlice1)

	var slice2 = months[9:9]

	//add new data to existed array on last position, if capacity full -> create new array
	var slice3 = append(slice2, "End")

	fmt.Println(slice2, slice3)
	fmt.Println(months)

	//create new slice (data type, length, capacity)
	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Iqbal"
	newSlice[1] = "Arrafi"

	newSlice = append(newSlice, "Udin")
	newSlice = append(newSlice, "Sugeng")
	newSlice = append(newSlice, "Raga")
	newSlice = append(newSlice, "David")
	newSlice = append(newSlice, "Naufal")

	fmt.Println(newSlice)

	//length can be less or greater than original slice
	//capacity can be greater than original slice, but cant be less than original slice
	copySlice := make([]string, len(newSlice), cap(newSlice))

	//copy slice (destination, source)
	copy(copySlice, newSlice)

	fmt.Println(copySlice)
}
