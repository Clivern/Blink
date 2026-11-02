// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"time"

	"github.com/samber/lo"
)

// Delivery statuses
const (
	DeliveryStatusAssigned  = "assigned"
	DeliveryStatusAtDoor    = "at_door"
	DeliveryStatusDelivered = "delivered"
	DeliveryStatusFailed    = "failed"
)

// Delivery is the last-mile ride for one order.
type Delivery struct {
	Id          Id
	OrderId     Id
	RiderUserId *string
	Status      string
	AssignedAt  *time.Time
	DeliveredAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// DeliveryRepository is the interface for delivery CRUD.
type DeliveryRepository interface {
	Create(delivery *Delivery) error
	GetById(id Id) (*Delivery, error)
	GetByOrder(orderId Id) (*Delivery, error)
	ListByRider(riderUserId Id, limit, offset int) ([]*Delivery, error)
	Assign(id Id, riderUserId Id) error
	UpdateStatus(id Id, status string, deliveredAt *time.Time) error
}

// DeliveryRepositoryPostgres implements DeliveryRepository for PostgreSQL
type DeliveryRepositoryPostgres struct {
	db *sql.DB
}

// NewDeliveryRepository returns the repository for the current driver
func NewDeliveryRepository(db *sql.DB) DeliveryRepository {
	return &DeliveryRepositoryPostgres{db: db}
}

// scanDelivery maps a row into a delivery. Column order must match the select list.
func scanDelivery(delivery *Delivery) []any {
	return []any{
		&delivery.Id,
		&delivery.OrderId,
		&delivery.RiderUserId,
		&delivery.Status,
		&delivery.AssignedAt,
		&delivery.DeliveredAt,
		&delivery.CreatedAt,
		&delivery.UpdatedAt,
	}
}

const deliveryColumns = `id, order_id, rider_user_id, status, assigned_at, delivered_at, created_at, updated_at`

// --- Postgres ---

// Create inserts a row
func (r *DeliveryRepositoryPostgres) Create(delivery *Delivery) error {
	if lo.IsEmpty(delivery.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		delivery.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO deliveries (id, order_id, rider_user_id, status, assigned_at, delivered_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		delivery.Id.String(),
		delivery.OrderId.String(),
		delivery.RiderUserId,
		delivery.Status,
		delivery.AssignedAt,
		delivery.DeliveredAt,
	)

	return err
}

// GetById returns a delivery by Id
func (r *DeliveryRepositoryPostgres) GetById(id Id) (*Delivery, error) {
	return QueryOne(
		r.db,
		`SELECT `+deliveryColumns+` FROM deliveries WHERE id = $1`,
		scanDelivery,
		id.String(),
	)
}

// GetByOrder returns the delivery of an order
func (r *DeliveryRepositoryPostgres) GetByOrder(orderId Id) (*Delivery, error) {
	return QueryOne(
		r.db,
		`SELECT `+deliveryColumns+` FROM deliveries WHERE order_id = $1`,
		scanDelivery,
		orderId.String(),
	)
}

// ListByRider returns the deliveries of a rider, newest first
func (r *DeliveryRepositoryPostgres) ListByRider(riderUserId Id, limit, offset int) ([]*Delivery, error) {
	return QueryMany(
		r.db,
		`SELECT `+deliveryColumns+` FROM deliveries WHERE rider_user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
		scanDelivery,
		riderUserId.String(),
		limit,
		offset,
	)
}

// Assign hands a delivery to a rider
func (r *DeliveryRepositoryPostgres) Assign(id Id, riderUserId Id) error {
	_, err := r.db.Exec(
		`UPDATE deliveries
		SET rider_user_id = $2, status = $3, assigned_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
			updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE id = $1`,
		id.String(),
		riderUserId.String(),
		DeliveryStatusAssigned,
	)

	return err
}

// UpdateStatus moves a delivery to a new status
func (r *DeliveryRepositoryPostgres) UpdateStatus(id Id, status string, deliveredAt *time.Time) error {
	_, err := r.db.Exec(
		`UPDATE deliveries SET status = $2, delivered_at = $3, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC') WHERE id = $1`,
		id.String(),
		status,
		deliveredAt,
	)

	return err
}
