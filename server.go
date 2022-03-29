package main

import (
	"pool-event/handlers"

	"github.com/Viva-con-Agua/vcago"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(vcago.Logger.Init("pool-core"))
	e.Use(vcago.CORS.Init())

	//error
	e.HTTPErrorHandler = vcago.HTTPErrorHandler
	e.Validator = vcago.JSONValidator

	events := e.Group("/events")
	event := e.Group("event")
	event.POST("", handlers.EventCreate)
	event.GET("", handlers.EventList)
	event.GET("/:id", handlers.EventGetByID)
	event.PUT("", handlers.EventUpdate)
	event.DELETE("/:id", handlers.EventDeleteByID)
	artist := events.Group("/artist")
	artist.POST("", handlers.ArtistCreate)
	artist.GET("", handlers.ArtistList)
	artist.GET("/:id", handlers.ArtistGetByID)
	artist.PUT("", handlers.ArtistUpdate)
	artist.DELETE("/:id", handlers.ArtistDeleteByID)
	//server
	port := vcago.Config.GetEnvString("APP_PORT", "n", "1323")
	e.Logger.Fatal(e.Start(":" + port))
}
