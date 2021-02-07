package account

import (
	"fmt"
	"net/http"
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
	account.Password = tools.HashAndSalt([]byte(r.FormValue("Password")))

	// return json
	response.Success(w, "login effettuato correttamente", account)

}
