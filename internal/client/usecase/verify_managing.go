package usecase

import (
	"context"
	"ms-auth/internal/client/domain"

	"github.com/pkg/errors"
)

func (s *Service) VerifyManaging(ctx context.Context, text string) (existsManaging *domain.Managing, err error) {

	existsManaging, err = s.pgRepository.FindManagingByText(ctx, text)

	if err != nil {
		return nil, errors.Wrap(err, "usecase.FindManagingByText")
	}

	return existsManaging, nil

}
