package usecase

import (
	"context"
	"ms-auth/internal/client/domain"
)

type PgRepository interface {
	Create(ctx context.Context, e *domain.Client) (*domain.Client, error)
	FindByText(ctx context.Context, text string, userID int64) (*domain.Client, error)
	FindClientByIDAndSecret(ctx context.Context, clientID, clientSecret string) (*domain.Client, error)
	FindManagingByText(ctx context.Context, text string) (*domain.Managing, error)
}

type UseCaseClient interface {
	RegisterClient(ctx context.Context, clientID, clientSecret, managingID string, userID int64) (*domain.Client, error)
	VerifyManaging(ctx context.Context, text string) (*domain.Managing, error)
	VerifyClient(ctx context.Context, clientID, clientSecret string) (*domain.Client, error)
}
