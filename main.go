package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-webapp/libraries/account"
	"todo-webapp/libraries/tools"
)

func main() {

	// load environment variables
	tools.LoadEnvironmentVariables()

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/users/login", account.Login)

	fmt.Printf("Starting server on port 8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
