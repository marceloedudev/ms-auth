package usecase

import (
	"context"
	"ms-auth/internal/user/domain"

	"github.com/pkg/errors"
)

var (
	TokenTypes = "Bearer"
)

func (s *Service) LoginUser(ctx context.Context, username, password string) (info *domain.TokenResponse, err error) {

	rUser := &domain.User{
		Username: username,
		Password: password,
	}

	err = rUser.ValidateUsernameOrEmail()

	if err != nil {
		return nil, errors.Wrap(err, "usecase.ValidateUsernameOrEmail")
	}

	user, err := s.pgRepository.FindByUsernameOrEmail(ctx, username, username)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.FindByUsernameOrEmai")
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	err = rUser.CheckPassword(user.Password)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.CheckPassword")
	}

	token, err := s.SaveToken(user)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.SaveToken")
	}

	return token, nil

}
