package main

import "fmt"

func main() {

	//
	// data type int
	//

	//int8 max value is 127 & min value is -128
	var intNum8 int8 = 127

	//int16 max value is 32767 & min value is -32768
	var intNum16 int16 = 32767

	//int32 max value is 2147483647 & min value is -2147483648
	var intNum32 int32 = 2147483647

	//int64 max value is 9223372036854775807 & min value is -9223372036854775808
	var intNum64 int64 = 9223372036854775807

	fmt.Println(intNum8, intNum16, intNum32, intNum64)

	//
	// data type uint
	//

	//uint8 max value is 255 & min value is 0
	var uintNum8 uint8 = 255

	//uint8 max value is 65535 & min value is 0
	var uintNum16 uint16 = 65535

	//uint8 max value is 4294967295 & min value is 0
	var uintNum32 uint32 = 4294967295

	//uint8 max value is 18446744073709551615 & min value is 0
	var uintNum64 uint64 = 18446744073709551615

	fmt.Println(uintNum8, uintNum16, uintNum32, uintNum64)

	//
	// data type float
	//

	//float32 max value is 3.4x10"38 & min value is 1.18x10"-38
	var floatNum32 float32 = 1.1

	//float64 max value is 1.80x10"308 & min value is 2.23x10"-308
	var floatNum64 float64 = 1.1

	fmt.Println(floatNum32, floatNum64)

	//
	// data type complex
	//

	//complex64 complex number with float32 real & imaginary parts
	var complexNum64 complex64 = 1

	//complex128 complex number with float64 real & imaginary parts
	var complexNum128 complex64 = 1

	fmt.Println(complexNum64, complexNum128)
}
