package main

import (
	"fmt"
	"log"
)

func main() {

	query := "CREATE TABLE Persons (PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255));"

	err := createTable(query)

	if err != nil {
		log.Fatalf("Could not Create Table using this query %v", query)
	
	}

	fmt.Println("Ran Mysql Query: ", query)
}