package serialize

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id"`
	FullName    string             `json:"fullname,omitempty"`
	Location    string             `json:"location,omitempty"`
	DateOfBirth time.Time          `json:"dateofbirth,omitempty"`
	Phone       string             `json:"phone,omitempty"`
	UserName    string             `json:"username,omitempty"`
	Password    string             `json:"password,omitempty"`
	Email       string             `json:"email,omitempty"`
}
