CREATE TABLE IF NOT EXISTS stock_movements (
	id UUID PRIMARY KEY,
	hub_id UUID NOT NULL REFERENCES hubs(id) ON DELETE CASCADE,
	product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
	qty_change INT NOT NULL,
	reason VARCHAR(30) NOT NULL,
	ref VARCHAR(120),
	created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);
CREATE INDEX IF NOT EXISTS idx_stock_movements_hub_product ON stock_movements(hub_id, product_id);
