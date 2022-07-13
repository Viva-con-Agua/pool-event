package models

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/vmdb"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	EventCreate struct {
		Name                  string           `json:"name" bson:"name"`
		TypeOfEvent           string           `json:"type_of_event" bson:"type_of_event"`
		AdditionalInformation string           `json:"additional_information" bson:"additional_information"`
		Website               string           `json:"website" bson:"website"`
		TourID                string           `json:"tour_id" bson:"tour_id"`
		Location              Location         `json:"location" bson:"location"`
		ArtistIDs             []string         `json:"artist_ids" bson:"artist_ids"`
		Organizer             Organizer        `json:"organizer" bson:"organizer"`
		StartAt               int64            `json:"start_at" bson:"start_at"`
		EndAt                 int64            `json:"end_at" bson:"end_at"`
		Crew                  CrewSimple       `json:"crew" bson:"crew"`
		EventASP              User             `json:"event_asp" bson:"event_asp"`
		InternalASP           User             `json:"internal_asp" bson:"internal_asp"`
		ExternalASP           UserExternal     `json:"external_asp" bson:"external_asp"`
		Application           EventApplication `json:"application" bson:"application"`
		EventTools            EventTools       `json:"event_tools" bson:"event_tools"`
	}
	EventDatabase struct {
		ID                    string           `json:"id" bson:"_id"`
		Name                  string           `json:"name" bson:"name"`
		TypeOfEvent           string           `json:"type_of_event" bson:"type_of_event"`
		AdditionalInformation string           `json:"additional_information" bson:"additional_information"`
		Website               string           `json:"website" bson:"website"`
		TourID                string           `json:"tour_id" bson:"tour_id"`
		Location              Location         `json:"location" bson:"location"`
		ArtistIDs             []string         `json:"artist_ids" bson:"artist_ids"`
		Organizer             string           `json:"organizer" bson:"organizer"`
		StartAt               int64            `json:"start_at" bson:"start_at"`
		EndAt                 int64            `json:"end_at" bson:"end_at"`
		Crew                  CrewSimple       `json:"crew" bson:"crew"`
		EventASPID            string           `json:"event_asp_id" bson:"event_asp_id"`
		InteralASPID          string           `json:"internal_asp_id" bson:"internal_asp_id"`
		ExternalASP           UserExternal     `json:"external_asp" bson:"external_asp"`
		Application           EventApplication `json:"application" bson:"application"`
		EventTools            EventTools       `json:"event_tools" bson:"event_tools"`
		CreatorID             string           `json:"creator_id" bson:"creator_id"`
		EventState            EventState       `json:"event_state" bson:"event_state"`
		Modified              vcago.Modified   `json:"modified" bson:"modified"`
	}
	EventTools struct {
		Tools   []string `json:"tools" bson:"tools"`
		Special string   `json:"special" bson:"special"`
	}
	//EventApplication represents the application type of an event.
	EventApplication struct {
		StartDate      int64 `json:"start_date" bson:"start_date"`
		EndDate        int64 `json:"end_date" bson:"end_date"`
		SupporterCount int   `json:"supporter_count" bson:"supporter_count"`
	}
	//EventState represents the state of an event.
	EventState struct {
		State                string `json:"state" bson:"state"`
		CrewConfirmation     string `json:"crew_confirmation" bson:"crew_confirmation"`
		InternalConfirmation string `json:"internal_confirmation" bson:"internal_confirmation"`
		TakingID             string `json:"taking_id" bson:"taking_id"`
		DepositID            string `json:"deposit_id" bson:"deposit_id"`
	}
	Event struct {
		ID                    string           `json:"id" bson:"_id"`
		Name                  string           `json:"name" bson:"name"`
		TypeOfEvent           string           `json:"type_of_event" bson:"type_of_event"`
		AdditionalInformation string           `json:"additional_information" bson:"additional_information"`
		Website               string           `json:"website" bson:"website"`
		TourID                string           `json:"tour_id" bson:"tour_id"`
		Location              Location         `json:"location" bson:"location"`
		ArtistIDs             []string         `json:"artist_ids" bson:"artist_ids"`
		Artists               []Artist         `json:"artists" bson:"artists"`
		Organizer             Organizer        `json:"organizer" bson:"organizer"`
		StartAt               int64            `json:"start_at" bson:"start_at"`
		EndAt                 int64            `json:"end_at" bson:"end_at"`
		Crew                  CrewSimple       `json:"crew" bson:"crew"`
		EventASP              User             `json:"event_asp" bson:"event_asp"`
		InteralASP            User             `json:"internal_asp" bson:"internal_asp"`
		ExternalASP           UserExternal     `json:"external_asp" bson:"external_asp"`
		Application           EventApplication `json:"application" bson:"application"`
		EventTools            EventTools       `json:"event_tools" bson:"event_tools"`
		Creator               User             `json:"creator" bson:"creator"`
		EventState            EventState       `json:"event_state" bson:"event_state"`
		Modified              vcago.Modified   `json:"modified" bson:"modified"`
	}
	EventUpdate struct {
		ID                    string           `json:"id" bson:"_id"`
		Name                  string           `json:"name" bson:"name"`
		TypeOfEvent           string           `json:"type_of_event" bson:"type_of_event"`
		AdditionalInformation string           `json:"additional_information" bson:"additional_information"`
		Website               string           `json:"website" bson:"website"`
		TourID                string           `json:"tour_id" bson:"tour_id"`
		Location              Location         `json:"location" bson:"location"`
		ArtistIDs             []string         `json:"artist_ids" bson:"artist_ids"`
		Organizer             Organizer        `json:"organizer" bson:"organizer"`
		StartAt               int64            `json:"start_at" bson:"start_at"`
		EndAt                 int64            `json:"end_at" bson:"end_at"`
		Crew                  CrewSimple       `json:"crew" bson:"crew"`
		EventASP              User             `json:"event_asp" bson:"event_asp"`
		InternalASP           User             `json:"internal_asp" bson:"internal_asp"`
		ExternalASP           UserExternal     `json:"external_asp" bson:"external_asp"`
		Application           EventApplication `json:"application" bson:"application"`
		EventTools            EventTools       `json:"event_tools" bson:"event_tools"`
		EventState            EventState       `json:"event_state" bson:"event_state"`
	}
	EventParam struct {
		ID string `param:"id"`
	}

	EventQuery struct {
		ID          string `query:"id" qs:"id"`
		Name        string `query:"name" qs:"name"`
		CrewID      string `query:"crew_id" qs:"crew_id"`
		UpdatedTo   string `query:"updated_to" qs:"updated_to"`
		UpdatedFrom string `query:"updated_from" qs:"updated_from"`
		CreatedTo   string `query:"created_to" qs:"created_to"`
		CreatedFrom string `query:"created_from" qs:"created_from"`
	}
)

func (i *EventCreate) EventDatabase(token *vcapool.AccessToken) *EventDatabase {
	return &EventDatabase{
		ID:                    uuid.NewString(),
		Name:                  i.Name,
		TypeOfEvent:           i.TypeOfEvent,
		AdditionalInformation: i.AdditionalInformation,
		Website:               i.Website,
		Location:              i.Location,
		ArtistIDs:             i.ArtistIDs,
		Organizer:             i.Organizer.ID,
		StartAt:               i.StartAt,
		EndAt:                 i.EndAt,
		Crew:                  i.Crew,
		EventASPID:            i.EventASP.ID,
		InteralASPID:          i.EventASP.ID,
		ExternalASP:           i.ExternalASP,
		Application:           i.Application,
		EventTools:            i.EventTools,
		CreatorID:             token.ID,
		EventState: EventState{
			State: "created",
		},
		Modified: vcago.NewModified(),
	}
}

func EventPipeline() (pipe *vmdb.Pipeline) {
	pipe = vmdb.NewPipeline()
	pipe.LookupUnwind("users", "event_asp_id", "_id", "event_asp")
	pipe.LookupUnwind("users", "internal_asp_id", "_id", "internal_asp")
	pipe.LookupUnwind("users", "creator_id", "_id", "creator")
	pipe.LookupUnwind("organizers", "organizer", "_id", "organizer")
	pipe.LookupList("artists", "artist_ids", "_id", "artists")
	return
}
func (i *EventDatabase) Match() *vmdb.Match {
	match := vmdb.NewMatch()
	match.EqualString("_id", i.ID)
	return match
}

func (i *EventParam) Match() *vmdb.Match {
	match := vmdb.NewMatch()
	match.EqualString("_id", i.ID)
	return match
}
func (i *EventUpdate) Match() *vmdb.Match {
	match := vmdb.NewMatch()
	match.EqualString("_id", i.ID)
	return match
}

func (i *EventUpdate) Filter() bson.D {
	return bson.D{{Key: "_id", Value: i.ID}}
}
func (i *EventParam) Filter() bson.D {
	return bson.D{{Key: "_id", Value: i.ID}}
}

func (i *EventQuery) Match() *vmdb.Match {
	match := vmdb.NewMatch()
	match.EqualString("_id", i.ID)
	match.LikeString("name", i.Name)
	match.GteInt64("modified.updated", i.UpdatedFrom)
	match.GteInt64("modified.created", i.CreatedFrom)
	match.LteInt64("modified.updated", i.UpdatedTo)
	match.LteInt64("modified.created", i.CreatedTo)
	return match
}
