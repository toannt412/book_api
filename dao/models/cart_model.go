package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	UserID        primitive.ObjectID `json:"userId" bson:"userId"`
	Books         []CartBook         `json:"books" bson:"books"`
	TotalQuantity int                `json:"totalQuantity" bson:"totalQuantity"`
	TotalAmount   float64            `json:"totalAmount" bson:"totalAmount"`
}

type CartBook struct {
	BookID   primitive.ObjectID `json:"bookId" bson:"bookId"`
	BookName string             `json:"bookName" bson:"bookName"`
	Price    float64            `json:"price" bson:"price"`
	Quantity int                `json:"quantity" bson:"quantity"`
	Total    float64            `json:"total" bson:"total"`
}
