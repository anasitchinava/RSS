package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from the headers of an HTTP request
// Example usage: Authorization ApiKey {value here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authentication info is missing")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("auth header is malformed")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("the first part of auth header is malformed")
	}

	return vals[1], nil
}
