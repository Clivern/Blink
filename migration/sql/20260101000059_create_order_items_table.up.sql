CREATE TABLE IF NOT EXISTS order_items (
	id UUID PRIMARY KEY,
	order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
	product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
	quantity INT NOT NULL,
	price_cents INT NOT NULL,
	picked_quantity INT NOT NULL DEFAULT 0,
	substitution VARCHAR(20) NOT NULL DEFAULT 'none',
	created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);
CREATE INDEX IF NOT EXISTS idx_order_items_order ON order_items(order_id);
