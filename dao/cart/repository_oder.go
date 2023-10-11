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

func (repo *OrderRepository) CreateOrder(cxt context.Context, order *serialize.Order) (model.Order, error) {
	result, err := repo.ordersCollection.InsertOne(cxt, order)
	if err != nil {
		return model.Order{}, err
	}
	if result.InsertedID != nil {
		err := repo.ordersCollection.FindOne(cxt, bson.M{"_id": result.InsertedID}).Decode(&order)
		if err != nil {
			return model.Order{}, err
		}
	}
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
	return model.Order{
		Id:            order.Id,
		UserID:        order.UserID,
		Books:         bookSlice,
		CartID:        order.CartID,
		TotalQuantity: order.TotalQuantity,
		TotalPrice:    order.TotalPrice,
		TotalAmount:   order.TotalAmount,
		Status:        order.Status,
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

func (repo *OrderRepository) EditOrder(cxt context.Context, orderID string, order *serialize.Order) (model.Order, error) {
	objID, _ := primitive.ObjectIDFromHex(orderID)
	result, err := repo.ordersCollection.UpdateOne(cxt, bson.M{"_id": objID}, bson.M{"$set": order})
	if err != nil {
		return model.Order{}, err
	}
	var updatedOrder model.Order
	if result.MatchedCount == 1 {
		err := repo.ordersCollection.FindOne(cxt, bson.M{"_id": objID}).Decode(&updatedOrder)
		if err != nil {
			return model.Order{}, err
		}
	}
	return model.Order{
		Id:            updatedOrder.Id,
		UserID:        updatedOrder.UserID,
		Books:         updatedOrder.Books,
		CartID:        updatedOrder.CartID,
		TotalQuantity: updatedOrder.TotalQuantity,
		TotalPrice:    updatedOrder.TotalPrice,
		TotalAmount:   updatedOrder.TotalAmount,
		Status:        updatedOrder.Status,
	}, nil
}
