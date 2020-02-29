package db

import (
	"github.com/lynnau/fuelfor.cheap-api/config"

	"github.com/go-pg/pg/v9"
)

// DB ...
var DB *pg.DB

// Connect to the configured database
func Connect() {
	hosturi := config.Get("db.uri").(string)
	options, err := pg.ParseURL(hosturi)
	if err != nil {
		panic(err)
	}

	DB = pg.Connect(options)
}
