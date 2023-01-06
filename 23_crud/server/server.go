package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser creates a new user on the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Getting the request body
	requestBody, err := ioutil.ReadAll(r.Body)

	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while reading the request body."))
		return
	}

	// Initializing the user var
	var user user

	// Trying to parse the JSON data
	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Error while converting user data to struct."))
		return
	}

	// Connecting to the database
	db, err := database.Connect()
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while connecting to the database."))
		return
	}
	// At the end, we'll close the database connection
	defer db.Close()

	// Defining the prepare statement (important to avoid SQL injection attacks)
	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while preparing the SQL query statement."))
		return
	}
	// At the end, we'll close the statement
	defer statement.Close()

	// Creating the SQL query, providing the data to be used
	insertion, err := statement.Exec(user.Name, user.Email)
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while creating the SQL query."))
		return
	}

	// Calling the SQL execution method
	insertedId, err := insertion.LastInsertId()
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while executing the SQL query."))
		return
	}

	// Finally, we'll return the created data
	// Defining the status code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created successfully! ID: %d", insertedId)))
}

// SearchUsers list all users from the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	// Connecting to the database
	db, err := database.Connect()
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while connecting to the database."))
		return
	}
	// At the end, we'll close the database connection
	defer db.Close()

	// Selecting users from database
	rows, err := db.Query("select * from users")
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while searching the users."))
		return
	}
	// At the end, we'll close the cursor
	defer rows.Close()

	var users []user

	// Reading the returned rows
	for rows.Next() {
		var user user

		// Scanning data and saving to the user
		// Data is returned according to table columns in database
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error while scanning user."))
			return
		}

		users = append(users, user)
	}

	// Returning data to the user
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error while parsing user data to JSON."))
		return
	}
}

// SearchUser retrieves a specific user from the database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	parameters := mux.Vars(r)

	// Converting the ID parameter from string to uint
	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while reading the item ID."))
		return
	}

	// Connecting to the database
	db, err := database.Connect()
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while connecting to the database."))
		return
	}
	// At the end, we'll close the database connection
	defer db.Close()

	// Selecting specific user from database
	row, err := db.Query("select * from users where id = ?", ID)
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while searching the user."))
		return
	}
	// At the end, we'll close the cursor
	defer row.Close()

	var user user

	// If data was returned
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error while scanning user."))
			return
		}
	}

	// Returning data to the user
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error while parsing user data to JSON."))
		return
	}
}

// UpdateUser updates data for a specific user on the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	parameters := mux.Vars(r)

	// Converting the ID parameter from string to uint
	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while reading the item ID."))
		return
	}

	// Getting the request body
	requestBody, err := ioutil.ReadAll(r.Body)

	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while reading the request body."))
		return
	}

	// Initializing the user var
	var user user

	// Trying to parse the JSON data
	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Error while converting user data to struct."))
		return
	}

	// Connecting to the database
	db, err := database.Connect()
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while connecting to the database."))
		return
	}
	// At the end, we'll close the database connection
	defer db.Close()

	// Defining the prepare statement (important to avoid SQL injection attacks)
	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while preparing the SQL query statement."))
		return
	}
	// At the end, we'll close the statement
	defer statement.Close()

	// Creating and executing the SQL query, providing the data to be used
	_, err = statement.Exec(user.Name, user.Email, ID)
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while updating the user."))
		return
	}

	// Finally, we'll return the updated data response
	// Defining the status code
	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser removes a specific user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	parameters := mux.Vars(r)

	// Converting the ID parameter from string to uint
	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while reading the item ID."))
		return
	}

	// Connecting to the database
	db, err := database.Connect()
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while connecting to the database."))
		return
	}
	// At the end, we'll close the database connection
	defer db.Close()

	// Defining the prepare statement (important to avoid SQL injection attacks)
	statement, err := db.Prepare("delete from users where id = ?")
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while preparing the SQL query statement."))
		return
	}
	// At the end, we'll close the statement
	defer statement.Close()

	// Creating and executing the SQL query, providing the data to be used
	_, err = statement.Exec(ID)
	// If something goes wrong
	if err != nil {
		w.Write([]byte("Error while deleting the user."))
		return
	}

	// Finally, we'll return the updated data response
	// Defining the status code
	w.WriteHeader(http.StatusNoContent)
}
