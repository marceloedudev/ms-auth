package domain_test

import (
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckPassword(t *testing.T) {

	t.Run("should be invalid password", func(t *testing.T) {

		rUserBD := &domain.User{
			Password: "123456",
		}

		err := rUserBD.ValidatePassword()

		rUserNew := &domain.User{
			Password: "xyz",
		}

		err = rUserNew.CheckPassword(rUserBD.Password)

		assert.NotNil(t, err)

	})

	t.Run("should be valid password", func(t *testing.T) {

		rUserBD := &domain.User{
			Password: "123456",
		}

		err := rUserBD.ValidatePassword()

		rUserNew := &domain.User{
			Password: "123456",
		}

		err = rUserNew.CheckPassword(rUserBD.Password)

		assert.Nil(t, err)

	})

}
