package main

import "fmt"

func main() {
	//if value more than int16, the number will be accumulate again from lowest value
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai16)
	fmt.Println(nilai64)
}
