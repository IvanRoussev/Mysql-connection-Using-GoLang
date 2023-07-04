package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func loadCredentialsFromEnv(key string) string{
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}


func main() {
	dbName := loadCredentialsFromEnv("DATABASE")
	username := loadCredentialsFromEnv("DBUSER")
	password := loadCredentialsFromEnv("Password")
	port := loadCredentialsFromEnv("PORT")
	host :=loadCredentialsFromEnv("HOST")

	// fmt.Println(dbName)
	// fmt.Println(username)
	// fmt.Println(password)
	// fmt.Println(port)
	// fmt.Println(host)


	mysql := username + ":" + password + ">@tcp(" + host + ":" + port + ")/" + dbName

	db, err := sql.Open("mysql", mysql)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Success!")
}
