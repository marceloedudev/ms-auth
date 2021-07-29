package usecase

import (
	"ms-auth/config"
)

type Service struct {
	config       *config.Config
	pgRepository PgRepository
}

func NewService(c *config.Config, p PgRepository) *Service {
	return &Service{
		config:       c,
		pgRepository: p,
	}
}
