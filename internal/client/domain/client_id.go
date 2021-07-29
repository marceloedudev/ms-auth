package domain

import "errors"

func (u *Client) ValidateClientID() error {

	if u.ClientID == "" {
		return errors.New("Field client id is required")
	}

	if len(u.ClientID) > 32 {
		return errors.New("Client id cannot be longer than 32")
	}

	return nil

}
