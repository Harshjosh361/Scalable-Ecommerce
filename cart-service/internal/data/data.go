package data

import (
	"context"
	"errors"
	"harsh/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartStore struct {
	Collection *mongo.Collection
}

func NewCartStore(db *mongo.Database) *CartStore {
	return &CartStore{
		Collection: db.Collection("cart"),
	}
}

// create cart insert a new document in db
// creating a cart for new users
func (c *CartStore) CreateCart(ctx context.Context, cart *model.Cart) error {
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()
	_, err := c.Collection.InsertOne(ctx, cart)
	return err

}

// search cart using userid , update items of the cart
func (c *CartStore) AddToCart(ctx context.Context, user_id primitive.ObjectID, item model.CartItem) error {
	currentTime := time.Now()
	filter := bson.M{"user_id": user_id}
	update := bson.M{
		"$push": bson.M{
			"items": item,
		},
		"$set": bson.M{"updatedAt": currentTime},
	}
	_, err := c.Collection.UpdateOne(ctx, filter, update)
	return err
}

// search cart by userid then modify using product id
func (c *CartStore) RemoveFromCart(ctx context.Context, product_id, user_id primitive.ObjectID) error {
	// time
	currentTime := time.Now()

	filter := bson.M{"user_id": user_id}
	update := bson.M{
		"$pull": bson.M{
			"items": bson.M{"product_id": product_id}},
		"$set": bson.M{"updatedAt": currentTime},
	}

	_, err := c.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (c *CartStore) ModifyCart(ctx context.Context, cart *model.Cart) error {
	cart.UpdatedAt = time.Now()

	filter := bson.M{"_id": cart.ID}
	update := bson.M{
		"$set": bson.M{
			"items":      cart.Items,
			"totalPrice": cart.TotalPrice,
			"status":     cart.Status,
			"updatedAt":  cart.UpdatedAt,
		},
	}
	res, err := c.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("failed to modify cart")
	}
	if res.ModifiedCount == 0 {
		return errors.New("no cart found")
	}
	return nil
}

func (c *CartStore) GetCartByUserId(ctx context.Context, userId primitive.ObjectID) (*model.Cart, error) {
	var cart *model.Cart
	err := c.Collection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&cart)
	if err != nil {
		return nil, errors.New("failed to get cart")
	}
	return cart, nil
}

func (c *CartStore) UpdatePrice(ctx context.Context, cart *model.Cart) error {

	filter := bson.M{"user_id": cart.UserID}
	update := bson.M{"$set": bson.M{
		"total_price": cart.TotalPrice,
		"updated_at":  time.Now(),
	}}

	_, err := c.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (c *CartStore) ClearCart(ctx context.Context, user_id primitive.ObjectID) error {
	filter := bson.M{"user_id": user_id}
	update := bson.M{"$set": bson.M{
		"items":      []model.CartItem{},
		"totalprice": 0.0,
	}}

	_, err := c.Collection.UpdateOne(ctx, filter, update)
	return err
}
