package dao

import (
	"context"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"go.mongodb.org/mongo-driver/bson"
)

var TourCollection = Database.Collection("tours")

type TourQuery struct {
	vcapool.TourQuery
}

type Tour vcapool.Tour

type TourCreate struct {
	vcapool.TourCreate
}

type TourUpdate vcapool.TourUpdate

func (i *TourCreate) Create(ctx context.Context, token *vcapool.AccessToken) (r *vcapool.Tour, err error) {
	database := i.Database(token)
	if err = TourCollection.InsertOne(ctx, database); err != nil {
		return
	}
	events := i.Events.Database(database.ID, token).Insert()
	if events != nil {
		if err = EventCollection.InsertMany(ctx, events); err != nil {
			return
		}
	}
	r = database.Tour()
	pipe := vcago.NewMongoPipe()
	pipe.Lookup(ArtistCollection.Name, "artist_ids", "_id", "artists")
	list := new(vcapool.EventList)
	match := vcago.NewMongoMatch()
	match.EqualString("tour_id", r.ID)
	pipe.Match(match)
	if err = EventCollection.Aggregate(ctx, pipe.Pipe, list); err != nil {
		return
	}
	r.Events = *list
	return
}

func (i *Tour) Get(ctx context.Context, filter bson.M) (err error) {
	if err = TourCollection.FindOne(ctx, filter, i); err != nil {
		return
	}
	pipe := vcago.NewMongoPipe()
	pipe.Lookup(ArtistCollection.Name, "artist_ids", "_id", "artists")
	list := new(vcapool.EventList)
	match := vcago.NewMongoMatch()
	match.EqualString("tour_id", i.ID)
	pipe.Match(match)
	if err = EventCollection.Aggregate(ctx, pipe.Pipe, list); err != nil {
		return
	}
	i.Events = *list
	artists := new(vcapool.ArtistList)

	if i.ArtistIDs != nil {
		if err = ArtistCollection.Find(ctx, bson.M{"_id": bson.M{"$in": i.ArtistIDs}}, artists); err != nil {
			return
		}
		i.Artists = *artists
	}
	return
}

func (i *TourUpdate) Update(ctx context.Context) (r *vcapool.Tour, err error) {
	i.Modified.Update()
	update := bson.M{"$set": i}
	if err = TourCollection.UpdateOne(ctx, bson.M{"_id": i.ID}, update); err != nil {
		return
	}
	r = new(vcapool.Tour)
	err = TourCollection.FindOne(ctx, bson.M{"_id": i.ID}, r)
	return
}

func (i *Tour) Delete(ctx context.Context, filter bson.M) (err error) {
	err = EventCollection.DeleteOne(ctx, filter)
	return
}

func (i *TourQuery) List(ctx context.Context) (r *vcapool.TourList, err error) {
	pipe := vcago.NewMongoPipe()
	pipe.Lookup(EventCollection.Name, "_id", "tour_id", "events")
	pipe.Lookup(ArtistCollection.Name, "artist_ids", "_id", "artists")
	pipe.Match(i.Match())
	r = new(vcapool.TourList)
	err = TourCollection.Aggregate(ctx, pipe.Pipe, r)
	return
}
