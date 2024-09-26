package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id           primitive.ObjectID `json:"id"`
	UserID       primitive.ObjectID `json:"userid"`
	Items        []OrderItem        `json:"items"`
	TotalAmount  float64            `json:"totalamount"`
	CreatedAt    time.Time          `json:"createdAt"`
	ShippingInfo ShippingInfo       `json:"shippingInfo"`
}

type OrderItem struct {
	ProductID   primitive.ObjectID `bson:"product_id" json:"product_id"`
	ProductName string             `bson:"product_name" json:"product_name"`
	Quantity    int                `bson:"quantity" json:"quantity"`
	Price       float64            `bson:"price" json:"price"`
}

type ShippingInfo struct {
	Address    string `bson:"address" json:"address"`
	City       string `bson:"city" json:"city"`
	State      string `bson:"state" json:"state"`
	PostalCode string `bson:"postal_code" json:"postal_code"`
	Country    string `bson:"country" json:"country"`
}
