package usecase_test

import (
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_SaveToken(t *testing.T) {

	model := getModel()

	t.Run("should be able invalid user", func(t *testing.T) {

		rUser := &domain.User{}

		token, err := model.SaveToken(rUser)
		assert.Nil(t, token)
		assert.NotNil(t, err)
	})

	t.Run("should be able valid user", func(t *testing.T) {

		rUser := &domain.User{
			ID:        1,
			Username:  faker.Username(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Password:  faker.Password(),
			Email:     faker.Email(),
		}

		token, err := model.SaveToken(rUser)
		assert.NotNil(t, token)
		assert.Nil(t, err)
	})

}
