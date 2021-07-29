package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_RequestLimit(t *testing.T) {

	model := getModel()

	t.Run("should be able not return errors", func(t *testing.T) {
		err := model.RequestLimit("test", 1, 5)
		assert.Nil(t, err)
	})

	t.Run("should be able block spam", func(t *testing.T) {
		key := faker.UUIDDigit()

		err := model.RequestLimit(key, 1, 5)
		assert.Nil(t, err)

		err = model.RequestLimit(key, 1, 5)
		assert.NotNil(t, err)
	})

	t.Run("should be able not block spam", func(t *testing.T) {
		err := model.RequestLimit("test1", 1, 5)
		assert.Nil(t, err)

		err = model.RequestLimit("test2", 1, 5)
		assert.Nil(t, err)
	})

}
