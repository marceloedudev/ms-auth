package usecase

import (
	"ms-auth/internal/user/domain"
	"ms-auth/pkg/utils"

	"github.com/pkg/errors"
)

func (s *Service) InfoToken(token string) (user *domain.UserDataToken, err error) {

	data, err := s.redisRepository.Get(token)

	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	err = utils.UnmarshalJSON([]byte(data), &user)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.UnmarshalJSON")
	}

	return user, nil

}
