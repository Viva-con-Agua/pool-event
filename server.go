package main

import (
	"pool-event/handlers/admin"
	"pool-event/handlers/token"

	"github.com/Viva-con-Agua/vcago"
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
	admins := e.Group("/admin/events")
	admin.User.Routes(admins.Group("/user"))
	events := e.Group("/events")
	token.Event.Routes(events.Group("/event"))
	token.Artist.Routes(events.Group("/artist"))
	token.Organizer.Routes(events.Group("/organizer"))
	token.Participation.Routes(events.Group("/participation"))
	//token.Tour.Routes(events.Group("/tour"))

	//server
	port := vcago.Config.GetEnvString("APP_PORT", "n", "1323")
	e.Logger.Fatal(e.Start(":" + port))
}
