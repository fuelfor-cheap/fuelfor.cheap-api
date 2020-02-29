package main

import (
	"errors"
	"fmt"

	"github.com/lynnau/fuelfor.cheap-api/db/models"
)

// TODO: make this a lot better
func fuelInitialToFuelType(initial string, types []*models.FuelType) (*models.FuelType, error) {
	switch initial {
	case "E10":
		for _, fuelType := range types {
			if fuelType.Name == "Special E10 94" {
				return fuelType, nil
			}
		}
		break

	case "U91":
		for _, fuelType := range types {
			if fuelType.Name == "Special Unleaded 91" {
				return fuelType, nil
			}
		}
		break

	case "U95":
		for _, fuelType := range types {
			if fuelType.Name == "Extra 95" {
				return fuelType, nil
			}
		}
		break

	case "U98":
		for _, fuelType := range types {
			if fuelType.Name == "Supreme+ 98" {
				return fuelType, nil
			}
		}
		break

	case "Diesel":
		for _, fuelType := range types {
			if fuelType.Name == "Special Diesel" {
				return fuelType, nil
			}
		}
		break

	case "LPG":
		for _, fuelType := range types {
			if fuelType.Name == "LPG" {
				return fuelType, nil
			}
		}
		break

	default:
		err := fmt.Sprintf("no matching fuel type for the given initial, %s", initial)
		return nil, errors.New(err)
	}

	err := fmt.Sprintf("no matching fuel type for the given initial, %s", initial)
	return nil, errors.New(err)
}
