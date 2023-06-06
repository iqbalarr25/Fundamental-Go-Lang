package main

import "fmt"

//type declaration for function getProfile
type Profile func() (string, int, bool)

func main() {
	//define function as parameter anoter function
	sayHelloProfile(getProfile)
}

func sayHelloProfile(profile Profile) {
	name, age, isMarried := profile()
	statusText := ""
	if isMarried {
		statusText = "sudah menikah"
	} else {
		statusText = "belum menikah"
	}

	fmt.Println("Halo", name, "dengan umur", age, "dan status", statusText)
}

// function with return named multiple value
func getProfile() (fullName string, age int, isMarried bool) {
	fullName = "Iqbal Arrafi"
	age = 22
	isMarried = true

	return
}
