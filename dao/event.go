package dao

import (
	"context"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"go.mongodb.org/mongo-driver/bson"
)

var EventCollection = Database.Collection("events")

type EventQuery struct {
	vcapool.EventQuery
}
type EventCreate struct {
	vcapool.EventCreate
}

type Event vcapool.Event

func (i *EventCreate) Create(ctx context.Context, token *vcapool.AccessToken) (r *Event, err error) {
	database := i.Database(token)
	if err = EventCollection.InsertOne(ctx, database); err != nil {
		return
	}
	r = (*Event)(database.Event())
	artists := new(vcapool.ArtistList)
	if err = ArtistCollection.Find(ctx, vcago.NewFilter().In("_id", database.ArtistIDs).Bson(), artists); err != nil {
		return
	}
	r.Artists = *artists
	return
}

func (i *Event) Get(ctx context.Context, filter bson.M) (err error) {
	if err = EventCollection.FindOne(ctx, filter, i); err != nil {
		return
	}

	artists := new(vcapool.ArtistList)
	if err = ArtistCollection.Find(ctx, vcago.NewFilter().In("_id", i.ArtistIDs).Bson(), artists); err != nil {
		return
	}
	i.Artists = *artists
	return
}

func (i *Event) Update(ctx context.Context) (err error) {
	update := bson.M{"$set": i}
	err = EventCollection.UpdateOne(ctx, bson.M{"_id": i.ID}, update)
	return
}

func (i *Event) Delete(ctx context.Context, filter bson.M) (err error) {
	err = EventCollection.DeleteOne(ctx, filter)
	return
}

func (i *EventQuery) List(ctx context.Context) (r *vcapool.EventList, err error) {
	pipe := vcago.NewMongoPipe()
	pipe.Lookup(ArtistCollection.Name, "artist_ids", "_id", "artists")
	pipe.Match(i.Match())
	r = new(vcapool.EventList)
	err = EventCollection.Aggregate(ctx, pipe.Pipe, r)
	return
}
