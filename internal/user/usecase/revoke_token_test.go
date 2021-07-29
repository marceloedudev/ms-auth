package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_RevokeToken(t *testing.T) {

	model := getModel()

	t.Run("should be able revoke token", func(t *testing.T) {
		user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.NotNil(t, user)
		assert.Nil(t, err)

		token, err := model.RevokeToken(user.RefreshToken)
		assert.NotNil(t, token)
		assert.Nil(t, err)
	})

	t.Run("should be able invalid revoke token", func(t *testing.T) {
		token, err := model.RevokeToken(faker.UUIDDigit())
		assert.Nil(t, token)
		assert.NotNil(t, err)
	})

}
