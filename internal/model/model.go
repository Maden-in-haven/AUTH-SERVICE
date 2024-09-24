package model

type AuthLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Tokens struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}