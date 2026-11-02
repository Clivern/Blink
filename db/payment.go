// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"time"

	"github.com/samber/lo"
)

// Payment statuses
const (
	PaymentStatusPending  = "pending"
	PaymentStatusPaid     = "paid"
	PaymentStatusFailed   = "failed"
	PaymentStatusRefunded = "refunded"
)

// Promotion kinds
const (
	PromotionKindPercent = "percent"
	PromotionKindAmount  = "amount"
)

// Payment is the charge for one order.
type Payment struct {
	Id          Id
	OrderId     Id
	Provider    string
	AmountCents int
	Status      string
	PaidAt      *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Promotion is a voucher campaign. Kind is percent or amount,
// Value the off in percent points or cents.
type Promotion struct {
	Id             Id
	Code           string
	Kind           string
	Value          int
	MinBasketCents int
	StartsAt       *time.Time
	EndsAt         *time.Time
	MaxRedemptions *int
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// VoucherRedemption records one voucher use against an order.
type VoucherRedemption struct {
	Id          Id
	PromotionId Id
	OrderId     Id
	UserId      Id
	AmountCents int
	CreatedAt   time.Time
}

// PaymentRepository is the interface for payment CRUD.
type PaymentRepository interface {
	Create(payment *Payment) error
	GetById(id Id) (*Payment, error)
	GetByOrder(orderId Id) (*Payment, error)
	UpdateStatus(id Id, status string, paidAt *time.Time) error
}

// PromotionRepository is the interface for promotion CRUD.
type PromotionRepository interface {
	Create(promotion *Promotion) error
	GetById(id Id) (*Promotion, error)
	GetByCode(code string) (*Promotion, error)
	ListActive() ([]*Promotion, error)
}

// VoucherRedemptionRepository is the interface for voucher redemption CRUD.
type VoucherRedemptionRepository interface {
	Create(redemption *VoucherRedemption) error
	GetByOrder(orderId Id) (*VoucherRedemption, error)
	ListByPromotion(promotionId Id) ([]*VoucherRedemption, error)
}

// PaymentRepositoryPostgres implements PaymentRepository for PostgreSQL
type PaymentRepositoryPostgres struct {
	db *sql.DB
}

// PromotionRepositoryPostgres implements PromotionRepository for PostgreSQL
type PromotionRepositoryPostgres struct {
	db *sql.DB
}

// VoucherRedemptionRepositoryPostgres implements VoucherRedemptionRepository for PostgreSQL
type VoucherRedemptionRepositoryPostgres struct {
	db *sql.DB
}

// NewPaymentRepository returns the repository for the current driver
func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &PaymentRepositoryPostgres{db: db}
}

// NewPromotionRepository returns the repository for the current driver
func NewPromotionRepository(db *sql.DB) PromotionRepository {
	return &PromotionRepositoryPostgres{db: db}
}

// NewVoucherRedemptionRepository returns the repository for the current driver
func NewVoucherRedemptionRepository(db *sql.DB) VoucherRedemptionRepository {
	return &VoucherRedemptionRepositoryPostgres{db: db}
}

// scanPayment maps a row into a payment. Column order must match the select list.
func scanPayment(payment *Payment) []any {
	return []any{
		&payment.Id,
		&payment.OrderId,
		&payment.Provider,
		&payment.AmountCents,
		&payment.Status,
		&payment.PaidAt,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	}
}

// scanPromotion maps a row into a promotion. Column order must match the select list.
func scanPromotion(promotion *Promotion) []any {
	return []any{
		&promotion.Id,
		&promotion.Code,
		&promotion.Kind,
		&promotion.Value,
		&promotion.MinBasketCents,
		&promotion.StartsAt,
		&promotion.EndsAt,
		&promotion.MaxRedemptions,
		&promotion.IsActive,
		&promotion.CreatedAt,
		&promotion.UpdatedAt,
	}
}

// scanVoucherRedemption maps a row into a voucher redemption.
func scanVoucherRedemption(redemption *VoucherRedemption) []any {
	return []any{
		&redemption.Id,
		&redemption.PromotionId,
		&redemption.OrderId,
		&redemption.UserId,
		&redemption.AmountCents,
		&redemption.CreatedAt,
	}
}

const paymentColumns = `id, order_id, provider, amount_cents, status, paid_at, created_at, updated_at`

const promotionColumns = `id, code, kind, value, min_basket_cents, starts_at, ends_at, max_redemptions, is_active, created_at, updated_at`

// --- Postgres ---

// Create inserts a row
func (r *PaymentRepositoryPostgres) Create(payment *Payment) error {
	if lo.IsEmpty(payment.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		payment.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO payments (id, order_id, provider, amount_cents, status, paid_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		payment.Id.String(),
		payment.OrderId.String(),
		payment.Provider,
		payment.AmountCents,
		payment.Status,
		payment.PaidAt,
	)

	return err
}

// GetById returns a payment by Id
func (r *PaymentRepositoryPostgres) GetById(id Id) (*Payment, error) {
	return QueryOne(
		r.db,
		`SELECT `+paymentColumns+` FROM payments WHERE id = $1`,
		scanPayment,
		id.String(),
	)
}

// GetByOrder returns the payment of an order
func (r *PaymentRepositoryPostgres) GetByOrder(orderId Id) (*Payment, error) {
	return QueryOne(
		r.db,
		`SELECT `+paymentColumns+` FROM payments WHERE order_id = $1`,
		scanPayment,
		orderId.String(),
	)
}

// UpdateStatus moves a payment to a new status
func (r *PaymentRepositoryPostgres) UpdateStatus(id Id, status string, paidAt *time.Time) error {
	_, err := r.db.Exec(
		`UPDATE payments SET status = $2, paid_at = $3, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC') WHERE id = $1`,
		id.String(),
		status,
		paidAt,
	)

	return err
}

// Create inserts a row
func (r *PromotionRepositoryPostgres) Create(promotion *Promotion) error {
	if lo.IsEmpty(promotion.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		promotion.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO promotions (id, code, kind, value, min_basket_cents, starts_at, ends_at, max_redemptions, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		promotion.Id.String(),
		promotion.Code,
		promotion.Kind,
		promotion.Value,
		promotion.MinBasketCents,
		promotion.StartsAt,
		promotion.EndsAt,
		promotion.MaxRedemptions,
		promotion.IsActive,
	)

	return err
}

// GetById returns a promotion by Id
func (r *PromotionRepositoryPostgres) GetById(id Id) (*Promotion, error) {
	return QueryOne(
		r.db,
		`SELECT `+promotionColumns+` FROM promotions WHERE id = $1`,
		scanPromotion,
		id.String(),
	)
}

// GetByCode returns a promotion by code
func (r *PromotionRepositoryPostgres) GetByCode(code string) (*Promotion, error) {
	return QueryOne(
		r.db,
		`SELECT `+promotionColumns+` FROM promotions WHERE code = $1`,
		scanPromotion,
		code,
	)
}

// ListActive returns active promotions by code
func (r *PromotionRepositoryPostgres) ListActive() ([]*Promotion, error) {
	return QueryMany(
		r.db,
		`SELECT `+promotionColumns+` FROM promotions WHERE is_active = true ORDER BY code`,
		scanPromotion,
	)
}

// Create inserts a row
func (r *VoucherRedemptionRepositoryPostgres) Create(redemption *VoucherRedemption) error {
	if lo.IsEmpty(redemption.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		redemption.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO voucher_redemptions (id, promotion_id, order_id, user_id, amount_cents)
		VALUES ($1, $2, $3, $4, $5)`,
		redemption.Id.String(),
		redemption.PromotionId.String(),
		redemption.OrderId.String(),
		redemption.UserId.String(),
		redemption.AmountCents,
	)

	return err
}

// GetByOrder returns the redemption of an order
func (r *VoucherRedemptionRepositoryPostgres) GetByOrder(orderId Id) (*VoucherRedemption, error) {
	return QueryOne(
		r.db,
		`SELECT id, promotion_id, order_id, user_id, amount_cents, created_at
		FROM voucher_redemptions WHERE order_id = $1`,
		scanVoucherRedemption,
		orderId.String(),
	)
}

// ListByPromotion returns the redemptions of a promotion, newest first
func (r *VoucherRedemptionRepositoryPostgres) ListByPromotion(promotionId Id) ([]*VoucherRedemption, error) {
	return QueryMany(
		r.db,
		`SELECT id, promotion_id, order_id, user_id, amount_cents, created_at
		FROM voucher_redemptions WHERE promotion_id = $1 ORDER BY created_at DESC`,
		scanVoucherRedemption,
		promotionId.String(),
	)
}
