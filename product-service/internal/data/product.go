package data

import (
	"context"
	"errors"
	model "harsh/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductStore struct {
	Collection *mongo.Collection
}

func NewProductStore(db *mongo.Database) *ProductStore {
	return &ProductStore{
		Collection: db.Collection("products"),
	}
}
func (p *ProductStore) GetAllProduct(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	cursor, err := p.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product *model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductStore) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	filter := bson.M{"_id": objectID}
	if err := p.Collection.FindOne(ctx, filter).Decode(&product); err != nil {
		return &product, err
	}
	return &product, nil
}

func (p *ProductStore) DeleteProduct(ctx context.Context, id string) error {
	// converting id to object id
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	filter := bson.M{"_id": objectID}
	res, err := p.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("cannot delete product")
	}
	if res.DeletedCount == 0 {
		return err
	}
	return nil
}

func (p *ProductStore) CreateProduct(ctx context.Context, product *model.Product) error {
	_, err := p.Collection.InsertOne(ctx, product)
	if err != nil {
		return errors.New("failed to create product")
	}
	return nil
}

func (p *ProductStore) ModifyProduct(ctx context.Context, newProduct *model.Product) error {
	// update condition
	update := bson.M{
		"$set": bson.M{
			"name":        newProduct.Name,
			"price":       newProduct.Price,
			"description": newProduct.Description,
			"category":    newProduct.Category,
			"stock":       newProduct.Stock,
		},
	}
	// modifying product by id
	res, err := p.Collection.UpdateByID(ctx, newProduct.Id, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("no document found")
	}
	return nil
}

func (p *ProductStore) FindProductByName(ctx context.Context, name string) (*model.Product, error) {
	var exist *model.Product
	filter := bson.M{"name": name}
	err := p.Collection.FindOne(ctx, filter).Decode(&exist)
	if err != nil {
		return nil, errors.New("no product found")
	}
	return exist, nil
}
