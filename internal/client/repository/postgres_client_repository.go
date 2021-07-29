package repository

import (
	"context"
	"database/sql"
	"ms-auth/internal/client/domain"

	"github.com/jmoiron/sqlx"
)

type ClientPgRepository struct {
	db *sqlx.DB
}

func NewClientPgRepository(db *sqlx.DB) *ClientPgRepository {
	return &ClientPgRepository{
		db: db,
	}
}

func (u *ClientPgRepository) Create(ctx context.Context, e *domain.Client) (user *domain.Client, err error) {

	var id int64
	if err := u.db.QueryRowContext(ctx, createClientQuery, e.ClientID, e.ClientSecret, e.UserID, e.ManagingID, e.Enabled, e.CreatedAt, e.UpdatedAt).Scan(&id); err != nil {
		return nil, err
	}

	e.ID = id

	return e, nil

}

func (u *ClientPgRepository) FindByText(ctx context.Context, text string, userID int64) (client *domain.Client, err error) {

	client = &domain.Client{}

	if err := u.db.QueryRowContext(ctx, findClientByTextAndUserIDQuery, text, userID).Scan(&client.ID, &client.ClientID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return client, nil

}

func (u *ClientPgRepository) FindClientByIDAndSecret(ctx context.Context, clientID, clientSecret string) (client *domain.Client, err error) {

	client = &domain.Client{}

	if err := u.db.QueryRowContext(ctx, findClientByIDAndSecret, clientID, clientSecret).Scan(&client.ID, &client.ClientID, &client.ClientSecret); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return client, nil

}

func (u *ClientPgRepository) FindManagingByText(ctx context.Context, text string) (managing *domain.Managing, err error) {

	managing = &domain.Managing{}

	if err := u.db.QueryRowContext(ctx, findManagingByTextQuery, text).Scan(&managing.ID, &managing.ManagingID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return managing, nil

}
