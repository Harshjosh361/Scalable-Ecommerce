package model

type CartItem struct {
	ProductID  string  `json:"productId" bson:"product_id"`
	UnitPrice  float64 `json:"unitPrice" bson:"unit_price"`
	Quantity   int     `json:"quantity" bson:"quantity"`
	TotalPrice float64 `json:"totalPrice" bson:"total_price"`
}
