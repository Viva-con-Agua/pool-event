package dao

import (
	"context"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Artist vcapool.Artist

type ArtistQuery vcapool.ArtistQuery

var ArtistCollection = Database.Collection("artists").CreateIndex("name", true)

func (i *Artist) Create(ctx context.Context) (err error) {
	i.ID = uuid.NewString()
	i.Modified = vcago.NewModified()
	err = ArtistCollection.InsertOne(ctx, i)
	return
}

func (i *Artist) Get(ctx context.Context, filter bson.M) (err error) {
	err = ArtistCollection.FindOne(ctx, filter, i)
	return
}

func (i *ArtistQuery) List(ctx context.Context) (r *vcapool.ArtistList, err error) {
	pipe := vcago.NewMongoPipe()
	pipe.Match((*vcapool.ArtistQuery)(i).Match())
	r = new(vcapool.ArtistList)
	err = ArtistCollection.Aggregate(ctx, pipe.Pipe, r)
	return
}

func (i *Artist) Update(ctx context.Context) (err error) {
	i.Modified.Update()
	update := bson.M{"$set": i}
	err = ArtistCollection.UpdateOne(ctx, bson.M{"_id": i.ID}, update)
	return
}

func (i Artist) Delete(ctx context.Context, filter bson.M) (err error) {
	err = ArtistCollection.DeleteOne(ctx, filter)
	return
}
