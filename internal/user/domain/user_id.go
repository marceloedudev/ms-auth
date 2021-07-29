package domain

import (
	"errors"
	"ms-auth/pkg/uuid"
)

func (u *User) ValidateID() error {
	if u.ID == 0 {
		return errors.New("Invalid ID")
	}

	return nil
}

func (u *User) ValidateUUID() error {
	u.UUID = uuid.NewUUIDString()

	return nil
}
