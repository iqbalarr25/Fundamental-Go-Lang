package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr." + man.name
}

func main() {
	iqbal := Man{"Iqbal Arrafi"}
	iqbal.Married()
	fmt.Println(iqbal.name)
}
