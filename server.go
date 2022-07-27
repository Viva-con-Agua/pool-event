package main

import (
	"pool-event/handlers/admin"
	"pool-event/handlers/token"

	"github.com/Viva-con-Agua/vcago"
)

func main() {
	e := vcago.NewEchoServer("pool-event")

	admins := e.Group("/admin/events")
	admin.User.Routes(admins.Group("/user"))
	events := e.Group("/events")
	token.Event.Routes(events.Group("/event"))
	token.Artist.Routes(events.Group("/artist"))
	token.Organizer.Routes(events.Group("/organizer"))
	token.Participation.Routes(events.Group("/participation"))
	//token.Tour.Routes(events.Group("/tour"))

	//server
	port := vcago.Settings.String("APP_PORT", "n", "1323")
	e.Logger.Fatal(e.Start(":" + port))
}
