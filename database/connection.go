package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Password>@tcp(127.0.0.1:3306)/testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Success!")
}
