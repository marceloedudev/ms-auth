package domain_test

import (
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateID(t *testing.T) {

	t.Run("should be id required", func(t *testing.T) {

		fakeUser := domain.User{
			ID: 0,
		}

		err := fakeUser.ValidateID()

		assert.NotNil(t, err)

	})

	t.Run("should be id validate", func(t *testing.T) {

		fakeUser := domain.User{
			ID: 10,
		}

		err := fakeUser.ValidateID()

		assert.Nil(t, err)

	})

}

func Test_ValidateUUID(t *testing.T) {

	t.Run("should be uuid validate", func(t *testing.T) {

		fakeUser := domain.User{
			UUID: faker.UUIDDigit(),
		}

		err := fakeUser.ValidateUUID()

		assert.Nil(t, err)

	})

}
