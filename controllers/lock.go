package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lynnau/fuelfor.cheap-api/models"
	"github.com/lynnau/fuelfor.cheap-api/seveneleven"
)

// NewLock ...
func NewLock(ctx *gin.Context) {
	payload := &models.LockPayload{}
	err := ctx.ShouldBind(payload)
	if err != nil {
		ctx.JSON(500, invalidScemaError(payload))
		return
	}

	// check if the email and or password is intact
	if payload.Longitude == 0 || payload.Latitude == 0 {
		ctx.JSON(500, invalidScemaError(payload))
		return
	}

	// ask the api to lock in a fuel price for a type at the specified location
	response := seveneleven.Lock(payload.Latitude, payload.Longitude, payload.AccessToken, payload.DeviceSecret, payload.DeviceID, payload.FuelType, payload.AccountID)
	fuelType := seveneleven.EanToFuelType[response.FuelGradeModel]
	expires := time.Unix(int64(response.ExpiresAt), 0)

	lock := &models.Lock{
		FuelType:      fuelType,
		CentsPerLitre: response.CentsPerLitre,
		Expires:       expires,
	}

	ctx.JSON(200, lock)
}

// GetExistingLock ...
func GetExistingLock(ctx *gin.Context) {
	details := &models.Auth{}
	err := ctx.ShouldBind(details)
	if err != nil {
		ctx.JSON(500, invalidScemaError(details))
		return
	}

	// check if the email and or password is intact
	if details.AccessToken == "" || details.DeviceID == "" {
		ctx.JSON(500, invalidScemaError(details))
		return
	}

	// ask the api to lock in a fuel price for a type at the specified location
	response := seveneleven.GetLock(details.AccessToken, details.DeviceSecret, details.DeviceID)
	locks := []*models.Lock{}

	for _, resp := range response {
		fuelType := seveneleven.EanToFuelType[resp.FuelGradeModel]
		expires := time.Unix(int64(resp.ExpiresAt), 0)

		lock := &models.Lock{
			FuelType:      fuelType,
			CentsPerLitre: resp.CentsPerLitre,
			// TotalLitres: resp.TotalLitres,
			Expires: expires,
		}
		locks = append(locks, lock)
	}

	ctx.JSON(200, locks)
}
