package domain

import "time"

type Client struct {
	ID           int64     `json:"id" example:"1"`
	ClientID     string    `json:"client_id" example:"management"`
	ClientSecret string    `json:"client_secret" example:"a0cc1f57-407d-4d87"`
	UserID       int64     `json:"user_id" example:"1"`
	ManagingID   string    `json:"managing_id" example:"master"`
	Enabled      bool      `json:"enabled" example:"true"`
	CreatedAt    time.Time `json:"created_at" example:"2021-04-10T21:09:38.180471Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2021-04-10T21:09:38.180471Z"`
}

type Managing struct {
	ID         int64  `json:"id"`
	ManagingID string `json:"managing_id"`
	Enabled    bool   `json:"enabled"`
}
