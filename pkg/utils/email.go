package utils

import "regexp"

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`.+@[^@]+\.[^@]{2,}$`)
	return !Re.MatchString(email)
}
