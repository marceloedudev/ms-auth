package usecase

import (
	"context"
	"ms-auth/internal/client/domain"
	"time"

	"github.com/pkg/errors"
)

func (s *Service) RegisterClient(ctx context.Context, clientID, clientSecret, managingID string, userID int64) (result *domain.Client, err error) {

	rClient := &domain.Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		UserID:       userID,
		ManagingID:   managingID,
		Enabled:      true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = rClient.ValidateAddClient()

	if err != nil {
		return nil, errors.Wrap(err, "usecase.ValidateAddClient")
	}

	existsClient, err := s.pgRepository.FindByText(ctx, rClient.ClientID, rClient.UserID)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.FindByText")
	}

	if existsClient != nil {
		return nil, domain.ErrIDOrNameAlreadyUsed
	}

	client, err := s.pgRepository.Create(ctx, rClient)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.Create")
	}

	return client, nil

}
