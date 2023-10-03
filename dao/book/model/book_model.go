package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	BookName          string             `json:"bookName" bson:"bookName,omitempty"`
	Price             float64            `json:"price" bson:"price,omitempty"`
	PublishingCompany string             `json:"publishingCompany" bson:"publishingCompany,omitempty"`
	PublicationDate   time.Time          `json:"publicationDate" bson:"publicationDate,omitempty"`
	Description       string             `json:"description,omitempty" bson:"description,omitempty"`
	CategoryIDs       []Category         `json:"categoryIds" bson:"categoryIds"`
	AuthorID          Author             `json:"authorId" bson:"authorId"`
}

type Category struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	CatName string             `json:"categoryname" bson:"categoryname"`
}

type Author struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	AuthorName  string             `json:"authorName" bson:"authorName,omitempty"`
	DateOfBirth time.Time          `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	HomeTown    string             `json:"homeTown" bson:"homeTown,omitempty"`
	Alive       bool               `json:"alive" bson:"alive,omitempty"`
}
