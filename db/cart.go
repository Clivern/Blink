// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"time"

	"github.com/samber/lo"
)

// Cart is the open basket of a user. One cart per user.
type Cart struct {
	Id        Id
	UserId    Id
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CartItem is a product line in a cart, pinned to the hub that serves it.
type CartItem struct {
	Id        Id
	CartId    Id
	ProductId Id
	HubId     Id
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CartRepository is the interface for cart CRUD.
type CartRepository interface {
	Create(cart *Cart) error
	GetById(id Id) (*Cart, error)
	GetByUser(userId Id) (*Cart, error)
	Delete(id Id) error
}

// CartItemRepository is the interface for cart item CRUD.
type CartItemRepository interface {
	Add(item *CartItem) error
	SetQuantity(cartId, productId Id, quantity int) error
	Remove(cartId, productId Id) error
	ListByCart(cartId Id) ([]*CartItem, error)
	ClearCart(cartId Id) error
}

// CartRepositoryPostgres implements CartRepository for PostgreSQL
type CartRepositoryPostgres struct {
	db *sql.DB
}

// CartItemRepositoryPostgres implements CartItemRepository for PostgreSQL
type CartItemRepositoryPostgres struct {
	db *sql.DB
}

// NewCartRepository returns the repository for the current driver
func NewCartRepository(db *sql.DB) CartRepository {
	return &CartRepositoryPostgres{db: db}
}

// NewCartItemRepository returns the repository for the current driver
func NewCartItemRepository(db *sql.DB) CartItemRepository {
	return &CartItemRepositoryPostgres{db: db}
}

// scanCart maps a row into a cart. Column order must match the select list.
func scanCart(cart *Cart) []any {
	return []any{
		&cart.Id,
		&cart.UserId,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	}
}

// scanCartItem maps a row into a cart item. Column order must match the select list.
func scanCartItem(item *CartItem) []any {
	return []any{
		&item.Id,
		&item.CartId,
		&item.ProductId,
		&item.HubId,
		&item.Quantity,
		&item.CreatedAt,
		&item.UpdatedAt,
	}
}

const cartItemColumns = `id, cart_id, product_id, hub_id, quantity, created_at, updated_at`

// --- Postgres ---

// Create inserts a row
func (r *CartRepositoryPostgres) Create(cart *Cart) error {
	if lo.IsEmpty(cart.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		cart.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO carts (id, user_id) VALUES ($1, $2)`,
		cart.Id.String(),
		cart.UserId.String(),
	)

	return err
}

// GetById returns a cart by Id
func (r *CartRepositoryPostgres) GetById(id Id) (*Cart, error) {
	return QueryOne(
		r.db,
		`SELECT id, user_id, created_at, updated_at FROM carts WHERE id = $1`,
		scanCart,
		id.String(),
	)
}

// GetByUser returns the cart of a user
func (r *CartRepositoryPostgres) GetByUser(userId Id) (*Cart, error) {
	return QueryOne(
		r.db,
		`SELECT id, user_id, created_at, updated_at FROM carts WHERE user_id = $1`,
		scanCart,
		userId.String(),
	)
}

// Delete removes a row
func (r *CartRepositoryPostgres) Delete(id Id) error {
	_, err := r.db.Exec(`DELETE FROM carts WHERE id = $1`, id.String())

	return err
}

// Add inserts a line, raising the quantity when it already exists
func (r *CartItemRepositoryPostgres) Add(item *CartItem) error {
	if lo.IsEmpty(item.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		item.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO cart_items (id, cart_id, product_id, hub_id, quantity)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (cart_id, product_id) DO UPDATE SET quantity = cart_items.quantity + $5,
			updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')`,
		item.Id.String(),
		item.CartId.String(),
		item.ProductId.String(),
		item.HubId.String(),
		item.Quantity,
	)

	return err
}

// SetQuantity overwrites the quantity of a line
func (r *CartItemRepositoryPostgres) SetQuantity(cartId, productId Id, quantity int) error {
	_, err := r.db.Exec(
		`UPDATE cart_items SET quantity = $3, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE cart_id = $1 AND product_id = $2`,
		cartId.String(),
		productId.String(),
		quantity,
	)

	return err
}

// Remove deletes a line from a cart
func (r *CartItemRepositoryPostgres) Remove(cartId, productId Id) error {
	_, err := r.db.Exec(
		`DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2`,
		cartId.String(),
		productId.String(),
	)

	return err
}

// ListByCart returns the lines of a cart
func (r *CartItemRepositoryPostgres) ListByCart(cartId Id) ([]*CartItem, error) {
	return QueryMany(
		r.db,
		`SELECT `+cartItemColumns+` FROM cart_items WHERE cart_id = $1 ORDER BY created_at`,
		scanCartItem,
		cartId.String(),
	)
}

// ClearCart removes every line of a cart
func (r *CartItemRepositoryPostgres) ClearCart(cartId Id) error {
	_, err := r.db.Exec(`DELETE FROM cart_items WHERE cart_id = $1`, cartId.String())

	return err
}
