package products

import "errors"

var (
	ErrCreatingProduct   = errors.New("an error occurred while creating the product")
	ErrInvalidPrice      = errors.New("price must be greater than zero")
	ErrInvalidStock      = errors.New("stock cannot be negative")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrProductNotFound   = errors.New("product not found")
	ErrSKUAlreadyExists  = errors.New("sku already exists")
)
