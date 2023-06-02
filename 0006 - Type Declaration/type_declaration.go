package main

import "fmt"

func main() {
	type DataPribadi map[string]interface{}
	type Married bool

	var dataPribadi = DataPribadi{
		"name":    "Iqbal Arrafi",
		"address": "Bekasi",
	}

	var isMarried Married = true
	var marriedStatus string

	if isMarried {
		marriedStatus = "already Married"
	} else {
		marriedStatus = "not Married yet"
	}

	fmt.Println(dataPribadi, " & ", marriedStatus)
}
