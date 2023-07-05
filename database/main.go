package main

import (
	"fmt"
	"log"
)

func main() {
	// person := Person{
	// 	PersonID: 1,
	// 	fname: "Ivan",
	// 	lname: "Roussev",
	// 	address: "123 mcdonald road",
	// 	city: "Vancouver",
	// }

	ctquery := "CREATE TABLE Persons (PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255));"

	err := createTable(ctquery)

	if err != nil {
		log.Fatalf("Could not Create Table using this query %v", ctquery)
	
	}

	fmt.Println("Ran Mysql Query: ", ctquery)


	itquery := ""

	err = insertTable(itquery)

	if err != nil {
		log.Fatalf("Could not Insert into Table using this query %v", itquery)
	
	}
	fmt.Println("Ran Mysql Query: ", itquery)
}