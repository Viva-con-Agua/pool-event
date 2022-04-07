package main

import (
	"pool-event/handlers"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(vcago.Logger.Init("pool-event"))
	e.Use(vcago.CORS.Init())

	//error
	e.HTTPErrorHandler = vcago.HTTPErrorHandler
	e.Validator = vcago.JSONValidator

	events := e.Group("/events", vcago.AccessCookieMiddleware(&vcapool.AccessToken{}))

	eventHandler := handlers.NewEventHandler()
	event := events.Group("/event")
	event.Use(eventHandler.Context)
	event.POST("", eventHandler.Create)
	event.GET("", eventHandler.List)
	event.GET("/:id", eventHandler.GetByID)
	event.PUT("", eventHandler.Update)
	event.DELETE("/:id", eventHandler.DeleteByID)

	artistHandler := handlers.NewArtistHandler()
	artist := events.Group("/artist")
	artist.Use(artistHandler.Context)
	artist.POST("", artistHandler.Create)
	artist.GET("", artistHandler.List)
	artist.GET("/:id", artistHandler.GetByID)
	artist.PUT("", artistHandler.Update)
	artist.DELETE("/:id", artistHandler.DeleteByID)

	organizerHandler := handlers.NewOrganizerHandler()
	organizer := events.Group("/organizer")
	organizer.Use(organizerHandler.Context)
	organizer.POST("", organizerHandler.Create)
	organizer.GET("", organizerHandler.List)
	organizer.GET("/:id", organizerHandler.GetByID)
	organizer.PUT("", organizerHandler.Update)
	organizer.DELETE("/:id", organizerHandler.DeleteByID)

	tourHandler := handlers.NewTourHandler()
	tour := events.Group("/tour")
	tour.Use(tourHandler.Context)
	tour.POST("", tourHandler.Create)
	tour.GET("/:id", tourHandler.GetByID)
	tour.GET("", tourHandler.List)
	//server
	port := vcago.Config.GetEnvString("APP_PORT", "n", "1323")
	e.Logger.Fatal(e.Start(":" + port))
}
