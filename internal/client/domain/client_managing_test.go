package domain_test

import (
	"fmt"
	"ms-auth/internal/client/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateManaging(t *testing.T) {

	t.Run("should be managing required", func(t *testing.T) {

		fakeUser := domain.Client{
			ManagingID: "",
		}

		err := fakeUser.ValidateManaging()

		assert.NotNil(t, err)

	})

	t.Run("should must be short managing", func(t *testing.T) {

		fakeUser := domain.Client{
			ManagingID: fmt.Sprintf("%s%s", faker.Paragraph(), faker.Paragraph()),
		}

		err := fakeUser.ValidateManaging()

		assert.NotNil(t, err)

	})

	t.Run("should be managing validate", func(t *testing.T) {

		fakeUser := domain.Client{
			ManagingID: faker.FirstName(),
		}

		err := fakeUser.ValidateManaging()

		assert.Nil(t, err)

	})

}

func Benchmark_ValidateManaging(b *testing.B) {

	b.Run("should be managing validate", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			fakeUser := domain.Client{
				ManagingID: faker.Word(),
			}

			err := fakeUser.ValidateManaging()

			if err != nil {
				b.Error("unexpected result: ", err)
			}
		}

	})

}
