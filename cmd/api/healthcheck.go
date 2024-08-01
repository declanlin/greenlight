package main

import (
	"encoding/json"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// Create a map which holds the information we want to send in the response.
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// Pass the map to the json.Marshal() function.
	// json.Marshal() takes a native Go object and converts it to a byte slice containing the JSON data.
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	js = append(js, '\n')

	// Set the "Content-Type: application/json" header on the response.
	// If you forget this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
	// header instead.

	// JSON RFC: JSON text exchanged between systems that are not part of a closed ecosystem MUST be
	// encoded using UTF-8. (hence why we are not using "application/json; charset=utf-8")
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON to the http.ResponseWriter as the response body.
	w.Write(js)
}
