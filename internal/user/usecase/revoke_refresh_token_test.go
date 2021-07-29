package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_RevokeRefreshToken(t *testing.T) {

	model := getModel()

	t.Run("should be able revoke refresh token", func(t *testing.T) {
		user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.NotNil(t, user)
		assert.Nil(t, err)

		token, err := model.RevokeRefreshToken(user.RefreshToken)
		assert.NotNil(t, token)
		assert.Nil(t, err)
	})

	t.Run("should be able invalid revoke refresh token", func(t *testing.T) {
		token, err := model.RevokeRefreshToken(faker.UUIDDigit())
		assert.Nil(t, token)
		assert.NotNil(t, err)
	})

}

func Benchmark_RevokeRefreshToken(b *testing.B) {

	b.StopTimer()

	model := getModel()

	b.StartTimer()

	b.Run("should be revoke refresh token benchmark", func(b *testing.B) {

		for n := 0; n < b.N; n++ {

			user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())

			if err != nil {
				b.Error("unexpected result: ", err)
			}

			_, err = model.RevokeRefreshToken(user.RefreshToken)

			if err != nil {
				b.Error("unexpected result: ", err)
			}

		}

	})

}
