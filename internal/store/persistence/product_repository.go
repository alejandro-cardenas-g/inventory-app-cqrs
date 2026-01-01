package persistence

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"inventory_cqrs/internal/domain/products"
	db "inventory_cqrs/internal/store/persistence/sqlc"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductRepository struct {
	queries *db.Queries
}

func NewProductRepository(tx DBTX) *ProductRepository {
	return &ProductRepository{queries: db.New(tx)}
}

func (r *ProductRepository) UseTX(tx DBTX) *ProductRepository{
	return &ProductRepository{queries: db.New(tx)}
}

func (r *ProductRepository) GetByID(ctx context.Context, id int64) (*products.Product, error) {
	product, err := r.queries.GetProductByID(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, products.ErrProductNotFound
		}
		return nil, err
	}


	var attributes map[string]any
	if err := json.Unmarshal(product.Attributes, &attributes); err != nil {
		return nil, ErrorPerformingOperation
	}

	p, err := products.New(
		product.Sku, 
		product.Name, 
		product.CategoryID,
		product.PriceAmount,
		product.PriceCurrency, 
		product.Stock,
		attributes,
		product.Description.String,
		product.Brand,
	)

	if err != nil {
		return nil, err
	}
	
	p.SetID(product.ID)
	p.SetAuditory(product.CreatedAt.Time, product.UpdatedAt.Time)
	return p, nil
}

type createResult struct {
	ID int64
}

func (r *ProductRepository) Create(ctx context.Context, product *products.Product) (*createResult, error) {

	attributes, err := json.Marshal(product.GetAttributes())
	if err != nil {
		return nil, ErrorPerformingOperation
	}

	params := db.CreateProductParams{
		Sku: product.GetSKU(),
		Name: product.GetName(),
		Description: pgtype.Text{String: product.GetDescription(), Valid: true},
		PriceAmount: product.GetPriceCents(),
		PriceCurrency: product.GetCurrency(),
		Stock: int32(product.GetStock()),
		IsActive: product.IsActive(),
		Attributes: attributes,
		CategoryID: product.GetCategoryID(),
		Brand: product.GetBrand(),
	}

	id, err := r.queries.CreateProduct(ctx, params)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case string(ErrorCodeUniqueViolation):
				return nil, products.ErrSKUAlreadyExists
			}
		}
		return nil, ErrorPerformingOperation
	}

	return &createResult{
		ID: id,
	}, nil
}

