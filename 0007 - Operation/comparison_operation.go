package main

import "fmt"

func main() {
	var greaterThan = 10 > 5
	var lessThan = 10 < 5
	var greaterThanEqualTo = 10 >= 5
	var lessThanEqualTo = 10 <= 5
	var equalTo = 10 == 5
	var notEqualTo = 10 != 5

	var nameResult = "Iqbal" == "iqbal"

	fmt.Println(greaterThan, lessThan, greaterThanEqualTo, lessThanEqualTo, equalTo, notEqualTo, nameResult)
}
