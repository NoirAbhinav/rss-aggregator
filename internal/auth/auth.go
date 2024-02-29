package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no api key found ")
	}
	vals := strings.Split(val, "=")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "apikey" {
		return "", errors.New("malformed first part of header")
	}
	if vals[1] == "" {
		return "", errors.New("no api key sent in header")
	}
	return vals[1], nil
}
