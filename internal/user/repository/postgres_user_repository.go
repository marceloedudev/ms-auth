package repository

import (
	"context"
	"database/sql"
	"ms-auth/internal/user/domain"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserPgRepository struct {
	db *sqlx.DB
}

func NewUserPgRepository(db *sqlx.DB) *UserPgRepository {
	return &UserPgRepository{
		db: db,
	}
}

func (u *UserPgRepository) Create(ctx context.Context, e *domain.User) (user *domain.User, err error) {

	var id int64
	if err := u.db.QueryRowContext(ctx, createUserQuery, e.UUID, e.Username, e.FirstName, e.LastName, e.Password, e.Email, e.EmailVerified, e.Enabled, e.CreatedAt, e.UpdatedAt).Scan(&id); err != nil {
		return nil, errors.Wrap(err, "db.QueryRowContext")
	}

	e.ID = id

	return e, nil

}

func (u *UserPgRepository) FindByID(ctx context.Context, userID int64) (result *domain.User, err error) {

	user := &domain.UserSQL{}

	if err := u.db.QueryRowContext(ctx, findUserByIDQuery, userID).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.ManagingID, &user.Enabled, &user.EmailVerified, &user.UUID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "db.QueryRowContext")
	}

	result = &domain.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ManagingID: user.ManagingID.String,
		Enabled: user.Enabled,
		EmailVerified: user.EmailVerified,
		UUID: user.UUID,
	}

	return result, nil

}

func (u *UserPgRepository) FindByUsernameOrEmail(ctx context.Context, username, email string) (result *domain.User, err error) {

	user := &domain.UserSQL{}

	if err := u.db.QueryRowContext(ctx, findUserByUsernameOrEmailQuery, username, email).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.ManagingID, &user.Enabled, &user.EmailVerified, &user.UUID); err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, errors.New("no rows in result set") // default
		// }
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "db.QueryRowContext")
	}

	result = &domain.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ManagingID: user.ManagingID.String,
		Enabled: user.Enabled,
		EmailVerified: user.EmailVerified,
		UUID: user.UUID,
	}

	return result, nil

}

func (u *UserPgRepository) Update(ctx context.Context, userID int64, doc *domain.User) (err error) {

	if _, err := u.db.ExecContext(ctx, updateUserQuery, doc.Username, doc.Email, doc.UpdatedAt, userID); err != nil {
		return errors.Wrap(err, "db.ExecContext")
	}

	return nil

}

func (u *UserPgRepository) Destroy(ctx context.Context, userID int64) error {

	if _, err := u.db.ExecContext(ctx, destroyUserQuery, userID); err != nil {
		return errors.Wrap(err, "db.ExecContext")
	}

	return nil

}
