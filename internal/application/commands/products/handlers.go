package products

import (
	"context"
	"inventory_cqrs/internal/store/persistence"
)

type CreateProductHandler struct {
	productsRepository *persistence.ProductRepository
}

func NewCreateProductHandler(productsRepository *persistence.ProductRepository) *CreateProductHandler {
	return &CreateProductHandler{productsRepository: productsRepository}
}

func (h *CreateProductHandler) Handler(ctx context.Context, command CreateProductCommand) (*CreateProductResult, error) {

	product, err := command.ToProduct()

	if err != nil {
		return nil, err
	}

	result, err := h.productsRepository.Create(ctx, product)

	if err != nil {
		return nil, err
	}

	return &CreateProductResult{
		ID: result.ID,
	}, nil
}