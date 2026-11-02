CREATE TABLE IF NOT EXISTS hub_assortment (
	id UUID PRIMARY KEY,
	hub_id UUID NOT NULL REFERENCES hubs(id) ON DELETE CASCADE,
	product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
	created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
	UNIQUE (hub_id, product_id)
);
CREATE INDEX IF NOT EXISTS idx_hub_assortment_hub ON hub_assortment(hub_id);
CREATE INDEX IF NOT EXISTS idx_hub_assortment_product ON hub_assortment(product_id);
