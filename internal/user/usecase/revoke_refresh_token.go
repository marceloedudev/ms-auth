package usecase

import (
	"ms-auth/internal/user/domain"

	"github.com/pkg/errors"
)

func (s *Service) RevokeRefreshToken(token string) (*domain.UserDataToken, error) {
	user, err := s.RevokeToken(token)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.RevokeToken")
	}

	return user, nil
}
