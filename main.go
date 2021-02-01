package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Username string
	Password string
}

// functon to get the query to create the user table
func getQueryCreateUserTable() string {
	return "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(128), password VARCHAR(128), date_creation TEXT, date_updated TEXT)"
}

// function to initialize the database
func initializeDatabase() bool {

	fmt.Println("Database empty... I will do something about that")

	database := dbConnection()

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

// function to hash & salt the password
func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// function to log in the user
func login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Sorry, only POST method is supported.", http.StatusBadRequest)
		return
	}

	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseMultipartForm(0); err != nil {
		fmt.Fprintf(w, "ParseMultipartForm() err: %v", err)
		return
	}

	// create object user
	account := new(Account)

	// fill the object with values
	account.Username = r.FormValue("Username")
	// hash password
	account.Password = hashAndSalt([]byte(r.FormValue("Password")))

	// convert it to json
	b, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
		return
	}
	// return it
	fmt.Fprintf(w, string(b))

}

// function to get the db connection
func dbConnection() *sql.DB {

	// create db connection
	database, err := sql.Open("sqlite3", "db/todo.db")

	if err != nil {
		log.Fatal(err)
	}

	// return db connection
	return database

}

// function to create the db folder if does not exist
func CreateDatabaseFolderIfDoesNotExist() bool {

	// if the folder db does not exist create it
	if _, err := os.Stat("db/todo.db"); os.IsNotExist(err) {

		// create the db folder
		if createFolder("db") {
			// Initialize the schema of the database
			if initializeDatabase() {
				return true
			}
		}
	}

	return false

}

// function to create a folder
func createFolder(path string) bool {

	// create db folder
	err := os.Mkdir(path, 0755)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}

func main() {

	// check for database existence
	if CreateDatabaseFolderIfDoesNotExist() {
		fmt.Println("Database Initiated, ready to go!")
	}

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/users/login", login)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
