package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// sqlite library used to connect to the sqlite3 db
	_ "github.com/go-sql-driver/mysql"
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
	con, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatal(err)
	}

	// return db connection
	return con

}
