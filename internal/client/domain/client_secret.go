package domain

import "errors"

// ValidateSecret validator
func (u *Client) ValidateSecret() error {

	if u.ClientSecret == "" {
		return errors.New("Field secret is required")
	}

	if len(u.ClientSecret) > 128 {
		return errors.New("Secret cannot be longer than 128")
	}

	return nil

}
