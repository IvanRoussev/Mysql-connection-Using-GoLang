package main

import "fmt"





func insertTable(query string) error{



	db := connection()

	err := db.Ping()

	if err != nil {
		fmt.Println("Unable to Connect to database")
	}

	fmt.Println("Connected to Database")

	_, err = db.Exec(query)    
	
	if err != nil {
		fmt.Errorf("Could not Insert Into Table: %s", err)
	}


	return err
} 