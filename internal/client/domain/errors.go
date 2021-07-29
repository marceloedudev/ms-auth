package domain

import "errors"

var (
	ErrIDOrNameAlreadyUsed = errors.New("ID/Name is already used")
)
