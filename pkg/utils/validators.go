package utils

import (
	"fmt"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

var (
	IsStringAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString
)

func ValidatorError(err error) string {

	defaultMessage := "Internal server error"

	if err == nil {
		return defaultMessage
	}

	var errorMessage string

	if os.Getenv("config") == "production" {

		if reason := errors.Cause(err); reason != nil {

			errorMessage = fmt.Sprint(reason)

		}

	} else {

		if reason := errors.Unwrap(err); reason != nil {

			errorMessage = fmt.Sprint(reason)

		}

	}

	if errorMessage == "" {
		defaultMessage = fmt.Sprint(err)
	} else {
		defaultMessage = errorMessage
	}

	return fmt.Sprint(defaultMessage)

}
