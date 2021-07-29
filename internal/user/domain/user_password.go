package domain

import (
	"errors"
	"ms-auth/pkg/hash"
)

func (u *User) CheckPassword(password string) error {

	err := hash.CompareHashAndPassword([]byte(password), []byte(u.Password))

	if err != nil {
		return errors.New("Incorrect password")
	}

	return nil

}

func (u *User) ValidatePassword() error {
	password, err := hash.GenerateHashPassword(u.Password)

	if err != nil {
		return nil
	}

	u.Password = password

	return nil
}
