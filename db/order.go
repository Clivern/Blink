// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"time"

	"github.com/samber/lo"
)

// Order statuses
const (
	OrderStatusReceived    = "received"
	OrderStatusPicking     = "picking"
	OrderStatusPacked      = "packed"
	OrderStatusHandedRider = "handed_to_rider"
	OrderStatusDelivered   = "delivered"
	OrderStatusCancelled   = "cancelled"
)

// Order item substitution outcomes
const (
	SubstitutionNone        = "none"
	SubstitutionSubstituted = "substituted"
	SubstitutionRefunded    = "refunded"
)

// Pick task statuses
const (
	PickTaskStatusPending  = "pending"
	PickTaskStatusAssigned = "assigned"
	PickTaskStatusPicking  = "picking"
	PickTaskStatusDone     = "done"
)

// Order is a placed basket. Money is stored in cents and prices
// are snapshotted onto the items at order time.
type Order struct {
	Id               Id
	UserId           Id
	HubId            Id
	AddressId        *string
	Status           string
	SubtotalCents    int
	DeliveryFeeCents int
	DiscountCents    int
	TotalCents       int
	VATCents         int
	PromisedBy       *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// OrderItem is a priced line of an order with its pick result.
type OrderItem struct {
	Id             Id
	OrderId        Id
	ProductId      Id
	Quantity       int
	PriceCents     int
	PickedQuantity int
	Substitution   string
	CreatedAt      time.Time
}

// PickTask is the in-hub picking work for one order.
type PickTask struct {
	Id           Id
	OrderId      Id
	PickerUserId *string
	Status       string
	StartedAt    *time.Time
	FinishedAt   *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// OrderRepository is the interface for order CRUD.
type OrderRepository interface {
	Create(order *Order) error
	GetById(id Id) (*Order, error)
	ListByUser(userId Id, limit, offset int) ([]*Order, error)
	ListByHub(hubId Id, status string, limit, offset int) ([]*Order, error)
	UpdateStatus(id Id, status string) error
}

// OrderItemRepository is the interface for order item CRUD.
type OrderItemRepository interface {
	Create(item *OrderItem) error
	ListByOrder(orderId Id) ([]*OrderItem, error)
	UpdatePicked(id Id, pickedQuantity int, substitution string) error
}

// PickTaskRepository is the interface for pick task CRUD.
type PickTaskRepository interface {
	Create(task *PickTask) error
	GetById(id Id) (*PickTask, error)
	GetByOrder(orderId Id) (*PickTask, error)
	Assign(id Id, pickerUserId Id) error
	Complete(id Id) error
}

// OrderRepositoryPostgres implements OrderRepository for PostgreSQL
type OrderRepositoryPostgres struct {
	db *sql.DB
}

// OrderItemRepositoryPostgres implements OrderItemRepository for PostgreSQL
type OrderItemRepositoryPostgres struct {
	db *sql.DB
}

// PickTaskRepositoryPostgres implements PickTaskRepository for PostgreSQL
type PickTaskRepositoryPostgres struct {
	db *sql.DB
}

// NewOrderRepository returns the repository for the current driver
func NewOrderRepository(db *sql.DB) OrderRepository {
	return &OrderRepositoryPostgres{db: db}
}

// NewOrderItemRepository returns the repository for the current driver
func NewOrderItemRepository(db *sql.DB) OrderItemRepository {
	return &OrderItemRepositoryPostgres{db: db}
}

// NewPickTaskRepository returns the repository for the current driver
func NewPickTaskRepository(db *sql.DB) PickTaskRepository {
	return &PickTaskRepositoryPostgres{db: db}
}

// scanOrder maps a row into an order. Column order must match the select list.
func scanOrder(order *Order) []any {
	return []any{
		&order.Id,
		&order.UserId,
		&order.HubId,
		&order.AddressId,
		&order.Status,
		&order.SubtotalCents,
		&order.DeliveryFeeCents,
		&order.DiscountCents,
		&order.TotalCents,
		&order.VATCents,
		&order.PromisedBy,
		&order.CreatedAt,
		&order.UpdatedAt,
	}
}

// scanOrderItem maps a row into an order item.
func scanOrderItem(item *OrderItem) []any {
	return []any{
		&item.Id,
		&item.OrderId,
		&item.ProductId,
		&item.Quantity,
		&item.PriceCents,
		&item.PickedQuantity,
		&item.Substitution,
		&item.CreatedAt,
	}
}

// scanPickTask maps a row into a pick task. Column order must match the select list.
func scanPickTask(task *PickTask) []any {
	return []any{
		&task.Id,
		&task.OrderId,
		&task.PickerUserId,
		&task.Status,
		&task.StartedAt,
		&task.FinishedAt,
		&task.CreatedAt,
		&task.UpdatedAt,
	}
}

const orderColumns = `id, user_id, hub_id, address_id, status, subtotal_cents, delivery_fee_cents, discount_cents, total_cents, vat_cents, promised_by, created_at, updated_at`

// --- Postgres ---

// Create inserts a row
func (r *OrderRepositoryPostgres) Create(order *Order) error {
	if lo.IsEmpty(order.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		order.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO orders (id, user_id, hub_id, address_id, status, subtotal_cents,
			delivery_fee_cents, discount_cents, total_cents, vat_cents, promised_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.Id.String(),
		order.UserId.String(),
		order.HubId.String(),
		order.AddressId,
		order.Status,
		order.SubtotalCents,
		order.DeliveryFeeCents,
		order.DiscountCents,
		order.TotalCents,
		order.VATCents,
		order.PromisedBy,
	)

	return err
}

// GetById returns an order by Id
func (r *OrderRepositoryPostgres) GetById(id Id) (*Order, error) {
	return QueryOne(
		r.db,
		`SELECT `+orderColumns+` FROM orders WHERE id = $1`,
		scanOrder,
		id.String(),
	)
}

// ListByUser returns the orders of a user, newest first
func (r *OrderRepositoryPostgres) ListByUser(userId Id, limit, offset int) ([]*Order, error) {
	return QueryMany(
		r.db,
		`SELECT `+orderColumns+` FROM orders WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
		scanOrder,
		userId.String(),
		limit,
		offset,
	)
}

// ListByHub returns the orders of a hub in a status, oldest first
func (r *OrderRepositoryPostgres) ListByHub(hubId Id, status string, limit, offset int) ([]*Order, error) {
	return QueryMany(
		r.db,
		`SELECT `+orderColumns+` FROM orders WHERE hub_id = $1 AND status = $2 ORDER BY created_at LIMIT $3 OFFSET $4`,
		scanOrder,
		hubId.String(),
		status,
		limit,
		offset,
	)
}

// UpdateStatus moves an order to a new status
func (r *OrderRepositoryPostgres) UpdateStatus(id Id, status string) error {
	_, err := r.db.Exec(
		`UPDATE orders SET status = $2, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC') WHERE id = $1`,
		id.String(),
		status,
	)

	return err
}

// Create inserts a row
func (r *OrderItemRepositoryPostgres) Create(item *OrderItem) error {
	if lo.IsEmpty(item.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		item.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO order_items (id, order_id, product_id, quantity, price_cents, picked_quantity, substitution)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		item.Id.String(),
		item.OrderId.String(),
		item.ProductId.String(),
		item.Quantity,
		item.PriceCents,
		item.PickedQuantity,
		item.Substitution,
	)

	return err
}

// ListByOrder returns the items of an order
func (r *OrderItemRepositoryPostgres) ListByOrder(orderId Id) ([]*OrderItem, error) {
	return QueryMany(
		r.db,
		`SELECT id, order_id, product_id, quantity, price_cents, picked_quantity, substitution, created_at
		FROM order_items WHERE order_id = $1`,
		scanOrderItem,
		orderId.String(),
	)
}

// UpdatePicked records the pick result of a line
func (r *OrderItemRepositoryPostgres) UpdatePicked(id Id, pickedQuantity int, substitution string) error {
	_, err := r.db.Exec(
		`UPDATE order_items SET picked_quantity = $2, substitution = $3 WHERE id = $1`,
		id.String(),
		pickedQuantity,
		substitution,
	)

	return err
}

// Create inserts a row
func (r *PickTaskRepositoryPostgres) Create(task *PickTask) error {
	if lo.IsEmpty(task.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		task.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO pick_tasks (id, order_id, picker_user_id, status, started_at, finished_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		task.Id.String(),
		task.OrderId.String(),
		task.PickerUserId,
		task.Status,
		task.StartedAt,
		task.FinishedAt,
	)

	return err
}

// GetById returns a pick task by Id
func (r *PickTaskRepositoryPostgres) GetById(id Id) (*PickTask, error) {
	return QueryOne(
		r.db,
		`SELECT id, order_id, picker_user_id, status, started_at, finished_at, created_at, updated_at
		FROM pick_tasks WHERE id = $1`,
		scanPickTask,
		id.String(),
	)
}

// GetByOrder returns the pick task of an order
func (r *PickTaskRepositoryPostgres) GetByOrder(orderId Id) (*PickTask, error) {
	return QueryOne(
		r.db,
		`SELECT id, order_id, picker_user_id, status, started_at, finished_at, created_at, updated_at
		FROM pick_tasks WHERE order_id = $1`,
		scanPickTask,
		orderId.String(),
	)
}

// Assign hands a task to a picker and starts it
func (r *PickTaskRepositoryPostgres) Assign(id Id, pickerUserId Id) error {
	_, err := r.db.Exec(
		`UPDATE pick_tasks
		SET picker_user_id = $2, status = $3, started_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
			updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE id = $1`,
		id.String(),
		pickerUserId.String(),
		PickTaskStatusPicking,
	)

	return err
}

// Complete closes a task
func (r *PickTaskRepositoryPostgres) Complete(id Id) error {
	_, err := r.db.Exec(
		`UPDATE pick_tasks
		SET status = $2, finished_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
			updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE id = $1`,
		id.String(),
		PickTaskStatusDone,
	)

	return err
}
