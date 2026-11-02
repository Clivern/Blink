CREATE TABLE IF NOT EXISTS voucher_redemptions (
	id UUID PRIMARY KEY,
	promotion_id UUID NOT NULL REFERENCES promotions(id) ON DELETE CASCADE,
	order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE UNIQUE,
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	amount_cents INT NOT NULL,
	created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);
CREATE INDEX IF NOT EXISTS idx_voucher_redemptions_promotion ON voucher_redemptions(promotion_id);
