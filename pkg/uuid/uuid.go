package uuid

import uuid "github.com/satori/go.uuid"

func NewUUID() uuid.UUID {
	return uuid.NewV4()
}

func NewUUIDString() string {
	return NewUUID().String()
}
