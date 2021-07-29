package usecase

import (
	"ms-auth/internal/user/domain"

	"github.com/pkg/errors"
)

func (s *Service) InfoAccessToken(token string) (user *domain.UserDataToken, err error) {

	user, err = s.InfoToken(token)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.InfoToken")
	}

	return user, nil

}
