package usecase

import (
	"time"

	"github.com/pkg/errors"
)

func (s *Service) SaveAccessToken(token, data string, expireTime time.Duration) error {

	err := s.redisRepository.Set(token, data, expireTime)

	if err != nil {
		return errors.Wrap(err, "usecase.Set")
	}

	return nil

}
