package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	FullName    string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Location    string             `json:"location,omitempty" bson:"location,omitempty" `
	DateOfBirth time.Time          `json:"dateofbirth,omitempty" bson:"dateofbirth,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	UserName    string             `json:"username,omitempty" bson:"username,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
}
