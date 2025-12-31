package products

import "inventory_cqrs/internal/domain/products"

type GetProductByIDResult struct {
	Product *products.Product
}