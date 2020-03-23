package controllers

import (
	"encoding/json"

	"github.com/lynnau/fuelfor.cheap-api/models"
	"github.com/lynnau/fuelfor.cheap-api/seveneleven"

	"github.com/gin-gonic/gin"
)

func invalidScemaError(scheme interface{}) gin.H {
	schema, err := json.Marshal(scheme)
	if err != nil {
		panic(err)
	}

	return gin.H{
		"error":   true,
		"message": "malformed payload details, request body should follow the follwing schema",
		"schema":  string(schema),
	}
}

// Login a user to the 7/11 api and return their info and access token
func Login(ctx *gin.Context) {
	details := &models.LoginDetails{}
	err := ctx.ShouldBind(details)
	if err != nil {
		ctx.JSON(500, invalidScemaError(details))
		return
	}

	// check if the email and or password is intact
	if details.Email == "" || details.Password == "" {
		ctx.JSON(500, invalidScemaError(details))
		return
	}

	// log the user in and return the api response from 7/11
	response, accessToken, deviceID := seveneleven.Login(details.Email, details.Password, details.AccessToken, details.DeviceID)

	ctx.Header("X-Accesstoken", accessToken)
	ctx.Header("X-DeviceID", deviceID)
	ctx.JSON(200, response)
}
