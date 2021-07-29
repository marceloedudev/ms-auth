package usecase_test

import (
	"context"
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_LoginUser(t *testing.T) {

	model := getModel()

	rUser := &domain.User{
		Username: faker.Username(),
		Password: faker.Password(),
	}

	user, err := model.RegisterUser(context.Background(), rUser.Username, faker.FirstName(), faker.LastName(), rUser.Password, faker.Email())
	assert.NotNil(t, user)
	assert.Nil(t, err)

	t.Run("should be able login user", func(t *testing.T) {

		info, err := model.LoginUser(context.Background(), rUser.Username, rUser.Password)
		assert.NotNil(t, info)
		assert.Nil(t, err)

	})

	t.Run("should be able invalid username/email login user", func(t *testing.T) {

		info, err := model.LoginUser(context.Background(), "_", rUser.Password)
		assert.Nil(t, info)
		assert.NotNil(t, err)

	})

	t.Run("should be able not found email login user", func(t *testing.T) {

		info, err := model.LoginUser(context.Background(), faker.Email(), faker.Password())

		assert.Nil(t, info)
		assert.NotNil(t, err)

	})

	t.Run("should be able invalid password login user", func(t *testing.T) {

		info, err := model.LoginUser(context.Background(), rUser.Username, faker.Password())

		assert.Nil(t, info)
		assert.NotNil(t, err)

	})

}
