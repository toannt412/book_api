package serialize

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	BookName          string             `json:"bookName,omitempty" bson:"bookName,omitempty"`
	Price             float64            `json:"price,omitempty" bson:"price,omitempty"`
	PublishingCompany string             `json:"publishingCompany,omitempty" bson:"publishingCompany,omitempty"`
	PublicationDate   time.Time          `json:"publicationDate,omitempty" bson:"publicationDate,omitempty"`
	Description       string             `json:"description,omitempty" bson:"description,omitempty"`
	CategoryIDs       []Category         `json:"categoryIds" bson:"categoryIds,omitempty"`
	AuthorID          []Author             `json:"authorId" bson:"authorId,omitempty"`
}

type Category struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	CatName string             `json:"categoryname" bson:"categoryname"`
}

type Author struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	AuthorName  string             `json:"authorName,omitempty" bson:"authorName,omitempty"`
	DateOfBirth time.Time          `json:"dateOfBirth,omitempty" bson:"dateOfBirth,omitempty"`
	HomeTown    string             `json:"homeTown,omitempty" bson:"homeTown,omitempty"`
	Alive       bool               `json:"alive,omitempty" bson:"alive,omitempty"`
}
