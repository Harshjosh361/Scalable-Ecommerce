package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	OrderID  primitive.ObjectID `json:"order_id" bson:"order_id"`
	Amount   float64            `json:"amount" bson:"amount"`
	Currency string             `json:"currency" bson:"currency"`
	Status   string             `json:"status" bson:"status"`
	// Token     string             `json:"token" bson:"token"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
