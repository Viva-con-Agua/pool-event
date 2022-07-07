package models

type (
	Location struct {
		Name        string   `json:"name" bson:"name"`
		Street      string   `json:"street" bson:"street"`
		City        string   `json:"city" bson:"city"`
		Country     string   `json:"country" bson:"country"`
		CountryCode string   `json:"country_code" bson:"country_code"`
		Number      string   `json:"number" bson:"number"`
		Position    Position `json:"position" bson:"position"`
		PlaceID     string   `json:"place_id" bson:"place_id"`
		Sublocality string   `json:"sublocality" bson:"sublocality"`
	}
	Position struct {
		Lat float64 `json:"lat" bson:"lat"`
		Lng float64 `json:"lng" bson:"lng"`
	}
)
