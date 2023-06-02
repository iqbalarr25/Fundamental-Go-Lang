package main

import "fmt"

func main() {
	//const must be declared & immutable
	const (
		firstName = "waduh"
		lastName  = "junaedi"
	)

	fmt.Println(firstName, lastName)

	const enemyName string = "David"
	const friendName = "Naufal"

	fmt.Println(enemyName, " musuhnya "+friendName)
}
