package domain_test

import (
	"ms-auth/internal/client/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateUser(t *testing.T) {

	t.Run("should be user required", func(t *testing.T) {

		fakeUser := domain.Client{
			UserID: 0,
		}

		err := fakeUser.ValidateUser()

		assert.NotNil(t, err)

	})

	t.Run("should be user validate", func(t *testing.T) {

		fakeUser := domain.Client{
			UserID: faker.RandomUnixTime(),
		}

		err := fakeUser.ValidateUser()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateUser(b *testing.B) {

	b.Run("should be secret validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.Client{
				UserID: faker.RandomUnixTime(),
			}

			err := fakeUser.ValidateUser()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}
