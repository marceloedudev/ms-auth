package utils

import (
	"net/http"
	"strings"
)

func ExtractTokenFromRequest(r *http.Request) string {

	bearerToken := r.Header.Get("Authorization")

	stringArray := strings.Split(bearerToken, " ")

	if len(stringArray) == 2 {

		return stringArray[1]

	}

	return ""

}
