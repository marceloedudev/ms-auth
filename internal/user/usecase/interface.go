package usecase

import (
	"context"
	"ms-auth/internal/user/domain"
	"time"
)

type PgRepository interface {
	Create(ctx context.Context, e *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, userID int64) (*domain.User, error)
	FindByUsernameOrEmail(ctx context.Context, username, email string) (*domain.User, error)
	Update(ctx context.Context, userID int64, u *domain.User) error
	Destroy(ctx context.Context, userID int64) error
}

type RedisRepository interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Del(key string) error
	ZRemRangeByScore(key, min, max string) (int64, error)
	ZRange(key string, start, stop int64) ([]string, error)
	Expire(key string, expiration time.Duration) bool
	ZAddNX(key string, score float64, member interface{}) int64
}

type UseCaseUser interface {
	LoginUser(ctx context.Context, username, password string) (*domain.TokenResponse, error)
	SaveToken(user *domain.User) (*domain.TokenResponse, error)
	SaveAccessToken(token, data string, expireTime time.Duration) error
	SaveRefreshToken(token, data string, expireTime time.Duration) error
	RevokeUser(token, tokenGrantType string) error
	RevokeToken(token string) (*domain.UserDataToken, error)
	RevokeAccessToken(token string) (*domain.UserDataToken, error)
	RevokeRefreshToken(token string) (*domain.UserDataToken, error)
	RefreshUserToken(ctx context.Context, refreshToken string) (*domain.TokenResponse, error)
	InfoToken(token string) (*domain.UserDataToken, error)
	InfoAccessToken(token string) (*domain.UserDataToken, error)
	InfoRefreshToken(token string) (*domain.UserDataToken, error)
	RegisterUser(ctx context.Context, username, first_name, last_name, password, email string) (*domain.TokenResponse, error)
	RequestLimit(key string, maxAttempts int, timeInterval int) error
	InfoUser(ctx context.Context, token string) (*domain.User, error)
}
