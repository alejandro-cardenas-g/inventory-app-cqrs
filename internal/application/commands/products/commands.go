package products

import "inventory_cqrs/internal/domain/products"

type CreateProductCommand struct {
	SKU         string
	Name        string
	Description string
	CategoryID  int64
	Brand       string
	Price       int64
	Currency    string
	Stock       int32
	Attributes  map[string]any
}


type GetProductByIDCommand struct {
	ID int64
}

func (c *CreateProductCommand) ToProduct() (*products.Product, error) {
	return products.New(c.SKU, c.Name, c.CategoryID, c.Price, c.Currency, c.Stock, c.Attributes, c.Description, c.Brand)
}