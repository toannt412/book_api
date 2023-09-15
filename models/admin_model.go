package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	FullName string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
	UserName string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Role     string             `json:"role,omitempty" bson:"role,omitempty"`
}
