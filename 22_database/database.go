package main

import (
	"database/sql"
	"fmt"
	"log"

	// This is an implicit import, since the module will be required for other files
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Defining the connection string to communicate with the database
	// "[USER]:[PASSWORD]@/[DATABASE]?PARAMS"
	connString := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	// Opening connection with the database
	db, err := sql.Open("mysql", connString)

	// If an error occurs
	if err != nil {
		fmt.Println("SQL Open error")
		log.Fatal(err)
	}
	// Setting a defer to close connection, when we finish using it
	defer db.Close()

	// Checking if connection was actually established
	if err = db.Ping(); err != nil {
		fmt.Println("Ping error")
		log.Fatal(err)
	}

	fmt.Println("Connection with database has been successfully established!")

	// Selecting data from the database
	rows, err := db.Query("select * from users;")

	// If an error occurs during the query
	if err != nil {
		log.Fatal(err)
	}

	// Setting defer to close the "rows cursor" at the end
	defer rows.Close()

	// Showing returned data
	fmt.Println(rows)

}
