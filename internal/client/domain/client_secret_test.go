package domain_test

import (
	"fmt"
	"ms-auth/internal/client/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateSecret(t *testing.T) {

	t.Run("should be secret required", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientSecret: "",
		}

		err := fakeUser.ValidateSecret()

		assert.NotNil(t, err)

	})

	t.Run("should must be short secret", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientSecret: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateSecret()

		assert.NotNil(t, err)

	})

	t.Run("should be secret validate", func(t *testing.T) {

		fakeUser := domain.Client{
			ClientSecret: faker.FirstName(),
		}

		err := fakeUser.ValidateSecret()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateSecret(b *testing.B) {

	b.Run("should be secret validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.Client{
				ClientSecret: faker.Jwt(),
			}

			err := fakeUser.ValidateSecret()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}
