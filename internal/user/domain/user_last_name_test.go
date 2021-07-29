package domain_test

import (
	"fmt"
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateLastName(t *testing.T) {

	t.Run("should be last name is required", func(t *testing.T) {

		fakeUser := domain.User{
			LastName: "",
		}

		err := fakeUser.ValidateLastName()

		assert.NotNil(t, err)

	})

	t.Run("should must be long last name", func(t *testing.T) {

		fakeUser := domain.User{
			LastName: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateLastName()

		assert.NotNil(t, err)

	})

	t.Run("should be last name validate", func(t *testing.T) {

		fakeUser := domain.User{
			LastName: faker.LastName(),
		}

		err := fakeUser.ValidateLastName()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateLastName(b *testing.B) {

	b.Run("should be last name validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.User{
				LastName: faker.LastName(),
			}

			err := fakeUser.ValidateLastName()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}
