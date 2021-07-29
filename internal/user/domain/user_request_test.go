package domain_test

import (
	"ms-auth/internal/user/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckMaxRequests(t *testing.T) {

	t.Run("should be valid current attempts", func(t *testing.T) {

		err := domain.CheckMaxRequests(9, 10)

		assert.Nil(t, err)

	})

	t.Run("should be invalid current attempts", func(t *testing.T) {

		err := domain.CheckMaxRequests(11, 10)

		assert.NotNil(t, err)

	})

}

func Test_FormatKeyRequestLimit(t *testing.T) {

	t.Run("should be returns key formatted", func(t *testing.T) {

		key := domain.FormatKeyRequestLimit("test")

		assert.NotEqual(t, key, "")

	})

}
