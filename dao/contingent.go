package dao

/*
import (
	"context"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Contingent vcapool.Contingent
type ContingentQuery vcapool.ContingentQuery

var ContingentCollection = Database.Collection("contingents")

func (i *Contingent) Create(ctx context.Context) (err error) {
	i.ID = uuid.NewString()
	i.Modified = vcago.NewModified()
	err = ContingentCollection.InsertOne(ctx, i)
	return
}

func (i *Contingent) Get(ctx context.Context, filter bson.M) (err error) {
	err = ContingentCollection.FindOne(ctx, filter, i)
	return
}

func (i *Contingent) Update(ctx context.Context) (err error) {
	i.Modified.Update()
	update := bson.M{"$set": i}
	err = ContingentCollection.UpdateOne(ctx, bson.M{"_id": i.ID}, update)
	return
}

func (i *Contingent) Delete(ctx context.Context, filter bson.M) (err error) {
	err = ContingentCollection.DeleteOne(ctx, filter)
	return
}

func (i *ContingentQuery) List(ctx context.Context) (r *vcapool.ContingentList, err error) {
	pipe := vcago.NewMongoPipe()
	pipe.Match((*vcapool.ContingentQuery)(i).Match())
	r = new(vcapool.ContingentList)

	return
}*/
