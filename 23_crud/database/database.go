package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Connect opens a connection with the database
func Connect() (*sql.DB, error) {
	// Defining the connection string to communicate with the database
	connString := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	// Opening connection with the database
	db, err := sql.Open("mysql", connString)

	// If an error occurs
	if err != nil {
		// We should return it
		return nil, err
	}

	// Checking if connection was actually established
	if err = db.Ping(); err != nil {
		// We should return it
		return nil, err
	}

	// Otherwise, we returned the connection
	return db, nil
}
