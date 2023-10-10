package main

import (
	"groupie-tracker/views"
	"net/http"
)

func main() {

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func ErrTest(w http.ResponseWriter, r *http.Request) {
	errorMessage := "An error occurred while parsing the template."
	views.ErrorHandler(w, r, errorMessage)
}
