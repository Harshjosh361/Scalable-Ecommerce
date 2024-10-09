package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name"`
	Price       float64            `bson:"price"`
	Description string             `bson:"description"`
	Category    string             `bson:"category"`
	Stock       int                `bson:"stock"`
}
