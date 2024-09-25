package service

import (
	"context"
	"errors"
	"harsh/internal/data"
	"harsh/internal/model"
	"harsh/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartService struct {
	CartStore *data.CartStore
}

func NewCartService(cartStore *data.CartStore) *CartService {
	return &CartService{
		CartStore: cartStore,
	}
}

func (c *CartService) CreateCart(ctx context.Context, cart *model.Cart) error {
	return c.CartStore.CreateCart(ctx, cart)
}

func (c *CartService) AddToCart(ctx context.Context, item model.CartItem, user_id primitive.ObjectID) error {
	err := c.CartStore.AddToCart(ctx, user_id, item)
	if err != nil {
		return err
	}

	// get the current cart to recalculate total price
	cart, err := c.CartStore.GetCartByUserId(ctx, user_id)
	if err != nil {
		return errors.New("failed to get cart")
	}

	// calculate price
	totalPrice := utils.CalculateTotalPrice(cart.Items)

	cart.TotalPrice = totalPrice

	// update cart price
	err = c.CartStore.UpdatePrice(ctx, cart)
	if err != nil {
		return err
	}
	return nil
}

func (c *CartService) RemoveFromCart(ctx context.Context, user_id, product_id primitive.ObjectID) error {
	err := c.CartStore.RemoveFromCart(ctx, product_id, user_id)
	if err != nil {
		return errors.New("failed to remove from cart")
	}

	// get the current cart to recalculate total price
	cart, err := c.CartStore.GetCartByUserId(ctx, user_id)
	if err != nil {
		return errors.New("failed to get cart")
	}

	// calculate price
	totalPrice := utils.CalculateTotalPrice(cart.Items)

	cart.TotalPrice = totalPrice

	// update cart price
	err = c.CartStore.UpdatePrice(ctx, cart)
	if err != nil {
		return err
	}
	return nil
}

func (c *CartService) GetCart(ctx context.Context, user_id primitive.ObjectID) (*model.Cart, error) {
	cart, err := c.CartStore.GetCartByUserId(ctx, user_id)
	if err != nil {
		return nil, errors.New("failed to get cart")
	}
	return cart, nil
}
func (c *CartService) ClearCart(ctx context.Context, user_id primitive.ObjectID) error {
	err := c.CartStore.ClearCart(ctx, user_id)
	if err != nil {
		return errors.New("failed to clear cart")
	}
	return nil
}
