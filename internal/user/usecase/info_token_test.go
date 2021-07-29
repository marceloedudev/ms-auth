package usecase_test

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_InfoToken(t *testing.T) {

	model := getModel()

	t.Run("should be able invalid token", func(t *testing.T) {
		user, err := model.InfoToken("xyz")
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("should be able invalid data storage token", func(t *testing.T) {
		token := faker.UUIDDigit()

		expiresIn := time.Now().Add(time.Duration(2) * time.Minute).Unix()

		err := model.SaveAccessToken(token, "{xyz}", time.Unix(expiresIn, 0).Sub(time.Now()))
		assert.Nil(t, err)

		user, err := model.InfoToken(token)
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

}
