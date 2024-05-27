package cart

import (
	"bookstore/dao"
	"bookstore/dao/cart/model"
	"bookstore/serialize"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	ordersCollection *mongo.Collection
}

func NewOrderRepository() *OrderRepository {
	var DB *mongo.Client = dao.ConnectDB()
	return &OrderRepository{
		ordersCollection: dao.GetCollection(DB, "orders"),
	}
}

func (repo *OrderRepository) CreateOrder(cxt context.Context, order *serialize.Order) (*serialize.Order, error) {
	bookSlice := make([]model.OrderBook, len(order.Books))
	for i, book := range order.Books {
		bookSlice[i] = model.OrderBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	model := model.Order{
		Id:            order.Id,
		UserID:        order.UserID,
		Books:         bookSlice,
		CartID:        order.CartID,
		TotalQuantity: order.TotalQuantity,
		TotalPrice:    order.TotalPrice,
		TotalAmount:   order.TotalAmount,
		OrderDate:     order.OrderDate,
		Status:        order.Status,
	}
	result, err := repo.ordersCollection.InsertOne(cxt, model)
	if err != nil {
		return &serialize.Order{}, err
	}
	if result.InsertedID != nil {
		err := repo.ordersCollection.FindOne(cxt, bson.M{"_id": result.InsertedID}).Decode(&model)
		if err != nil {
			return &serialize.Order{}, err
		}
	}
	books := make([]serialize.OrderBook, len(model.Books))
	for i, book := range model.Books {
		books[i] = serialize.OrderBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	return &serialize.Order{
		Id:            model.Id,
		UserID:        model.UserID,
		Books:         books,
		CartID:        model.CartID,
		TotalQuantity: model.TotalQuantity,
		TotalPrice:    model.TotalPrice,
		TotalAmount:   model.TotalAmount,
		OrderDate:     model.OrderDate,
		Status:        model.Status,
	}, nil
}

func (repo *OrderRepository) GetOrderByID(cxt context.Context, orderID string) (model.Order, error) {
	var order model.Order
	objID, _ := primitive.ObjectIDFromHex(orderID)
	err := repo.ordersCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (repo *OrderRepository) DeleteOrder(cxt context.Context, orderID string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(orderID)
	result, err := repo.ordersCollection.DeleteOne(cxt, bson.M{"_id": objID})
	if err != nil {
		return "Deleted failed", err
	}
	if result.DeletedCount == 0 {
		return "", mongo.ErrNoDocuments
	}
	return "Deleted successfully", nil
}

func (repo *OrderRepository) EditOrder(cxt context.Context, orderID string, order *serialize.Order) (*serialize.Order, error) {
	objID, _ := primitive.ObjectIDFromHex(orderID)
	bookSlice := make([]model.OrderBook, len(order.Books))
	for i, book := range order.Books {
		bookSlice[i] = model.OrderBook{
			BookID:   book.BookID,
			BookName: book.BookName,
			Price:    book.Price,
			Quantity: book.Quantity,
			Total:    book.Total,
		}
	}
	model := model.Order{
		Id:            objID,
		UserID:        order.UserID,
		Books:         bookSlice,
		CartID:        order.CartID,
		TotalQuantity: order.TotalQuantity,
		TotalPrice:    order.TotalPrice,
		TotalAmount:   order.TotalAmount,
		OrderDate:     order.OrderDate,
		Status:        order.Status,
	}
	result, err := repo.ordersCollection.UpdateOne(cxt, bson.M{"_id": objID}, bson.M{"$set": model})
	if err != nil {
		return &serialize.Order{}, err
	}
	var updatedOrder *serialize.Order
	if result.MatchedCount == 1 {
		err := repo.ordersCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&updatedOrder)
		if err != nil {
			return &serialize.Order{}, err
		}
	}
	return &serialize.Order{
		Id:            objID,
		UserID:        updatedOrder.UserID,
		Books:         updatedOrder.Books,
		CartID:        updatedOrder.CartID,
		TotalQuantity: updatedOrder.TotalQuantity,
		TotalPrice:    updatedOrder.TotalPrice,
		TotalAmount:   updatedOrder.TotalAmount,
		OrderDate:     updatedOrder.OrderDate,
		Status:        updatedOrder.Status,
	}, nil
}
