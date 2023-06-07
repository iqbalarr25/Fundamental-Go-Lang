package main

import "fmt"

type Biodata struct {
	name      string
	age       int
	isMarried bool
}

func main() {
	biodata := Biodata{"Iqbal Arrafi", 22, true}
	addOneAge(&biodata)
	fmt.Println(biodata)
}

func addOneAge(biodata *Biodata) {
	biodata.age++
}
