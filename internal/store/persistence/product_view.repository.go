package persistence

import (
	"context"
	"encoding/json"
	"inventory_cqrs/internal/store/persistence/readmodel"
	db "inventory_cqrs/internal/store/persistence/sqlc"
)

type ProductViewRepository struct {
	queries *db.Queries
}

func NewProductViewRepository(tx DBTX) *ProductViewRepository {
	return &ProductViewRepository{queries: db.New(tx)}
}

func (r *ProductViewRepository) GetByID(ctx context.Context, id int64) (*readmodel.ProductView, error) {
	product, err := r.queries.GetProductForReadModel(ctx, id)

	if err != nil {
		return nil, err
	}

	var attributes map[string]any
	if err := json.Unmarshal(product.Attributes, &attributes); err != nil {
		return nil, err
	}

	p := &readmodel.ProductView{
		ID:          product.ID,
		SKU:         product.Sku,
		Name:        product.Name,
		Description: product.Description.String,
		Category:    product.CategoryName.String,
		CategoryID:  product.CategoryID,
		PriceCents:  product.PriceAmount,
		Currency:    product.PriceCurrency,
		Stock:       int32(product.Stock),
		Attributes:  attributes,
	}

	return p, nil
}
