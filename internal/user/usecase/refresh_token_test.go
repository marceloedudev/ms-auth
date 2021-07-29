package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_RefreshUserToken(t *testing.T) {

	model := getModel()

	t.Run("should be able refresh token user", func(t *testing.T) {
		user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.NotNil(t, user)
		assert.Nil(t, err)

		newToken, err := model.RefreshUserToken(context.Background(), user.RefreshToken)
		assert.NotEqual(t, user.RefreshToken, newToken)
		assert.Nil(t, err)
	})

	t.Run("should be able invalid token", func(t *testing.T) {
		newToken, err := model.RefreshUserToken(context.Background(), "")
		assert.Nil(t, newToken)
		assert.NotNil(t, err)
	})

	t.Run("should be able not founds token", func(t *testing.T) {
		newToken, err := model.RefreshUserToken(context.Background(), faker.UUIDDigit())
		assert.Nil(t, newToken)
		assert.NotNil(t, err)
	})

}
