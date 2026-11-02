CREATE TABLE IF NOT EXISTS product_categories (
	id UUID PRIMARY KEY,
	product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
	category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
	created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
	UNIQUE (product_id, category_id)
);
CREATE INDEX IF NOT EXISTS idx_product_categories_product ON product_categories(product_id);
CREATE INDEX IF NOT EXISTS idx_product_categories_category ON product_categories(category_id);
