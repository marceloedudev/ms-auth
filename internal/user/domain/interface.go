package domain

import (
	"database/sql"
	"time"
)

type User struct {
	ID            int64     `json:"id"`
	UUID          string    `json:"uuid"`
	Username      string    `json:"username"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Password      string    `json:"-"`
	ManagingID    string    `json:"managing_id"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	Enabled       bool      `json:"enabled"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// only pg/db
type UserSQL struct {
	ID        int64  `json:"id"`
	UUID      string `json:"uuid"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`
	ManagingID    sql.NullString `json:"managing_id"`
	Email         string         `json:"email"`
	EmailVerified bool           `json:"email_verified"`
	Enabled       bool           `json:"enabled"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type TokenUUID struct {
	AccessUUID  string
	RefreshUUID string
}

type TokenExpires struct {
	AtExpires int64
	RtExpires int64
}

type TokenData struct {
	TokenUUID
	TokenExpires
}

type TokenJWT struct {
	UserID      int64  `json:"user_id"`
	AccessUUID  string `json:"access_uuid"`
	RefreshUUID string `json:"refresh_uuid"`
}

type ExpireToken struct {
	ExpireTimeAt time.Duration
	ExpireTimeRt time.Duration
}

type TokenResponse struct {
	UserID       int64  `json:"user_id" example:"1"`
	AccessToken  string `json:"access_token" example:"3f0dcdfe-7117-461d-b54e-bfc6c039a173"`
	ExpiresIn    int    `json:"expires_in" example:"3600"`
	TokenType    string `json:"token_type" example:"Bearer"`
	RefreshToken string `json:"refresh_token" example:"28591e38-6f52-4fbc-97ce-747d76edfa2f"`
}

type UserDataToken struct {
	ID         int64  `json:"id" example:"1"`
	Username   string `json:"username" example:"_example_"`
	FirstName  string `json:"first_name" example:"Example"`
	LastName   string `json:"last_name" example:"Admin"`
	Email      string `json:"email" example:"example@gmail.com"`
	ManagingID string `json:"managing_id" example:"master"`
}
