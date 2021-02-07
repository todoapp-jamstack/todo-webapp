package account

import (
	"database/sql"
	"fmt"
	"net/http"
	"todo-webapp/libraries/database"
	"todo-webapp/libraries/response"
	"todo-webapp/libraries/tools"
)

// Account struct rapresent the user
type Account struct {
	ID       int
	Username string
	Password string
}

// Login function used to check if the user can enter in the platform
func Login(w http.ResponseWriter, r *http.Request) {

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
	account.Password = r.FormValue("Password")

	// return json
	//tools.HashAndSalt([]byte(account.Password)
	response.Success(w, "login effettuato correttamente", checkPassword(account))

}

func getQuerySelectUser() string {
	return "SELECT `id`, `username`, `password` FROM `users` WHERE `username` = ?"
}

// function to select a user by its username
func selectUser(username string) *Account {

	// get the db connection
	db := database.DbConnection()

	// when all the other action are finished close the connection with the database
	defer db.Close()

	row := db.QueryRow(getQuerySelectUser(), username)

	user := new(Account)

	err := row.Scan(&user.ID, &user.Username, &user.Password)

	// if the error is different from there are no result then throw an error
	if err != nil {
		if err != sql.ErrNoRows {
			panic(err)
		}
	}

	return user
}

// function to check if the password is correct and the login can proceed
func checkPassword(account *Account) bool {

	// get the user with the username inserted
	user := selectUser(account.Username)

	// if the user ID is 0 then there is no user with that username
	if user.ID == 0 {
		return false
	}

	// else check if the password is correct
	return tools.ComparePasswords(user.Password, []byte(account.Password))

}
