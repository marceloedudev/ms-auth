package usecase

import (
	"ms-auth/internal/user/domain"

	"github.com/pkg/errors"
)

func (s *Service) RevokeToken(token string) (*domain.UserDataToken, error) {

	user, err := s.InfoToken(token)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.InfoToken")
	}

	err = s.redisRepository.Del(token)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.Del")
	}

	return user, nil
}
