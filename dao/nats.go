package dao

import (
	"context"
	"pool-event/models"

	"github.com/Viva-con-Agua/vcago"
	"go.mongodb.org/mongo-driver/bson"
)

func InitialNats() {
	vcago.Nats.Connect()
	vcago.Nats.Subscribe("pool-user.user.created", SubscribeUserCreate)
	vcago.Nats.Subscribe("pool-user.user.updated", SubscribeUserUpdate)
}

func SubscribeUserCreate(m *models.User) {
	ctx := context.Background()
	if err := UserCollection.InsertOne(ctx, m); err != nil {
		output := vcago.NewError(err, "ERROR", "nats")
		output.Print("internal")
	}
}

func SubscribeUserUpdate(m *models.User) {
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
}
