package fakes

import (
	"context"
	"log"
	"ms-auth/internal/user/domain"
	"ms-auth/internal/user/repository"
	"ms-auth/pkg/postgres"
)

type UserPgRepositoryFake struct {
	repo *repository.UserPgRepository
}

func NewUserPgRepositoryFake() *UserPgRepositoryFake {

	db, err := postgres.InitPostgresTest()

	if err != nil {
		log.Fatalf("Postgres failed %s", err)
	}

	repo := repository.NewUserPgRepository(db)

	return &UserPgRepositoryFake{
		repo: repo,
	}
}

func (f *UserPgRepositoryFake) Create(ctx context.Context, e *domain.User) (*domain.User, error) {

	user, err := f.repo.Create(ctx, e)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (f *UserPgRepositoryFake) FindByID(ctx context.Context, userID int64) (user *domain.User, err error) {

	user, err = f.repo.FindByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (f *UserPgRepositoryFake) FindByUsernameOrEmail(ctx context.Context, username, email string) (user *domain.User, err error) {

	user, err = f.repo.FindByUsernameOrEmail(ctx, username, email)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (f *UserPgRepositoryFake) Update(ctx context.Context, userID int64, u *domain.User) (err error) {

	err = f.repo.Update(ctx, userID, u)

	if err != nil {
		return err
	}

	return nil

}

func (f *UserPgRepositoryFake) Destroy(ctx context.Context, userID int64) error {

	err := f.repo.Destroy(ctx, userID)

	if err != nil {
		return err
	}

	return nil

}
