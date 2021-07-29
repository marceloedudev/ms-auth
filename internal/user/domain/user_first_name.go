package domain

import "errors"

func (u *User) ValidateFirstName() error {

	if u.FirstName == "" {
		return errors.New("Field firstname is required")
	}

	if len(u.FirstName) > 30 {
		return errors.New("First name cannot be longer than 30")
	}

	return nil

}
