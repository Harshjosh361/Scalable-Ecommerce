package service

import (
	"context"
	"errors"
	"harsh/internal/data"
	model "harsh/internal/models"
)

type ProductService struct {
	productStore *data.ProductStore
}

func NewProductService(productStore *data.ProductStore) *ProductService {
	return &ProductService{
		productStore: productStore,
	}
}

func (s *ProductService) GetAllProduct(ctx context.Context) ([]*model.Product, error) {
	products, err := s.productStore.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	return s.productStore.GetProductById(ctx, id)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	return s.productStore.DeleteProduct(ctx, id)
}

func (s *ProductService) CreateProduct(ctx context.Context, product *model.Product) error {
	// check if product is already created
	_, err := s.productStore.FindProductByName(ctx, product.Name)
	if err == nil {
		return errors.New("product already exists")
	}

	return s.productStore.CreateProduct(ctx, product)
}

func (s *ProductService) ModifyProduct(ctx context.Context, newProduct *model.Product) error {
	return s.productStore.ModifyProduct(ctx, newProduct)
}
