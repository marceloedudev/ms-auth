package domain

import "errors"

func (u *User) ValidateLastName() error {

	if u.LastName == "" {
		return errors.New("Field lastname is required")
	}

	if len(u.LastName) > 30 {
		return errors.New("Last name cannot be longer than 30")
	}

	return nil

}
