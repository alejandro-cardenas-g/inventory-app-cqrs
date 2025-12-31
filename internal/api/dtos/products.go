package dtos

import "time"

type CreateProductDTO struct {
	SKU         string                 `json:"sku" validate:"required,alphanum,min=3,max=64"`
	Name        string                 `json:"name" validate:"required,min=3,max=150"`
	Description string                 `json:"description" validate:"omitempty,max=2000"`

	CategoryID  int64                  `json:"category_id" validate:"required,gt=0"`
	Brand       string                 `json:"brand" validate:"omitempty,max=100"`

	Price       int64               	`json:"price" validate:"required,gt=0"`
	Currency    string                 `json:"currency" validate:"required,len=3,uppercase"`

	Stock       int                    `json:"stock" validate:"gte=0"`

	Attributes  map[string]any         `json:"attributes" validate:"omitempty,dive,required"`
}

type CreateProductResultDTO struct {
	ID int64 `json:"id"`
}

type GetProductByIDResultDTO struct {
	ID int64 `json:"id"`
	SKU string `json:"sku"`
	Name string `json:"name"`
	Description string `json:"description"`
	CategoryID int64 `json:"category_id"`
	Brand string `json:"brand"`
	Price int64 `json:"price"`
	Currency string `json:"currency"`
	Stock int `json:"stock"`
	Attributes map[string]any `json:"attributes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}