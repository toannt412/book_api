package serialize

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id            primitive.ObjectID `json:"id" `
	UserID        primitive.ObjectID `json:"userID" `
	Books         []OrderBook        `json:"books" `
	CartID        primitive.ObjectID `json:"cartID" `
	TotalQuantity int                `json:"totalQuantity" `
	TotalPrice    float64            `json:"totalPrice" `
	TotalAmount   float64            `json:"totalAmount" `
	OrderDate     time.Time          `json:"orderDate" `
	Status        string             `json:"status" `
}

type OrderBook struct {
	BookID   primitive.ObjectID `json:"bookId" `
	BookName string             `json:"bookName" `
	Price    float64            `json:"price" `
	Quantity int                `json:"quantity" `
	Total    float64            `json:"total" `
}
