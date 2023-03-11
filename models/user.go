package models

import "time"

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Name      string    `bson:"name" json:",omitempty"`
	LastName  string    `bson:"lastName" json:"lastName,omitempty"`
	BirthDate time.Time `bson:"birth_date" json:"birthDate,omitempty"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password,omitempty"`
	Avatar    string    `bson:"avatar" json:"avatar,omitempty"`
	Banner    string    `bson:"banner" json:"banner,omitempty"`
	Location  string    `bson:"location" json:"location,omitempty"`
	WebSite   string    `bson:"webSite" json:"webSite,omitempty"`
}