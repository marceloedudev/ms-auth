package fakes

import (
	"context"
	"log"
	"ms-auth/internal/client/domain"
	"ms-auth/internal/client/repository"
	"ms-auth/pkg/postgres"
)

type ClientPgRepositoryFake struct {
	repo *repository.ClientPgRepository
}

func NewClientPgRepositoryFake() *ClientPgRepositoryFake {
	db, err := postgres.InitPostgresTest()

	if err != nil {
		log.Fatalf("Postgres failed %s", err)
	}

	repo := repository.NewClientPgRepository(db)

	return &ClientPgRepositoryFake{
		repo: repo,
	}
}

func (f *ClientPgRepositoryFake) Create(ctx context.Context, e *domain.Client) (*domain.Client, error) {

	client, err := f.repo.Create(ctx, e)

	if err != nil {
		return nil, err
	}

	return client, nil

}

func (f *ClientPgRepositoryFake) FindByText(ctx context.Context, text string, userID int64) (client *domain.Client, err error) {

	client, err = f.repo.FindByText(ctx, text, userID)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (f *ClientPgRepositoryFake) FindClientByIDAndSecret(ctx context.Context, clientID, clientSecret string) (client *domain.Client, err error) {

	client, err = f.repo.FindClientByIDAndSecret(ctx, clientID, clientSecret)

	if err != nil {
		return nil, err
	}

	return client, nil

}

func (f *ClientPgRepositoryFake) FindManagingByText(ctx context.Context, text string) (managing *domain.Managing, err error) {

	managing, err = f.repo.FindManagingByText(ctx, text)

	if err != nil {
		return nil, err
	}

	return managing, nil

}
