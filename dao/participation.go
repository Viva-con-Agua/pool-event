package dao

import (
	"context"
	"log"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"go.mongodb.org/mongo-driver/bson"
)

type ParticipationCreate struct {
	vcapool.ParticipationCreate
}

type ParticipationUpdate struct {
	vcapool.ParticipationUpdate
}

type ParticipationParam struct {
	vcapool.ParticipationParam
}

type ParticipationQuery struct {
	vcapool.ParticipationQuery
}

var ParticipationCollection = Database.Collection("participations").CreateMultiIndex(bson.D{{Key: "user_id", Value: 1}, {Key: "event_id", Value: 1}}, true)

func (i *ParticipationCreate) Create(ctx context.Context, token *vcapool.AccessToken) (r *vcapool.Participation, err error) {
	event := new(vcapool.Event)
	if err = EventCollection.FindOne(ctx, bson.M{"_id": i.EventID}, event); err != nil {
		return
	}
	r = vcapool.NewParticipation(token, i.EventID)
	err = ParticipationCollection.InsertOne(ctx, r)
	return
}

func (i *ParticipationUpdate) Update(ctx context.Context, token *vcapool.AccessToken) (r *vcapool.Participation, err error) {
	r = new(vcapool.Participation)
	if err = ParticipationCollection.FindOne(ctx, bson.M{"_id": i.ID}, r); err != nil {
		return
	}
	event := new(vcapool.Event)
	if err = EventCollection.FindOne(ctx, bson.M{"_id": r.EventID}, event); err != nil {
		return
	}
	if i.Status == "confirmed" || i.Status == "rejected" {
		log.Print(i.Status)
		if !(token.PoolRoles.Validate("operation") || token.Roles.Validate("employee;admin") || token.ID == event.EventASP.ID) {
			err = vcago.NewPermissionDenied("participation", i.ID)
			log.Print("asd")
			return
		}
	} else if i.Status == "withdrawn" {
		if token.ID != r.User.ID {
			err = vcago.NewPermissionDenied("participation", i.ID)
			return
		}
	}
	if err = ParticipationCollection.UpdateOneSet(ctx, bson.M{"_id": i.ID}, i.ParticipationUpdate); err != nil {
		return
	}
	err = ParticipationCollection.FindOne(ctx, bson.M{"_id": i.ID}, r)
	return
}

func (i *ParticipationParam) Get(ctx context.Context, token *vcapool.AccessToken) (r *vcapool.Participation, err error) {
	r = new(vcapool.Participation)
	if err = ParticipationCollection.FindOne(ctx, bson.M{"_id": i.ID}, r); err != nil {
		return
	}
	event := new(vcapool.Event)
	if err = EventCollection.FindOne(ctx, bson.M{"_id": r.EventID}, event); err != nil {
		return
	}
	if !(token.PoolRoles.Validate("operation") || token.Roles.Validate("employee;admin") || token.ID == event.EventASP.ID || token.ID == r.User.ID) {
		err = vcago.NewPermissionDenied("participation", i.ID)
		return
	}
	return
}

func (i *ParticipationParam) Delete(ctx context.Context, token *vcapool.AccessToken) (err error) {
	if !token.Roles.Validate("employee;admin") {
		err = vcago.NewPermissionDenied("participation", i.ID)
		return
	}
	err = ParticipationCollection.DeleteOne(ctx, bson.M{"_id": i.ID})
	return
}

func (i *ParticipationQuery) List(ctx context.Context, token *vcapool.AccessToken) (r *vcapool.ParticipationList, err error) {
	pipe := vcago.NewMongoPipe()
	pipe.Match(i.Match())
	r = new(vcapool.ParticipationList)
	err = ParticipationCollection.Aggregate(ctx, pipe.Pipe, r)
	return
}
