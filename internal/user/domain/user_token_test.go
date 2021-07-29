package domain_test

import (
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_CreateToken(t *testing.T) {

	t.Run("should be valid create token", func(t *testing.T) {

		var user domain.User

		data := user.CreateToken(1000, 5000)

		assert.NotNil(t, data)

	})

}

func Test_IsValidToken(t *testing.T) {

	t.Run("should be token required", func(t *testing.T) {

		err := domain.IsValidToken("")

		assert.NotNil(t, err)

	})

	t.Run("should be validate", func(t *testing.T) {

		err := domain.IsValidToken(faker.UUIDDigit())

		assert.Nil(t, err)

	})

}

func Test_ExpiresDurationToken(t *testing.T) {

	t.Run("should be valid token expires", func(t *testing.T) {

		data := domain.ExpiresDurationToken(1000, 5000)

		assert.NotNil(t, data)

	})

}
