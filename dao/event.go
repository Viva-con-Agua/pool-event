package dao

import (
	"context"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type EventDB struct {
	ID                    string                   `json:"id" bson:"_id"`
	Name                  string                   `json:"name" bson:"name"`
	TypeOfEvent           string                   `json:"type_of_event" bson:"type_of_event"`
	AdditionalInformation string                   `json:"additional_information" bson:"additional_information"`
	Website               string                   `json:"website" bson:"website"`
	TourneeID             string                   `json:"tournee_id" bson:"tournee_id"`
	Location              vcapool.Location         `json:"location" bson:"location"`
	Artists               []string                 `json:"artists" bson:"artists"`
	Organizer             vcapool.Organizer        `json:"organizer" bson:"organizer"`
	StartAt               int64                    `json:"start_at" bson:"start_at"`
	EndAt                 int64                    `json:"end_at" bson:"end_at"`
	Crew                  vcapool.CrewSimple       `json:"crew" bson:"crew"`
	EventASP              vcapool.EventASP         `json:"event_asp" bson:"event_asp"`
	InteralASP            vcapool.EventASP         `json:"interal_asp" bson:"internal_asp"`
	ExternalASP           vcapool.EventASPExternal `json:"external_asp" bson:"external_asp"`
	Application           vcapool.EventApplication `json:"application" bson:"application"`
	EventTools            vcapool.EventTools       `json:"event_tools" bson:"event_tools"`
	CreatorID             string                   `json:"creator_id" bson:"creator_id"`
	EventState            vcapool.EventState       `json:"event_state" bson:"event_state"`
	Modified              vcago.Modified           `json:"modified" bson:"modified"`
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
