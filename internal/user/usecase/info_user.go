package usecase

import (
	"context"
	"ms-auth/internal/user/domain"

	"github.com/pkg/errors"
)

func (s *Service) InfoUser(ctx context.Context, token string) (user *domain.User, err error) {

	userDecoded, err := s.InfoAccessToken(token)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.InfoAccessToken")
	}

	user, err = s.pgRepository.FindByID(ctx, userDecoded.ID)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.FindByID")
	}


	return user, nil

}
