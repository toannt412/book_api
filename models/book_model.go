package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	BookName          string             `json:"bookName" bson:"bookName,omitempty"`
	Price             int                `json:"price" bson:"price"`
	Author            string             `json:"author" bson:"author"`
	Category          string             `json:"category" bson:"category"`
	PublishingCompany string             `json:"publishingCompany" bson:"publishingCompany"`
	PublicationDate   time.Time          `json:"publicationDate" bson:"publicationDate"`
	Description       *string            `json:"description,omitempty" bson:"description,omitempty"`
}
