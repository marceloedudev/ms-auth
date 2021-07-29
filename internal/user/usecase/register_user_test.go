package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {

	model := getModel()

	t.Run("should be able create user", func(t *testing.T) {

		user, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.NotNil(t, user)
		assert.Nil(t, err)

	})

	t.Run("should be able user already used", func(t *testing.T) {

		userName := faker.Username()

		user, err := model.RegisterUser(context.Background(), userName, faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.NotNil(t, user)
		assert.Nil(t, err)

		user, err = model.RegisterUser(context.Background(), userName, faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
		assert.Nil(t, user)
		assert.NotNil(t, err)

	})

	t.Run("should be able error invalid data", func(t *testing.T) {

		user, err := model.RegisterUser(context.Background(), faker.Username(), "", "", faker.Password(), "")
		assert.Nil(t, user)
		assert.NotNil(t, err)

	})

}

func Benchmark_CreateUser(b *testing.B) {

	b.StopTimer()

	model := getModel()

	b.StartTimer()

	b.Run("should be create user benchmark", func(b *testing.B) {

		for n := 0; n < b.N; n++ {

			_, err := model.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())

			if err != nil {
				b.Error("unexpected result: ", err)
				b.Fatal(err)
			}

		}

	})

}
