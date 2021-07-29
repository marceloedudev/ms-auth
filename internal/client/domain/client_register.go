package domain

func (u *Client) ValidateAddClient() error {

	err := u.ValidateClientID()

	if err != nil {
		return err
	}

	err = u.ValidateSecret()

	if err != nil {
		return err
	}

	err = u.ValidateUser()

	if err != nil {
		return err
	}

	err = u.ValidateManaging()

	if err != nil {
		return err
	}

	return nil

}
