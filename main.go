package main

import (
	"github.com/lynnau/fuelfor.cheap-api/config"
)

func main() {
	port := config.Get("port").(string)
	defer Server.Run(port)

	// email := ""
	// password := ""
	// resp, accessToken, deviceID := seveneleven.Login(email, password, "")
	// seveneleven.Lock(
	// 	-37.920945,
	// 	145.167859,
	// 	accessToken,
	// 	resp.DeviceSecretToken,
	// 	deviceID,
	// 	"Extra 95",
	// 	resp.AccountID)
}
