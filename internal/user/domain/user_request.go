package domain

import (
	"errors"
	"fmt"
)

func CheckMaxRequests(currentAttempts, maxAttempts int) error {

	if currentAttempts >= maxAttempts {

		return errors.New("You made a lot of requests in a short time")

	}

	return nil

}

func FormatKeyRequestLimit(key string) string {

	return fmt.Sprint("user:limit:", key)

}
