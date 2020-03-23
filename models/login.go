package models

// Auth ...
type Auth struct {
	AccessToken  string `json:"access_token"`
	DeviceID     string `json:"device_id"`
	DeviceSecret string `json:"device_secret"`
}

// LoginDetails ...
type LoginDetails struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
	DeviceID    string `json:"device_id"`
}
