package readmodel

type ProductView struct {
	ID          int64
	SKU         string
	Name        string
	Description string
	CategoryID  int64
	Category    string
	PriceCents  int64
	Currency    string
	Stock       int32
	Attributes  map[string]any
}
