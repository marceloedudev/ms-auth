package domain

import "errors"

func (u *Client) ValidateUser() error {

	if u.UserID == 0 {
		return errors.New("Field user is required")
	}

	return nil

}
