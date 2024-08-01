package main

import (
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
