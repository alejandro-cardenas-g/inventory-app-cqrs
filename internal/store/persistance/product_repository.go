package persistance

import (
	"context"
	"inventory_cqrs/internal/domain/products"
)

type ProductRepository struct {
	db DBTX
}

func (r *ProductRepository) UseTX(tx DBTX) *ProductRepository{
	return &ProductRepository{db: tx}
}

func (r *ProductRepository) GetByID(ctx context.Context, id int64) (*products.Product, error) {
	return nil, nil
}

type createResult struct {
	ID int64
}

func (r *ProductRepository) Create(ctx context.Context, product *products.Product) (*createResult, error) {
	return nil, nil
}

