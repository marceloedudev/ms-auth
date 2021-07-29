package usecase

import (
	"fmt"
	"ms-auth/internal/user/domain"
	"time"

	"github.com/pkg/errors"
)

func (s *Service) RequestLimit(key string, maxAttempts int, timeInterval int) error {

	timeLimit := time.Duration(timeInterval) * time.Second

	currentTime := time.Now().UnixNano()

	maxTime := fmt.Sprint(currentTime - (timeLimit.Nanoseconds()))

	_, err := s.redisRepository.ZRemRangeByScore(key, "0", maxTime)

	if err != nil {

		return errors.Wrap(err, "usecase.ZRemRangeByScore")

	}

	requests, err := s.redisRepository.ZRange(key, 0, -1)

	if err != nil {

		return errors.Wrap(err, "usecase.ZRange")

	}

	err = domain.CheckMaxRequests(len(requests), maxAttempts)

	if err != nil {

		return errors.Wrap(err, "usecase.IsMaxRequests")

	}

	s.redisRepository.ZAddNX(key, float64(currentTime), float64(currentTime))

	s.redisRepository.Expire(key, timeLimit)

	return nil

}
