package products

import (
	"context"
	"errors"
	"inventory_cqrs/internal/store/persistance"
)

var ErrProductNotFound = errors.New("product not found")

type GetProductByIDHandler struct {
	productsRepository *persistance.ProductRepository
}

func NewGetProductByIDHandler(productsRepository *persistance.ProductRepository) *GetProductByIDHandler {
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