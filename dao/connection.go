package dao

import (
	"context"
	"log"
	"pool-event/models"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/vmdb"
	"go.mongodb.org/mongo-driver/bson"
)

var Logger = vcago.Logger
var AdminRequest = vcago.NewAdminRequest()

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

func InitialCollections() (err error) {
	log.Print("--- Initial Collections --- ")
	if ArtistCollection, err = Database.Collection("artists").CreateIndex("name", true); err != nil {
		return
	}
	if UserCollection, err = Database.Collection("users").CreateIndex("email", true); err != nil {
		return
	}
	if ParticipationCollection, err = Database.Collection("participations").CreateMultiIndex(
		bson.D{
			{Key: "user_id", Value: 1},
			{Key: "event_id", Value: 1},
		}, true); err != nil {
		return
	}
	if OrganizerCollection, err = Database.Collection("organizers").CreateIndex("name", true); err != nil {
		return
	}
	EventCollection = Database.Collection("events")
	SourceCollection = Database.Collection("sources")
	if TakingCollection, err = Database.Collection("takings").CreateIndex("event_id", true); err != nil {
		return
	}
	return
}

func SubscribeUserCreate() {
	vcago.Nats.Subscribe("pool-user.user.created", func(m *models.User) {
		ctx := context.Background()
		if err := UserCollection.InsertOne(ctx, m); err != nil {
			output := vcago.NewError(err, "ERROR", "nats")
			output.Print("internal")
		}
	})
}

func SubscribeUserUpdate() {
	vcago.Nats.Subscribe("pool-user.user.updated", func(m *models.User) {
		ctx := context.Background()
		if err := UserCollection.UpdateOne(
			ctx,
			bson.D{{Key: "_id", Value: m.ID}},
			bson.D{{Key: "$set", Value: m}},
			nil,
		); err != nil {
			output := vcago.NewError(err, "ERROR", "nats")
			output.Print("internal")
		}
	})
}
