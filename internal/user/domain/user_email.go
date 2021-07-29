package domain

import (
	"errors"
	"ms-auth/pkg/utils"
)

func (u *User) ValidateEmail() error {

	if u.Email == "" {
		return errors.New("Field email is required")
	}

	if len(u.Email) < 3 {
		return errors.New("Email cannot be less than 3")
	}

	if len(u.Email) > 128 {
		return errors.New("Email cannot be longer than 128")
	}

	if utils.ValidateEmail(u.Email) {
		return errors.New("Email is invalid")
	}

	return nil

}
