package usecase

import (
	"context"
	"ms-auth/internal/client/domain"

	"github.com/pkg/errors"
)

func (s *Service) VerifyClient(ctx context.Context, clientID, clientSecret string) (client *domain.Client, err error) {

	client, err = s.pgRepository.FindClientByIDAndSecret(ctx, clientID, clientSecret)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.FindClientByIDAndSecret")
	}

	return client, nil

}
