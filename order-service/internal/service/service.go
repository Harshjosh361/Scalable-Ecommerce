package service

import (
	"context"
	"harshy/internal/data"
	"harshy/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	OrderData *data.OrderData
}

func NewOrderService(orderData *data.OrderData) *OrderService {
	return &OrderService{
		OrderData: orderData,
	}
}

func (os *OrderService) CreateOrder(ctx context.Context, order *models.Order) error {
	return os.OrderData.CreateOrder(ctx, order)
}

func (os *OrderService) GetOrderById(ctx context.Context, id primitive.ObjectID) (*models.Order, error) {
	return os.OrderData.GetOrderById(ctx, id)
}

func (os *OrderService) GetOrderByUserId(ctx context.Context, userid primitive.ObjectID) ([]*models.Order, error) {
	return os.OrderData.GetOrderByUserId(ctx, userid)
}
