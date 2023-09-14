package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	UserName    string             `json:"username,omitempty" bson:"username,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	FullName    string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	DateOfBirth time.Time          `json:"dateofbirth,omitempty" bson:"dateofbirth,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
}
