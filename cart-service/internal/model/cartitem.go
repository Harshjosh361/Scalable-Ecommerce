package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	ProductID  primitive.ObjectID `json:"productId" bson:"product_id"`
	UnitPrice  float64            `json:"unitPrice" bson:"unit_price"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	TotalPrice float64            `json:"totalPrice" bson:"total_price"`
}
