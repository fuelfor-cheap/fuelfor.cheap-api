package models

import "time"

// PriceHistory is the collection of fuel prices at a given timestamp
type PriceHistory struct {
	tableName struct{}  `pg:"price_history"`
	Datetime  time.Time `pg:",pk"`
	FuelType  string    `pg:",pk"`
	Price     float64
	Postcode  int
}
