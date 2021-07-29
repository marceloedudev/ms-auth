package usecase

import (
	"ms-auth/internal/user/domain"

	"github.com/pkg/errors"
)

func (s *Service) RevokeUser(token, tokenGrantType string) error {

	err := domain.IsValidToken(token)

	if err != nil {
		return errors.Wrap(err, "usecase.IsValidToken")
	}

	grantTypes := map[string]func(token string) (*domain.UserDataToken, error){
		"access_token":  s.RevokeAccessToken,
		"refresh_token": s.RevokeRefreshToken,
	}

	grantHandler, ok := grantTypes[tokenGrantType]

	if !ok {
		return domain.ErrInvalidGrantType
	}

	_, err = grantHandler(token)

	if err != nil {
		return errors.Wrap(err, "usecase.grantHandler")
	}

	return nil

}
