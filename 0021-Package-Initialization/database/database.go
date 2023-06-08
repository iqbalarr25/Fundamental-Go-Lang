package database

import "fmt"

var connection string

func init() {
	fmt.Println("Initialization database")
	connection = "Postgre"
}

func GetDatabase() string {
	return connection
}
