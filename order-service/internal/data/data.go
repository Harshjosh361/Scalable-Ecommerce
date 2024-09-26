package data

import (
	"context"
	"harshy/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderData struct {
	Collection *mongo.Collection
}

func NewOrderData(db *mongo.Database) *OrderData {
	return &OrderData{
		Collection: db.Collection("orders"),
	}
}

func (o *OrderData) CreateOrder(ctx context.Context, order *models.Order) error {
	order.CreatedAt = time.Now()
	_, err := o.Collection.InsertOne(ctx, order)
	return err
}

func (o *OrderData) GetOrderByUserId(ctx context.Context, userId primitive.ObjectID) ([]*models.Order, error) {

	var orders []*models.Order
	filter := bson.M{"userid": userId}
	cursor, err := o.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// iterate over cursor
	for cursor.Next(ctx) {
		var order *models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (o *OrderData) GetOrderById(ctx context.Context, id primitive.ObjectID) (*models.Order, error) {
	var order models.Order
	err := o.Collection.FindOne(ctx, bson.M{"id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
