package products

import "time"

type Product struct {
	id          int64
	sku         string
	name        string
	description string

	categoryID int64
	brand      string

	priceCents int64
	currency   string
	stock      int

	isActive   bool
	attributes map[string]any

	createdAt time.Time
	updatedAt time.Time
}


func New(
	sku string,
	name string,
	categoryID int64,
	priceCents int64,
	currency string,
	stock int,
	attributes map[string]any,
) (*Product, error) {

	if priceCents <= 0 {
		return nil, ErrInvalidPrice
	}

	if stock < 0 {
		return nil, ErrInvalidStock
	}

	return &Product{
		sku:         sku,
		name:        name,
		categoryID:  categoryID,
		priceCents: priceCents,
		currency:   currency,
		stock:      stock,
		isActive:   true,
		attributes: attributes,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}, nil
}

func (p *Product) GetID() int64 {
	return p.id
}

func (p *Product) GetSKU() string {
	return p.sku
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetDescription() string {
	return p.description
}

func (p *Product) GetCategoryID() int64 {
	return p.categoryID
}

func (p *Product) GetBrand() string {
	return p.brand
}

func (p *Product) GetPriceCents() int64 {
	return p.priceCents
}

func (p *Product) GetCurrency() string {
	return p.currency
}

func (p *Product) GetStock() int {
	return p.stock
}

func (p *Product) GetAttributes() map[string]any {
	return p.attributes
}

func (p *Product) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *Product) GetUpdatedAt() time.Time {
	return p.updatedAt
}