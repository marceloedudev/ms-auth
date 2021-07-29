package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_RevokeAccessToken(t *testing.T) {

	model := getModel()

	t.Run("should be able revoke access token", func(t *testing.T) {
		user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.NotNil(t, user)
		assert.Nil(t, err)

		token, err := model.RevokeAccessToken(user.AccessToken)
		assert.NotNil(t, token)
		assert.Nil(t, err)
	})

	t.Run("should be able invalid revoke access token", func(t *testing.T) {
		token, err := model.RevokeAccessToken(faker.UUIDDigit())
		assert.Nil(t, token)
		assert.NotNil(t, err)
	})

}

func Benchmark_RevokeAccessToken(b *testing.B) {

	b.StopTimer()

	model := getModel()

	b.StartTimer()

	b.Run("should be revoke access token benchmark", func(b *testing.B) {

		for n := 0; n < b.N; n++ {

			user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())

			if err != nil {
				b.Error("unexpected result: ", err)
			}

			_, err = model.RevokeAccessToken(user.AccessToken)

			if err != nil {
				b.Error("unexpected result: ", err)
			}

		}

	})

}
