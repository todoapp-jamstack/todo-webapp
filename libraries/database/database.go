package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo-webapp/libraries/tools"

	// sqlite library used to connect to the sqlite3 db
	_ "github.com/mattn/go-sqlite3"
)

// function to get the query to create the user table
func getQueryCreateUserTable() string {
	return "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(128), password VARCHAR(128), date_creation TEXT, date_updated TEXT)"
}

// function to initialize the database
func initializeDatabase() bool {

	fmt.Println("Database empty... I will do something about that")

	database := DbConnection()

	statement, _ := database.Prepare(getQueryCreateUserTable())
	_, err := statement.Exec()

	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("User table created ✔️")

	fmt.Println("----\nAll Database created ✔️")

	return true

}

// DbConnection used to get the DB connection and doing stuff
func DbConnection() *sql.DB {

	// create db connection
	database, err := sql.Open("sqlite3", "db/todo.db")

	if err != nil {
		log.Fatal(err)
	}

	// return db connection
	return database

}

// CheckDBfolder is used to check if the db folder is created or
// in case is not then it will be created
func CheckDBfolder() bool {

	// if the folder db does not exist create it
	if _, err := os.Stat("db/todo.db"); os.IsNotExist(err) {

		// create the db folder
		if tools.CreateFolder("db") {
			// Initialize the schema of the database
			if initializeDatabase() {
				return true
			}
		}
	}

	return false

}
