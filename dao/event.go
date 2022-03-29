package dao

import (
	"context"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type EventDB struct {
	ID           string                        `json:"id" bson:"_id"`
	Name         string                        `json:"name" bson:"name"`
	TypeOfEvent  string                        `json:"type_of_event" bson:"type_of_event"`
	TourneeID    string                        `json:"tournee_id" bson:"tournee_id"`
	Location     vcapool.Location              `json:"location" bson:"location"`
	Artists      []string                      `json:"artists" bson:"artists"`
	Organizer    vcapool.Organizer             `json:"organizer" bson:"organizer"`
	StartAt      int64                         `json:"start_at" bson:"start_at"`
	EndAt        int64                         `json:"end_at" bson:"end_at"`
	Application  vcapool.EventApplication      `json:"application" bson:"application"`
	Organisation vcapool.EventOrganisationList `json:"organistion"`
	Tools        []string                      `json:"tools" bson:"tools"`
	CreaterID    string                        `json:"creater_id" bson:"creater_id"`
	Modified     vcago.Modified                `json:"modified" bson:"modified"`
}

var EventCollection = Database.Collection("events")

type EventQuery vcapool.EventQuery

type Event vcapool.Event

func (i *EventDB) Create(ctx context.Context) (err error) {
	i.ID = uuid.NewString()
	i.Modified = vcago.NewModified()
	err = EventCollection.InsertOne(ctx, i)
	return
}

func (i *Event) Get(ctx context.Context, filter bson.M) (err error) {
	err = EventCollection.FindOne(ctx, filter, i)
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
	pipe.Match((*vcapool.EventQuery)(i).Match())
	r = new(vcapool.EventList)
	err = EventCollection.Aggregate(ctx, pipe.Pipe, r)
	return
}
