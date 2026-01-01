package products

import (
	"context"
	"errors"
	"inventory_cqrs/internal/store/persistence"
)

var ErrProductNotFound = errors.New("product not found")

type GetProductByIDHandler struct {
	productsRepository *persistence.ProductRepository
}

func NewGetProductByIDHandler(productsRepository *persistence.ProductRepository) *GetProductByIDHandler {
	return &GetProductByIDHandler{productsRepository: productsRepository}
}

func (h *GetProductByIDHandler) Handler(ctx context.Context, query GetProductByIDQuery) (*GetProductByIDResult, error) {
	product, err := h.productsRepository.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, ErrProductNotFound
	}

	return &GetProductByIDResult{Product: product}, nil
}