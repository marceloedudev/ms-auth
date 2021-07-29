package usecase

import (
	"context"
	"ms-auth/internal/user/domain"
	"time"

	"github.com/pkg/errors"
)

func (s *Service) RegisterUser(ctx context.Context, username, first_name, last_name, password, email string) (info *domain.TokenResponse, err error) {

	rUser := &domain.User{
		Username:      username,
		FirstName:     first_name,
		LastName:      last_name,
		Password:      password,
		Email:         email,
		EmailVerified: false,
		Enabled:       true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	existsUser, err := s.pgRepository.FindByUsernameOrEmail(ctx, username, email)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.FindByUsernameOrEmail")
	}

	if existsUser != nil {
		return nil, domain.ErrEmailOrUsernameAlreadyUsed
	}

	err = rUser.ValidateAddUser()

	if err != nil {
		return nil, errors.Wrap(err, "usecase.ValidateAddUser")
	}

	user, err := s.pgRepository.Create(ctx, rUser)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.Create")
	}

	token, err := s.SaveToken(user)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.SaveToken")
	}

	return token, nil

}
