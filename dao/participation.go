package dao

import (
	"context"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
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
	r = &vcapool.Participation{
		ID: uuid.NewString(),
		User: vcapool.UserInternal{
			UserID:      token.ID,
			Email:       token.Email,
			FullName:    token.FullName,
			DisplayName: token.DisplayName,
			Phone:       token.Profile.Phone,
		},
		EventID:  i.EventID,
		Status:   "requested",
		Modified: vcago.NewModified(),
	}
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
		if !(token.PoolRoles.Validate("operation") || token.Roles.Validate("employee;admin") || token.ID == event.EventASP.UserID) {
			err = vcago.NewPermissionDenied("participation", i.ID)
			return
		}
	} else if i.Status == "withdrawn" {
		if token.ID != r.User.UserID {
			err = vcago.NewPermissionDenied("participation", i.ID)
			return
		}
	}
	if err = ParticipationCollection.UpdateOneSet(ctx, bson.M{"_id": i.ID}, i); err != nil {
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
	if !(token.PoolRoles.Validate("operation") || token.Roles.Validate("employee;admin") || token.ID == event.EventASP.UserID || token.ID == r.User.UserID) {
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
