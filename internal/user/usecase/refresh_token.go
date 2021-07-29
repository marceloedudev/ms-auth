package usecase

import (
	"context"
	"ms-auth/internal/user/domain"

	"github.com/pkg/errors"
)

func (s *Service) RefreshUserToken(ctx context.Context, refreshToken string) (info *domain.TokenResponse, err error) {

	err = domain.IsValidToken(refreshToken)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.RefreshUserToken")
	}

	tokenDecoded, err := s.RevokeRefreshToken(refreshToken)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.RevokeRefreshToken")
	}

	user, err := s.pgRepository.FindByID(ctx, tokenDecoded.ID)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.FindByID")
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	token, err := s.SaveToken(user)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.SaveToken")
	}

	return token, nil

}
