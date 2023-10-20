package serialize

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id                primitive.ObjectID `json:"id" `
	BookName          string             `json:"bookName,omitempty" `
	Price             float64            `json:"price,omitempty" `
	PublishingCompany string             `json:"publishingCompany,omitempty" `
	PublicationDate   time.Time          `json:"publicationDate,omitempty" `
	Description       string             `json:"description,omitempty" `
	CategoryIDs       []Category         `json:"categoryIds" `
	AuthorID          []Author           `json:"authorId" `
}

type Category struct {
	Id      primitive.ObjectID `json:"id" `
	CatName string             `json:"categoryname" `
}

type Author struct {
	Id          primitive.ObjectID `json:"id" `
	AuthorName  string             `json:"authorname" `
	DateOfBirth time.Time          `json:"dateofbirth" `
	HomeTown    string             `json:"hometown" `
	Alive       bool               `json:"alive" `
}
