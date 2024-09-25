package utils

import "harsh/internal/model"

func CalculateTotalPrice(items []model.CartItem) float64 {
	totalPrice := 0.0
	for _, item := range items {
		totalPrice += item.UnitPrice * float64(item.Quantity)
	}
	return totalPrice
}
