package data

import (
	"fmt"
	"strconv"
)

// Declare a custom Runtime type which has the same underlying type as the Runtime field in the Movie struct.
type Runtime int32

// Implement a MarshalJSON() method on the Runtime type so that it satisfies the
// json.Marshaler interface. This should return the JSON-encoded value for the movie
// runtime (in our case, it will return a string in the format "<runtime> mins").
func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a formatted string containing the movie runtime.
	jsonValue := fmt.Sprintf("%d mins", r)

	// Wrap the string in double quotes in order to make it a valid JSON string.
	quotedJSONValue := strconv.Quote(jsonValue)

	// Convert the quoted string to a byte slice and return it
	return []byte(quotedJSONValue), nil
}
