package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	BookName          string             `json:"bookname" bson:"bookname,omitempty"`
	Price             float64            `json:"price" bson:"price,omitempty"`
	PublishingCompany string             `json:"publishingcompany" bson:"publishingcompany,omitempty"`
	PublicationDate   time.Time          `json:"publicationdate" bson:"publicationdate,omitempty"`
	Description       string             `json:"description,omitempty" bson:"description,omitempty"`
	CategoryIDs       []Category         `json:"categoryIds" bson:"categoryIds,omitempty"`
	AuthorID          []Author           `json:"authorId" bson:"authorId,omitempty"`
}

type Category struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	CatName string             `json:"categoryname" bson:"categoryname"`
}

type Author struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	AuthorName  string             `json:"authorname" bson:"authorname,omitempty"`
	DateOfBirth time.Time          `json:"dateofbirth" bson:"dateofbirth,omitempty"`
	HomeTown    string             `json:"hometown" bson:"hometown,omitempty"`
	Alive       bool               `json:"alive" bson:"alive"`
}
