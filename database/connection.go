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

type DBconfig struct {
	name string
	username string
	password string
	port string
	host string
}


func connection() *sql.DB{
	
	config := DBconfig{
	name : loadCredentialsFromEnv("DATABASE"),
	username : loadCredentialsFromEnv("DBUSER"),
	password : loadCredentialsFromEnv("Password"),
	port : loadCredentialsFromEnv("PORT"),
	host : loadCredentialsFromEnv("HOST"),
	}

	fmt.Println(config.name)
	fmt.Println(config.username)
	fmt.Println(config.password)
	fmt.Println(config.port)
	fmt.Println(config.host)


	mysql := config.username + ":" + config.password + "@tcp(" + config.host + ":" + config.port + ")/" + config.name 
	fmt.Println(mysql)
	
	db, err := sql.Open("mysql", mysql)
	if err != nil {
		fmt.Println("Error connecting to database")
		fmt.Println(err)
		panic(err.Error())
	}



	err = db.Ping()

	if err != nil {
		fmt.Println("Unable to Connect to database")
	}



	return db


}


