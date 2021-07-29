package domain

func (u *User) ValidateAddUser() error {

	err := u.ValidateUUID()

	if err != nil {
		return err
	}

	err = u.ValidateUsername()

	if err != nil {
		return err
	}

	err = u.ValidateFirstName()

	if err != nil {
		return err
	}

	err = u.ValidateLastName()

	if err != nil {
		return err
	}

	err = u.ValidatePassword()

	if err != nil {
		return err
	}

	err = u.ValidateEmail()

	if err != nil {
		return err
	}

	return nil

}
