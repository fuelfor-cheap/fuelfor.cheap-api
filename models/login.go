package models

// LoginDetails ...
type LoginDetails struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}
