package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_RevokeUser(t *testing.T) {

	model := getModel()

	t.Run("should be able invalid token", func(t *testing.T) {
		err := model.RevokeUser("", "")
		assert.NotNil(t, err)
	})

	t.Run("should be able invalid grant type", func(t *testing.T) {
		err := model.RevokeUser(faker.UUIDDigit(), "")
		assert.NotNil(t, err)
	})

	t.Run("should be able invalid grant hanlder", func(t *testing.T) {
		err := model.RevokeUser(faker.UUIDDigit(), "access_token")
		assert.NotNil(t, err)
	})

	user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
	assert.NotNil(t, user)
	assert.Nil(t, err)

	t.Run("should be able revoke access token", func(t *testing.T) {
		err := model.RevokeUser(user.AccessToken, "access_token")
		assert.Nil(t, err)
	})

	t.Run("should be able revoke refresh token", func(t *testing.T) {
		err := model.RevokeUser(user.RefreshToken, "refresh_token")
		assert.Nil(t, err)
	})

}
