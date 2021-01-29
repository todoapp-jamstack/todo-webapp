package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Nickname string
	Password string
}

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
	user := new(User)

	// fill the object with values
	user.Nickname = r.FormValue("Username")
	user.Password = r.FormValue("Password")

	// convert it to json
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	// return it
	fmt.Fprintf(w, string(b))

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/users/login", login)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
