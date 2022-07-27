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
var Database = vmdb.NewDatabase(
	"pool-event",
	vcago.Settings.String("DB_HOST", "w", "localhost"),
	vcago.Settings.String("DB_PORT", "w", "27017"),
)

var ArtistCollection = Database.Collection("artists").CreateIndex("name", true)
var UserCollection = Database.Collection("users").CreateIndex("email", true)
var ParticipationCollection = Database.Collection("participations").CreateMultiIndex(bson.D{{Key: "user_id", Value: 1}, {Key: "event_id", Value: 1}}, true)
var OrganizerCollection = Database.Collection("organizers").CreateIndex("name", true)
var EventCollection = Database.Collection("events")

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
	vcago.Nats.Subscribe("user.created", func(m *models.User) {
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
