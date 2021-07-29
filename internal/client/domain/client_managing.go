package domain

import "errors"

func (u *Client) ValidateManaging() error {

	if u.ManagingID == "" {
		return errors.New("Field managing is required")
	}

	if len(u.ManagingID) > 32 {
		return errors.New("Managing cannot be longer than 32")
	}

	return nil

}
