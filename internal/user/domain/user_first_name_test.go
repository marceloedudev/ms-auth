package domain_test

import (
	"fmt"
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateFirstName(t *testing.T) {

	t.Run("should be first name required", func(t *testing.T) {

		fakeUser := domain.User{
			FirstName: "",
		}

		err := fakeUser.ValidateFirstName()

		assert.NotNil(t, err)

	})

	t.Run("should must be long first name", func(t *testing.T) {

		fakeUser := domain.User{
			FirstName: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateFirstName()

		assert.NotNil(t, err)

	})

	t.Run("should be first name validate", func(t *testing.T) {

		fakeUser := domain.User{
			FirstName: faker.FirstName(),
		}

		err := fakeUser.ValidateFirstName()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateFirstName(b *testing.B) {

	b.Run("should be first name validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.User{
				FirstName: faker.FirstName(),
			}

			err := fakeUser.ValidateFirstName()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}
