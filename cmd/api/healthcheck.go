package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a fixed JSON format from a string.
	// Using raw string literals to use double quotes without the need to escape them.
	// Use the %q format specifier to wrap the interpolated values in double quotes.
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Set the "Content-Type: application/json" header on the response.
	// If you forget this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
	// header instead.

	// JSON RFC: JSON text exchanged between systems that are not part of a closed ecosystem MUST be
	// encoded using UTF-8. (hence why we are not using "application/json; charset=utf-8")
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON to the http.ResponseWriter as the response body.
	w.Write([]byte(js))
}
