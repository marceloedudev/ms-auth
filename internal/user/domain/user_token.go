package domain

import (
	"errors"
	"ms-auth/pkg/utils"
	"ms-auth/pkg/uuid"
	"time"
)

func (u *User) CreateToken(atExpiresIn int, rtExpiresIn int) *TokenData {

	accessUUID := uuid.NewUUIDString()
	refreshUUID := uuid.NewUUIDString()
	accessTokenExpire := time.Now().Add(time.Duration(atExpiresIn) * time.Second).Unix()
	refreshTokenExpire := time.Now().Add(time.Duration(rtExpiresIn) * time.Second).Unix()

	token := &TokenData{
		TokenUUID: TokenUUID{
			AccessUUID:  accessUUID,
			RefreshUUID: refreshUUID,
		},
		TokenExpires: TokenExpires{
			AtExpires: accessTokenExpire,
			RtExpires: refreshTokenExpire,
		},
	}

	return token

}

func IsValidToken(token string) error {

	if !!(len(token) > 0 && token != "") {
		return nil
	}

	return errors.New("Token is required")

}

func ExpiresDurationToken(AtExpires int64, RtExpires int64) *ExpireToken {

	currentTime := time.Now()

	timeAccessToken := time.Unix(AtExpires, 0)
	expireTimeAt := timeAccessToken.Sub(currentTime)

	timeRefreshToken := time.Unix(RtExpires, 0)
	expireTimeRt := timeRefreshToken.Sub(currentTime)

	return &ExpireToken{
		ExpireTimeAt: expireTimeAt,
		ExpireTimeRt: expireTimeRt,
	}

}

func (t *UserDataToken) ToJSON() (data []byte, err error) {

	rUser := &User{
		ID:        t.ID,
		Username:  t.Username,
		FirstName: t.FirstName,
		LastName:  t.LastName,
		Email:     t.Email,
	}

	err = rUser.ValidateID()

	if err != nil {
		return nil, err
	}

	err = rUser.ValidateUsername()

	if err != nil {
		return nil, err
	}

	err = rUser.ValidateFirstName()

	if err != nil {
		return nil, err
	}

	err = rUser.ValidateLastName()

	if err != nil {
		return nil, err
	}

	data, err = utils.MarshalJSON(t)

	if err != nil {
		return nil, err
	}

	return data, nil

}
