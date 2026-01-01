CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,

    sku VARCHAR(64) NOT NULL UNIQUE,

    name VARCHAR(100) NOT NULL,
    description TEXT,

    category_id BIGINT NOT NULL,
    brand VARCHAR(100) NOT NULL,

    price_amount BIGINT NOT NULL CHECK (price_amount >= 0),

    price_currency CHAR(3) NOT NULL,

    stock INTEGER NOT NULL CHECK (stock >= 0),

    is_active BOOLEAN NOT NULL DEFAULT true,

    attributes JSONB NOT NULL DEFAULT '{}'::jsonb,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE INDEX idx_products_name_fts
ON products
USING gin (to_tsvector('simple', name));

CREATE INDEX idx_products_is_active
ON products (is_active);

CREATE INDEX idx_products_created_at
ON products (created_at DESC);

CREATE UNIQUE INDEX idx_products_sku_unique ON products (sku);

-- create trigger to set updated_at

CREATE OR REPLACE FUNCTION set_products_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_products_updated_at
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION set_products_updated_at();
