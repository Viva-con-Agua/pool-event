package dao

import (
	"log"

	"github.com/Viva-con-Agua/vcago/vmdb"
	"go.mongodb.org/mongo-driver/bson"
)

var Database *vmdb.Database

var ArtistCollection *vmdb.Collection
var UserCollection *vmdb.Collection
var ParticipationCollection *vmdb.Collection
var OrganizerCollection *vmdb.Collection
var EventCollection *vmdb.Collection
var SourceCollection *vmdb.Collection
var TakingCollection *vmdb.Collection

func InitialDatabase() {
	Database = vmdb.NewDatabase("pool-event").Connect()
}
func InitialTestDatabase() {
	Database = vmdb.NewDatabase("pool-event-test").Connect()
}

func InitialCollections() {
	log.Print("--- Initial Collections Start --- ")
	ArtistCollection = Database.Collection("artists").CreateIndex("name", true)
	UserCollection = Database.Collection("users").CreateIndex("email", true)
	ParticipationCollection = Database.Collection("participations").CreateMultiIndex(
		bson.D{
			{Key: "user_id", Value: 1},
			{Key: "event_id", Value: 1},
		}, true)
	OrganizerCollection = Database.Collection("organizers").CreateIndex("name", true)
	EventCollection = Database.Collection("events")
	SourceCollection = Database.Collection("sources")
	TakingCollection = Database.Collection("takings").CreateIndex("event_id", true)
	log.Print("--- Initial Collections Finish")
}
