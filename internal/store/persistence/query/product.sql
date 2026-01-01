-- name: CreateProduct :one
INSERT INTO products (
    sku,
    name,
    description,
    price_amount,
    price_currency,
    stock,
    is_active,
    attributes,
    category_id,
    brand
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING id;

-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1 LIMIT 1;