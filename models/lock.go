package models

import "time"

// LockPayload ...
type LockPayload struct {
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	AccessToken  string  `json:"access_token"`
	DeviceID     string  `json:"device_id"`
	DeviceSecret string  `json:"device_secret"`
	FuelType     string  `json:"fuel_type"`
	AccountID    string  `json:"account_id"`
}

// Lock ...
type Lock struct {
	FuelType      string  `json:"fuel_type"`
	CentsPerLitre float64 `json:"cents_per_litre"`
	// TotalLitres   float64   `json:"total_litres"`
	Expires time.Time `json:"expires"`
}
