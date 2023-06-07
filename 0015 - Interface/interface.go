package main

import "fmt"

type Person struct {
	name string
}

type Animal struct {
	name string
}

func (person Person) getName() string {
	return person.name
}

func (animal Animal) getName() string {
	return animal.name
}

type HasName interface {
	getName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

func main() {
	person := Person{"Iqbal Arrafi"}
	sayHello(person)

	animal := Animal{"Raga"}
	sayHello(animal)
}
