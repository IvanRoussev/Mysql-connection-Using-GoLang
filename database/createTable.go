package main

import "fmt"


type Person struct {
	PersonID int
	lname string
	fname string
	address string
	city string
}

func createTable(query string) error {
	db := connection()
	

	err := db.Ping()

	if err != nil {
		fmt.Println("Unable to Connect to database")
	}

	fmt.Println("Connected to Database")

	_, err = db.Exec(query)

	return err
}