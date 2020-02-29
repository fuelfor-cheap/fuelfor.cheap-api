package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lynnau/fuelfor.cheap-api/db"
	"github.com/lynnau/fuelfor.cheap-api/db/models"
)

func main() {
	// connect to the db
	db.Connect()

	// wait until everything is finished before we close the database connection
	// to save ourselves resources...
	defer db.DB.Close()

	// fetch all the timestamps that we know in the database
	pHistory := []*models.PriceHistory{}
	err := db.DB.Model(&pHistory).Column("datetime").Group("datetime").Select()
	if err != nil {
		panic(err)
	}

	// now that we have a list of known datetimes,
	// we need to fetch some new price from the api
	response, err := fetch()
	if err != nil {
		panic(err)
	}

	// compare if the api response timestamp is one we already have data for from the api
	var hasExisting bool
	for _, priceHistory := range pHistory {
		have := priceHistory.Datetime.Unix()
		if have == response.Updated {
			hasExisting = true
			break
		}
	}

	if hasExisting {
		// exit the script, we don't need to do anything else...
		fmt.Println("api has not updated yet, timestamped:", response.Updated)
		os.Exit(1)
	}

	// populate the fuel types so we know what's what from the API
	fTypes := []*models.FuelType{}
	err = db.DB.Model(&fTypes).Column("id", "name").Select()
	if err != nil {
		panic(err)
	}

	// and now that we've gotten this far, we'll need to add the records to the db
	histories := []*models.PriceHistory{}
	timestamp := time.Unix(response.Updated, 0)
	for _, region := range response.Regions {
		if region.Region != "All" {
			continue
		}

		for _, price := range region.Prices {
			// find our corrosponding fuel type that matches this one
			fuel, err := fuelInitialToFuelType(price.Type, fTypes)
			if err != nil {
				continue
			}

			postcode, _ := strconv.Atoi(price.Postcode)
			history := &models.PriceHistory{
				Datetime: timestamp,
				FuelType: fuel.ID,
				Price:    price.Price,
				Postcode: postcode,
			}
			histories = append(histories, history)
		}
	}

	// insert the list of price histories into the database
	_, err = db.DB.Model(&histories).Insert()
	if err != nil {
		panic(err)
	}

	fmt.Println("fetched new prices from the api, timestamped:", response.Updated)
}
