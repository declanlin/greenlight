package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retrieve the "id" URL parameter from the current request context.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Retrieve the interpolated "id" parameter from the requested URL.
	// ParamsFromContext() retrieves a slice of parameter names and values.
	params := httprouter.ParamsFromContext(r.Context())

	// Get the value of the "id" parameter as a base 10, 64 bit integer.
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

type envelope map[string]interface{}

// Define a helper for sending HTTP responses.
// - w: io.Writer to write the response to
// - status: the HTTP status code to send
// - data: the data to encode to a JSON
// - headers: a header map containing additional HTTP headers to include in the response.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Use the json.MarshalIndent() function so that whitespace is added to the encoded
	// JSON. Here we use no line prefix ("") and tab indents ("\t") for each element.
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append a new line to the JSON body to make it easier to read in terminal application outputs.
	js = append(js, '\n')

	// We won't run into errors before writing the response. It is safe to add any headers that we want to include.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Set the "Content-Type: application/json" header on the response.
	// If you forget this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
	// header instead.

	// JSON RFC: JSON text exchanged between systems that are not part of a closed ecosystem MUST be
	// encoded using UTF-8. (hence why we are not using "application/json; charset=utf-8")
	w.Header().Set("Content-Type", "application/json")

	// Send the response headers with the specified status code.
	w.WriteHeader(status)
	// Send the JSON response body.
	w.Write(js)
	return nil
}
