package main

import (
	"pool-event/dao"
	"pool-event/handlers/admin"
	"pool-event/handlers/token"

	"github.com/Viva-con-Agua/vcago"
)

func main() {
	e := vcago.NewServer()
	dao.InitialDatabase()
	dao.InitialCollections()
	dao.InitialNats()
	admins := e.Group("/admin/events")
	admin.User.Routes(admins.Group("/user"))
	events := e.Group("/events")
	token.Event.Routes(events.Group("/event"))
	token.Artist.Routes(events.Group("/artist"))
	token.Organizer.Routes(events.Group("/organizer"))
	token.Participation.Routes(events.Group("/participation"))
	//token.Tour.Routes(events.Group("/tour"))
	e.Run()
}
