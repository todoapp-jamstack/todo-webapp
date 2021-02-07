package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// this struct rapresent the response json
type Response struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

// function that return a json response of an action that succeeded
func Success(w http.ResponseWriter, message string, data interface{}) {

	// set HTTP header to 200 OK
	w.WriteHeader(http.StatusOK)

	// create response code
	response := Response{Message: message, Status: true, Data: data}

	// return return JSON
	returnJson(w, response)

}

// function that return a json response of an action that failed
func Error(w http.ResponseWriter, errorCode int, message string) {

	// set HTTP header to wathever error code given by the dev
	w.WriteHeader(errorCode)

	// create response code
	response := Response{Message: message, Status: false}

	// return return JSON
	returnJson(w, response)

}

// this function return a json
func returnJson(w http.ResponseWriter, response Response) {
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	// return it
	fmt.Fprintf(w, string(b))
}
