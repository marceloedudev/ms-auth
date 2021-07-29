package usecase_test

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_SaveAccessToken(t *testing.T) {

	model := getModel()

	t.Run("should be able valid token", func(t *testing.T) {
		err := model.SaveAccessToken(faker.UUIDDigit(), "{}", time.Duration(time.Now().Unix()))
		assert.Nil(t, err)
	})

}
