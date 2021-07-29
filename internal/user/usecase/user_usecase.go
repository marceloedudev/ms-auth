package usecase

import (
	"ms-auth/config"
)

type Service struct {
	config          *config.Config
	pgRepository    PgRepository
	redisRepository RedisRepository
}

func NewService(c *config.Config, p PgRepository, r RedisRepository) *Service {
	return &Service{
		config:          c,
		pgRepository:    p,
		redisRepository: r,
	}
}
