package domain

import (
	"errors"
	"ms-auth/pkg/utils"
)

func (u *User) ValidateUsername() error {

	if u.Username == "" {
		return errors.New("Field username is required")
	}

	if len(u.Username) < 4 {
		return errors.New("Username cannot be less than 4")
	}

	if len(u.Username) > 30 {
		return errors.New("Username cannot be longer than 30")
	}

	if !utils.IsStringAlphanumeric(u.Username) {
		return errors.New("Username must be alphanumeric or special characters only")
	}

	return nil

}

func (u *User) ValidateUsernameOrEmail() error {

	if u.Username == "" {
		return errors.New("Field username is required")
	}

	if len(u.Username) < 3 {
		return errors.New("Username cannot be less than 3")
	}

	if len(u.Username) > 128 {
		return errors.New("Username cannot be longer than 128")
	}

	return nil

}
