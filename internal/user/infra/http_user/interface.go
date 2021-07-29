package http_user

type UserTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}

type UserTokenLoginRequest struct {
	GrantType string `json:"grant_type" example:"password"`
	Username  string `json:"username" example:"_example_"`
	Password  string `json:"password" example:"123456"`
}

type UserTokenRefreshRequest struct {
	GrantType    string `json:"grant_type" example:"refresh_token"`
	RefreshToken string `json:"refresh_token" example:"28591e38-6f52-4fbc-97ce-747d76edfa2f"`
}

type UserRevokeTokenRequest struct {
	Token     string `json:"token" example:"6f19ebb1-0e5c-4845-a48f-70bbcf4fcd1a"`
	GrantType string `json:"grant_type" example:"refresh_token"`
}

type UserInfoTokenRequest struct {
	ClientID     string `json:"client_id" example:"management"`
	ClientSecret string `json:"client_secret" example:"hash-a0cc1f57-407d-4d87"`
	AccessToken  string `json:"access_token" example:"3f0dcdfe-7117-461d-b54e-bfc6c039a173"`
}

type UserRegisterRequest struct {
	Username  string `json:"username" example:"_example+"`
	FirstName string `json:"first_name" example:"Example"`
	LastName  string `json:"last_name" example:"Admin"`
	Password  string `json:"password" example:"123456"`
	Email     string `json:"email" example:"example@gmail.com"`
}
