package http_client

type ClientRegisterRequest struct {
	ClientID     string `json:"client_id" example:"management"`
	ClientSecret string `json:"client_secret" example:"a0cc1f57-407d-4d87"`
	UserID       int64  `json:"user_id" example:"1"`
	ManagingID   string `json:"managing_id" example:"master"`
}
