DROP TRIGGER IF EXISTS trg_products_updated_at ON products;
DROP FUNCTION IF EXISTS set_products_updated_at();

DROP INDEX IF EXISTS idx_products_name_fts;
DROP INDEX IF EXISTS idx_products_is_active;
DROP INDEX IF EXISTS idx_products_created_at;
DROP INDEX IF EXISTS idx_products_sku_unique;

DROP TABLE IF EXISTS products;