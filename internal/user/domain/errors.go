package domain

import "errors"

var (
	ErrUnknownError               = errors.New("Unknown error")
	ErrInvalidGrantType           = errors.New("missing or unknown grant type")
	ErrInvalidToken               = errors.New("token contains an invalid number of segments")
	ErrUserNotFound               = errors.New("User not found")
	ErrEmailOrUsernameAlreadyUsed = errors.New("Email or username is already used")
)
