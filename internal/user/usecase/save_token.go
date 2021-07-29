package usecase

import (
	"ms-auth/internal/user/domain"

	"github.com/pkg/errors"
)

func (s *Service) SaveToken(user *domain.User) (*domain.TokenResponse, error) {

	rUser := &domain.UserDataToken{
		ID:         user.ID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		ManagingID: user.ManagingID,
	}

	token := user.CreateToken(s.config.Auth.AccessTokenLifeTime, s.config.Auth.RefreshTokenLifeTime)

	data, err := rUser.ToJSON()

	if err != nil {
		return nil, errors.Wrap(err, "usecase.ToJSON")
	}

	expireTime := domain.ExpiresDurationToken(token.AtExpires, token.RtExpires)

	err = s.SaveAccessToken(token.AccessUUID, string(data), expireTime.ExpireTimeAt)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.SaveAccessToken")
	}

	err = s.SaveRefreshToken(token.RefreshUUID, string(data), expireTime.ExpireTimeRt)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.SaveRefreshToken")
	}

	return &domain.TokenResponse{
		UserID:       user.ID,
		AccessToken:  token.AccessUUID,
		ExpiresIn:    s.config.Auth.AccessTokenLifeTime,
		TokenType:    TokenTypes,
		RefreshToken: token.RefreshUUID,
	}, nil

}
