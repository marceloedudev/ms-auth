package domain_test

import (
	"fmt"
	"ms-auth/internal/client/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateClientID(t *testing.T) {

	t.Run("should be client id required", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID: "",
		}

		err := fakeUser.ValidateClientID()

		assert.NotNil(t, err)

	})

	t.Run("should must be short client id", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateClientID()

		assert.NotNil(t, err)

	})

	t.Run("should be client id validate", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientID: faker.FirstName(),
		}

		err := fakeUser.ValidateClientID()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateClientID(b *testing.B) {

	b.Run("should be client id validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.Client{
				ClientID: faker.Word(),
			}

			err := fakeUser.ValidateClientID()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}
