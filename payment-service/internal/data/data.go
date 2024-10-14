package data

import (
	"context"
	model "harsh/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentData struct {
	Collection *mongo.Collection
}

func NewPaymentData(db *mongo.Database) *PaymentData {
	return &PaymentData{
		Collection: db.Collection("payments"),
	}
}

func (p *PaymentData) SavePayment(ctx context.Context, payment *model.Payment) error {
	_, err := p.Collection.InsertOne(ctx, payment)
	return err
}

func (p *PaymentData) GetPaymentById(ctx context.Context, id primitive.ObjectID) (*model.Payment, error) {

	var payment *model.Payment
	err := p.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
