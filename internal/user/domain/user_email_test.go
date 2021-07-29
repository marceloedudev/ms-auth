package domain_test

import (
	"fmt"
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateEmail(t *testing.T) {

	t.Run("should be email required", func(t *testing.T) {

		fakeUser := domain.User{
			Email: "",
		}

		err := fakeUser.ValidateEmail()

		assert.NotNil(t, err)

	})

	t.Run("should must be short email", func(t *testing.T) {

		fakeUser := domain.User{
			Email: "x",
		}

		err := fakeUser.ValidateEmail()

		assert.NotNil(t, err)

	})

	t.Run("should must be long email", func(t *testing.T) {

		fakeUser := domain.User{
			Email: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateEmail()

		assert.NotNil(t, err)

	})

	t.Run("should must be invalid email", func(t *testing.T) {

		fakeUser := domain.User{
			Email: "example-gmail.com",
		}

		err := fakeUser.ValidateEmail()

		assert.NotNil(t, err)

	})

	t.Run("should be email validate", func(t *testing.T) {

		fakeUser := domain.User{
			Email: faker.Email(),
		}

		err := fakeUser.ValidateEmail()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateEmail(b *testing.B) {

	b.Run("should be email benchmark", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.User{
				Email: faker.Email(),
			}

			err := fakeUser.ValidateEmail()

			if err != nil {
				b.Error("Unexpected result: ", err)
			}
		}

	})

}
