package serialize

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id                primitive.ObjectID `json:"id"`
	BookName          string             `json:"bookName,omitempty"`
	Price             float64            `json:"price,omitempty"`
	PublishingCompany string             `json:"publishingCompany,omitempty"`
	PublicationDate   time.Time          `json:"publicationDate,omitempty"`
	Description       string             `json:"description,omitempty"`
	CategoryIDs       []Category         `json:"categoryIds"`
	AuthorID          Author             `json:"authorid"`
}

type Category struct {
	Id      primitive.ObjectID `json:"id"`
	CatName string             `json:"categoryname"`
}

type Author struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	AuthorName  string             `json:"authorName,omitempty" bson:"authorName,omitempty"`
	DateOfBirth time.Time          `json:"dateOfBirth,omitempty" bson:"dateOfBirth,omitempty"`
	HomeTown    string             `json:"homeTown,omitempty" bson:"homeTown,omitempty"`
	Alive       bool               `json:"alive,omitempty" bson:"alive,omitempty"`
}
