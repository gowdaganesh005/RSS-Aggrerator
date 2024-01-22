package auth

import (
	"errors"
	"net/http"
	"strings"
)

// authoriztion : api {apikey here}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("authorization")
	if val == "" {
		return "", errors.New("no authentication info")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication")
	}
	if vals[0] != "api" {
		return "", errors.New("malformed first part of authentication")
	}
	return vals[1], nil

}
