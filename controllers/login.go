package controllers

import (
	"encoding/json"

	"github.com/lynnau/fuelfor.cheap-api/models"
	"github.com/lynnau/fuelfor.cheap-api/seveneleven"

	"github.com/gin-gonic/gin"
)

func invalidScemaError() gin.H {
	schema, err := json.Marshal(&models.LoginDetails{})
	if err != nil {
		panic(err)
	}

	return gin.H{
		"error":   true,
		"message": "malformed login details, request body should follow the follwing schema",
		"schema":  string(schema),
		"notes":   "the access_token in the schema is optional",
	}
}

// Login a user to the 7/11 api and return their info and access token
func Login(ctx *gin.Context) {
	details := &models.LoginDetails{}
	err := ctx.ShouldBind(details)
	if err != nil {
		ctx.JSON(500, invalidScemaError())
		return
	}

	// check if the email and or password is intact
	if details.Email == "" || details.Password == "" {
		ctx.JSON(500, invalidScemaError())
		return
	}

	// log the user in and return the api response from 7/11
	response, accessToken, deviceID := seveneleven.Login(details.Email, details.Password, details.AccessToken)

	ctx.Header("X-Accesstoken", accessToken)
	ctx.Header("X-DeviceID", deviceID)
	ctx.JSON(200, response)
}
