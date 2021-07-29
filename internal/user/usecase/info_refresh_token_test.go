package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_InfoRefreshToken(t *testing.T) {

	model := getModel()

	t.Run("should be able info refresh token user", func(t *testing.T) {
		user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.NotNil(t, user)
		assert.Nil(t, err)

		newToken, err := model.InfoRefreshToken(user.RefreshToken)
		assert.NotEqual(t, user.RefreshToken, newToken)
		assert.Nil(t, err)

	})

	t.Run("should be able invalid info refresh token", func(t *testing.T) {
		token, err := model.InfoRefreshToken(faker.UUIDDigit())
		assert.Nil(t, token)
		assert.NotNil(t, err)
	})

}
