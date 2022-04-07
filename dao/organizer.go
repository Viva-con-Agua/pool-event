package dao

import (
	"context"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Organizer vcapool.Organizer

type OrganizerQuery vcapool.OrganizerQuery

var OrganizerCollection = Database.Collection("organizers").CreateIndex("name", true)

func (i *Organizer) Create(ctx context.Context) (err error) {
	i.ID = uuid.NewString()
	i.Modified = vcago.NewModified()
	err = OrganizerCollection.InsertOne(ctx, i)
	return
}

func (i *Organizer) Get(ctx context.Context, filter bson.M) (err error) {
	err = OrganizerCollection.FindOne(ctx, filter, i)
	return
}

func (i *OrganizerQuery) List(ctx context.Context) (r *vcapool.OrganizerList, err error) {
	pipe := vcago.NewMongoPipe()
	pipe.Match((*vcapool.OrganizerQuery)(i).Match())
	r = new(vcapool.OrganizerList)
	err = OrganizerCollection.Aggregate(ctx, pipe.Pipe, r)
	return
}

func (i *Organizer) Update(ctx context.Context) (err error) {
	i.Modified.Update()
	update := bson.M{"$set": i}
	err = OrganizerCollection.UpdateOne(ctx, bson.M{"_id": i.ID}, update)
	return
}

func (i Organizer) Delete(ctx context.Context, filter bson.M) (err error) {
	err = OrganizerCollection.DeleteOne(ctx, filter)
	return
}
