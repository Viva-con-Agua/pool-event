package models

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/vmdb"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	ParticipationCreate struct {
		EventID string `json:"event_id" bson:"event_id"`
		Comment string `json:"comment" bson:"comment"`
	}

	ParticipationUpdate struct {
		ID      string `json:"id" bson:"_id"`
		Status  string `json:"status" bson:"status"`
		Comment string `json:"comment" bson:"comment"`
		//Confirmer UserInternal `json:"confirmer" bson:"confirmer"`
	}
	ParticipationDatabase struct {
		ID      string     `json:"id" bson:"_id"`
		UserID  string     `json:"user" bson:"user"`
		EventID string     `json:"event_id" bson:"event_id"`
		Comment string     `json:"comment" bson:"comment"`
		Status  string     `json:"status" bson:"status"`
		Crew    CrewSimple `json:"crew" bson:"crew"`
		//Confirmer UserInternal   `json:"confirmer" bson:"confirmer"`
		Modified vcago.Modified `json:"modified" bson:"modified"`
	}
	Participation struct {
		ID      string     `json:"id" bson:"_id"`
		User    User       `json:"user" bson:"user"`
		EventID string     `json:"event_id" bson:"event_id"`
		Comment string     `json:"comment" bson:"comment"`
		Status  string     `json:"status" bson:"status"`
		Event   Event      `json:"event" bson:"event"`
		Crew    CrewSimple `json:"crew" bson:"crew"`
		//Confirmer UserInternal   `json:"confirmer" bson:"confirmer"`
		Modified vcago.Modified `json:"modified" bson:"modified"`
	}

	ParticipationParam struct {
		ID string `param:"id"`
	}

	ParticipationQuery struct {
		ID       []string `query:"id" qs:"id"`
		EventID  []string `query:"event_id" qs:"event_id"`
		Comment  []string `query:"comment" bson:"comment"`
		Status   []string `query:"status" bson:"status"`
		UserId   []string `query:"user_id" bson:"user_id"`
		CrewName []string `query:"crew_name" bson:"crew_name"`
		CrewId   []string `query:"crew_id" qs:"crew_id"`
	}
	ParticipationStateRequest struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
)

func (i *ParticipationCreate) ParticipationDatabase(token *vcapool.AccessToken) *ParticipationDatabase {
	return &ParticipationDatabase{
		ID:       uuid.NewString(),
		UserID:   token.ID,
		EventID:  i.EventID,
		Comment:  i.Comment,
		Status:   "requested",
		Crew:     CrewSimple(*token.CrewSimple()),
		Modified: vcago.NewModified(),
	}
}

func ParticipationPipeline() (pipe *vmdb.Pipeline) {
	pipe = vmdb.NewPipeline()
	pipe.LookupUnwind("users", "user_id", "_id", "user")
	pipe.LookupUnwind("events", "event_id", "_id", "event")
	return
}

func (i *ParticipationQuery) Match() (r *vmdb.Match) {
	r = vmdb.NewMatch()
	r.EqualStringList("_id", i.ID)
	r.EqualStringList("event_id", i.EventID)
	r.EqualStringList("status", i.Status)
	r.EqualStringList("comment", i.Comment)
	r.EqualStringList("user._id", i.UserId)
	r.EqualStringList("crew.name", i.CrewName)
	r.EqualStringList("crew.id", i.CrewId)
	return
}

func (i *ParticipationParam) Match() (r *vmdb.Match) {
	r = vmdb.NewMatch()
	r.EqualString("_id", i.ID)
	return
}

func (i *ParticipationDatabase) Match() (r *vmdb.Match) {
	r = vmdb.NewMatch()
	r.EqualString("_id", i.ID)
	return
}

func (i *ParticipationUpdate) Match() *vmdb.Match {
	match := vmdb.NewMatch()
	match.EqualString("_id", i.ID)
	return match
}

func (i *ParticipationUpdate) Filter() bson.D {
	return bson.D{{Key: "_id", Value: i.ID}}
}

func (i *ParticipationParam) Filter() bson.D {
	return bson.D{{Key: "_id", Value: i.ID}}
}

func (i *ParticipationStateRequest) Permission(token *vcapool.AccessToken) bson.D {
	if i.Status == "withdrawn" {
		return bson.D{{Key: "_id", Value: i.ID}, {Key: "user", Value: token.ID}}
	} else if i.Status == "confirmed" || i.Status == "rejected" {
		if token.Roles.Validate("employee;admin") {
			return bson.D{{Key: "_id", Value: i.ID}}
		} else if token.Roles.Validate("operation") {
			return bson.D{{Key: "_id", Value: i.ID}, {Key: "crew.id", Value: token.CrewID}}
		}
	}
	return bson.D{{Key: "_id", Value: "not_defined"}}
}

func (i *ParticipationStateRequest) Match() (r *vmdb.Match) {
	r = vmdb.NewMatch()
	r.EqualString("_id", i.ID)
	return
}
