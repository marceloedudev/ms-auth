package domain_test

import (
	"fmt"
	"ms-auth/internal/client/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateAddClient(t *testing.T) {

	t.Run("should be add new client required", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID:     "",
			ClientSecret: "",
			UserID:       0,
			ManagingID:   "",
		}

		err := fakeUser.ValidateAddClient()

		assert.NotNil(t, err)

	})

	t.Run("should must be short add new client", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID:     fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
			ClientSecret: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
			UserID:       0,
			ManagingID:   fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateAddClient()

		assert.NotNil(t, err)

	})

	t.Run("should must be invalid secret", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID:     faker.FirstName(),
			ClientSecret: "",
			UserID:       faker.RandomUnixTime(),
			ManagingID:   faker.FirstName(),
		}

		err := fakeUser.ValidateAddClient()

		assert.NotNil(t, err)

	})

	t.Run("should must be invalid userid", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID:     faker.FirstName(),
			ClientSecret: faker.FirstName(),
			UserID:       0,
			ManagingID:   faker.FirstName(),
		}

		err := fakeUser.ValidateAddClient()

		assert.NotNil(t, err)

	})

	t.Run("should must be invalid managing", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID:     faker.FirstName(),
			ClientSecret: faker.FirstName(),
			UserID:       faker.RandomUnixTime(),
			ManagingID:   "",
		}

		err := fakeUser.ValidateAddClient()

		assert.NotNil(t, err)

	})

	t.Run("should be add new client validate", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID:     faker.FirstName(),
			ClientSecret: faker.FirstName(),
			UserID:       faker.RandomUnixTime(),
			ManagingID:   faker.FirstName(),
		}

		err := fakeUser.ValidateAddClient()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateAddClient(b *testing.B) {

	b.Run("should be add new client validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.Client{
				ClientID:     faker.Word(),
				ClientSecret: faker.Jwt(),
				UserID:       faker.RandomUnixTime(),
				ManagingID:   faker.Word(),
			}

			err := fakeUser.ValidateAddClient()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}
