package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("authorization header not found")
	}
	vals := strings.Split(apiKey, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization header format")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid authorization scheme")
	}
	return vals[1], nil
}
