package main

import (
	"github.com/lynnau/fuelfor.cheap-api/config"
)

func main() {
	port := config.Get("port").(string)
	defer Server.Run(port)
}
