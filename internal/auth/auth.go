package auth

import (
	"errors"
	"net/http"
	"strings"
)

// get api key extacts an api key
// the headres on an http request
// Example
// Authorization: ApiKey <apikey>
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth key found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("infalid authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("infalid first path of authorization header")
	}
	return vals[1], nil
}
