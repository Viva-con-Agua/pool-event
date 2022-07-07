package models

import "github.com/Viva-con-Agua/vcago"

type (
	User struct {
		ID          string  `json:"id,omitempty" bson:"_id"`
		Email       string  `json:"email" bson:"email"`
		FirstName   string  `bson:"first_name" json:"first_name"`
		LastName    string  `bson:"last_name" json:"last_name"`
		FullName    string  `bson:"full_name" json:"full_name"`
		DisplayName string  `bson:"display_name" json:"display_name"`
		Profile     Profile `json:"profile" bson:"profile"`
	}
	Profile struct {
		ID        string         `bson:"_id" json:"id"`
		Gender    string         `bson:"gender" json:"gender"`
		Phone     string         `bson:"phone" json:"phone"`
		Birthdate int64          `bson:"birthdate" json:"birthdate"`
		UserID    string         `bson:"user_id" json:"user_id"`
		Modified  vcago.Modified `bson:"modified" json:"modified"`
	}
	//ExternalASP represents an external asp without user_id.
	UserExternal struct {
		FullName    string `json:"full_name" bson:"full_name"`
		DisplayName string `json:"display_name" bson:"display_name"`
		Email       string `json:"email" bson:"email"`
		Phone       string `json:"phone" bson:"phone"`
	}
)
