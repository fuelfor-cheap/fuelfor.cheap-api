package models

// FuelType contains the information needed to describe a specific fuel
type FuelType struct {
	ID           string `pg:",pk"`
	Name         string
	OctaneRating float64
}
