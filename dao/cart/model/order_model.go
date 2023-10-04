package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	UserID        primitive.ObjectID `json:"userID" bson:"userID"`
	Books         []OrderBook        `json:"books" bson:"books"`
	CartID        []Cart             `json:"cartID" bson:"cartID"`
	TotalQuantity int                `json:"totalQuantity" bson:"totalQuantity"`
	TotalPrice    float64            `json:"totalPrice" bson:"totalPrice"`
	TotalAmount   float64            `json:"totalAmount" bson:"totalAmount"`
	OrderDate     time.Time          `json:"orderDate" bson:"orderDate"`
	Status        string             `json:"status" bson:"status"`
}

type OrderBook struct {
	BookID   primitive.ObjectID `json:"bookId" bson:"bookId"`
	BookName string             `json:"bookName" bson:"bookName"`
	Price    float64            `json:"price" bson:"price"`
	Quantity int                `json:"quantity" bson:"quantity"`
	Total    float64            `json:"total" bson:"total"`
}
