package models

import (
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/google/uuid"
)

type (
	Tour struct {
		ID        string        `json:"id" bson:"_id"`
		Name      string        `json:"name" bson:"name"`
		ArtistIDs []string      `json:"artist_ids" bson:"artist_ids"`
		Artists   []Artist      `json:"artists" bson:"artists"`
		Events    []Event       `json:"events" bson:"events"`
		Creator   User          `json:"creator" bson:"creator"`
		Modified  vmod.Modified `json:"modified" bson:"modified"`
	}

	TourCreate struct {
		Name      string        `json:"name" bson:"name"`
		ArtistIDs []string      `json:"artist_ids" bson:"artist_ids"`
		Creator   User          `json:"creator" bson:"creator"`
		Events    []EventCreate `json:"events" bson:"events"`
	}

	TourUpdate struct {
		ID        string   `json:"id" bson:"_id"`
		Name      string   `json:"name" bson:"name"`
		ArtistIDs []string `json:"artist_ids" bson:"artist_ids"`
	}

	TourDatabase struct {
		ID        string        `json:"id" bson:"_id"`
		Name      string        `json:"name" bson:"name"`
		ArtistIDs []string      `json:"artist_ids" bson:"artist_ids"`
		Creator   string        `json:"creator" bson:"creator"`
		Modified  vmod.Modified `json:"modified" bson:"modified"`
	}
)

func (i *TourCreate) TourDatabase(token *vcapool.AccessToken) *TourDatabase {
	return &TourDatabase{
		ID:        uuid.NewString(),
		Name:      i.Name,
		ArtistIDs: i.ArtistIDs,
		Creator:   *&token.ID,
		Modified:  vmod.NewModified(),
	}
}
