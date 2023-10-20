package serialize

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	Id            primitive.ObjectID `json:"id" `
	UserID        primitive.ObjectID `json:"userId" `
	Books         []CartBook         `json:"books" `
	TotalQuantity int                `json:"totalQuantity" `
	TotalAmount   float64            `json:"totalAmount" `
}

type CartBook struct {
	BookID   primitive.ObjectID `json:"bookId" `
	BookName string             `json:"bookName" `
	Price    float64            `json:"price" `
	Quantity int                `json:"quantity" `
	Total    float64            `json:"total" `
}
