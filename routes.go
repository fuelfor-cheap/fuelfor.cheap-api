package main

import (
	"github.com/lynnau/fuelfor.cheap-api/config"
	"github.com/lynnau/fuelfor.cheap-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server ...
var Server = gin.Default()

func init() {
	// setup cors
	conf := cors.DefaultConfig()
	origins := config.GetStringList("cors.origins")
	conf.AllowOrigins = origins
	conf.AddExposeHeaders("X-Accesstoken", "X-DeviceID")
	Server.Use(cors.New(conf))

	// create the routes
	api := Server.Group("/api/v1")
	api.POST("/lock", controllers.NewLock)
	api.POST("/lock/existing", controllers.GetExistingLock)

	// sessions
	sessions := api.Group("/sessions")
	sessions.POST("/login", controllers.Login)
}
