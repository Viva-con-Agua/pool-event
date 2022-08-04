package dao

import (
	"context"
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

func InitialCollections() {
	ArtistCollection = Database.Collection("artists").CreateIndex("name", true)
	UserCollection = Database.Collection("users").CreateIndex("email", true)
	ParticipationCollection = Database.Collection("participations").CreateMultiIndex(bson.D{{Key: "user_id", Value: 1}, {Key: "event_id", Value: 1}}, true)
	OrganizerCollection = Database.Collection("organizers").CreateIndex("name", true)
	EventCollection = Database.Collection("events")
	SourceCollection = Database.Collection("sources")
	TakingCollection = Database.Collection("takings").CreateIndex("event_id", true)

}

func SubscribeUserCreate() {
	vcago.Nats.Subscribe("user.created", func(m *models.User) {
		ctx := context.Background()
		if err := UserCollection.InsertOne(ctx, m); err != nil {
			output := vcago.NewError(err, "ERROR", "nats")
			output.Print("internal")
		}
	})
}

func SubscribeUserUpdate() {
	vcago.Nats.Subscribe("user.updated", func(m *models.User) {
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
