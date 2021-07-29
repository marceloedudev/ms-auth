package domain_test

import (
	"fmt"
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateUsername(t *testing.T) {

	t.Run("should be username required", func(t *testing.T) {

		fakeUser := domain.User{
			Username: "",
		}

		err := fakeUser.ValidateUsername()

		assert.NotNil(t, err)

	})

	t.Run("should must be short username", func(t *testing.T) {

		fakeUser := domain.User{
			Username: "abc",
		}

		err := fakeUser.ValidateUsername()

		assert.NotNil(t, err)

	})

	t.Run("should must be long username", func(t *testing.T) {

		fakeUser := domain.User{
			Username: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateUsername()

		assert.NotNil(t, err)

	})

	t.Run("should must be invalid username", func(t *testing.T) {

		fakeUser := domain.User{
			Username: "username#;",
		}

		err := fakeUser.ValidateUsername()

		assert.NotNil(t, err)

	})

	t.Run("should be username validate", func(t *testing.T) {

		fakeUser := domain.User{
			Username: faker.Username(),
		}

		err := fakeUser.ValidateUsername()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateUsername(b *testing.B) {

	b.Run("should be username validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.User{
				Username: faker.Username(),
			}

			err := fakeUser.ValidateUsername()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}

func Test_ValidateUsernameOrEmail(t *testing.T) {

	t.Run("should be username/email required", func(t *testing.T) {

		fakeUser := domain.User{
			Username: "",
		}

		err := fakeUser.ValidateUsernameOrEmail()

		assert.NotNil(t, err)

	})

	t.Run("should must be short username/email", func(t *testing.T) {

		fakeUser := domain.User{
			Username: "ab",
		}

		err := fakeUser.ValidateUsernameOrEmail()

		assert.NotNil(t, err)

	})

	t.Run("should must be long username/email", func(t *testing.T) {

		fakeUser := domain.User{
			Username: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateUsernameOrEmail()

		assert.NotNil(t, err)

	})

	t.Run("should be username/email validate", func(t *testing.T) {

		fakeUser := domain.User{
			Username: faker.Username(),
		}

		err := fakeUser.ValidateUsernameOrEmail()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateUsernameOrEmail(b *testing.B) {

	b.Run("should be username validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.User{
				Username: faker.Username(),
			}

			err := fakeUser.ValidateUsernameOrEmail()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}
