package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_InfoAccessToken(t *testing.T) {

	model := getModel()

	t.Run("should be able info access token user", func(t *testing.T) {
		user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.NotNil(t, user)
		assert.Nil(t, err)

		newToken, err := model.InfoAccessToken(user.AccessToken)
		assert.NotEqual(t, user.AccessToken, newToken)
		assert.Nil(t, err)

	})

	t.Run("should be able invalid info access token", func(t *testing.T) {
		token, err := model.InfoAccessToken(faker.UUIDDigit())
		assert.Nil(t, token)
		assert.NotNil(t, err)
	})

}
