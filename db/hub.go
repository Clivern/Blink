// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"time"

	"github.com/samber/lo"
)

// Hub statuses
const (
	HubStatusOpen   = "open"
	HubStatusClosed = "closed"
)

// Stock movement reasons
const (
	StockMovementInbound    = "inbound"
	StockMovementSale       = "sale"
	StockMovementWaste      = "waste"
	StockMovementCorrection = "correction"
)

// Hub is a dark store serving the addresses around it.
type Hub struct {
	Id        Id
	Name      string
	City      string
	Street    *string
	Lat       float64
	Lng       float64
	RadiusM   int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// HubAssortment marks a product as carried by a hub.
type HubAssortment struct {
	Id        Id
	HubId     Id
	ProductId Id
	CreatedAt time.Time
}

// HubStock is the live count of a product at a hub. Quantity is
// the physical count, Reserved the units held by open baskets.
type HubStock struct {
	Id           Id
	HubId        Id
	ProductId    Id
	Quantity     int
	Reserved     int
	ReorderPoint int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// StockMovement is an immutable log of every unit in or out.
type StockMovement struct {
	Id        Id
	HubId     Id
	ProductId Id
	QtyChange int
	Reason    string
	Ref       *string
	CreatedAt time.Time
}

// HubRepository is the interface for hub CRUD.
type HubRepository interface {
	Create(hub *Hub) error
	GetById(id Id) (*Hub, error)
	List(limit, offset int) ([]*Hub, error)
	Update(hub *Hub) error
	Delete(id Id) error
}

// HubAssortmentRepository is the interface for hub assortment links.
type HubAssortmentRepository interface {
	Add(hubId, productId Id) error
	Remove(hubId, productId Id) error
	ListByHub(hubId Id) ([]*HubAssortment, error)
}

// HubStockRepository is the interface for per-hub stock levels.
type HubStockRepository interface {
	Create(stock *HubStock) error
	GetByHubProduct(hubId, productId Id) (*HubStock, error)
	ListByHub(hubId Id) ([]*HubStock, error)
	ListLowStock(hubId Id) ([]*HubStock, error)
	Adjust(hubId, productId Id, delta int) error
	Delete(id Id) error
}

// StockMovementRepository is the interface for the stock ledger.
type StockMovementRepository interface {
	Record(movement *StockMovement) error
	ListByHubProduct(hubId, productId Id, limit int) ([]*StockMovement, error)
}

// HubRepositoryPostgres implements HubRepository for PostgreSQL
type HubRepositoryPostgres struct {
	db *sql.DB
}

// HubAssortmentRepositoryPostgres implements HubAssortmentRepository for PostgreSQL
type HubAssortmentRepositoryPostgres struct {
	db *sql.DB
}

// HubStockRepositoryPostgres implements HubStockRepository for PostgreSQL
type HubStockRepositoryPostgres struct {
	db *sql.DB
}

// StockMovementRepositoryPostgres implements StockMovementRepository for PostgreSQL
type StockMovementRepositoryPostgres struct {
	db *sql.DB
}

// NewHubRepository returns the repository for the current driver
func NewHubRepository(db *sql.DB) HubRepository {
	return &HubRepositoryPostgres{db: db}
}

// NewHubAssortmentRepository returns the repository for the current driver
func NewHubAssortmentRepository(db *sql.DB) HubAssortmentRepository {
	return &HubAssortmentRepositoryPostgres{db: db}
}

// NewHubStockRepository returns the repository for the current driver
func NewHubStockRepository(db *sql.DB) HubStockRepository {
	return &HubStockRepositoryPostgres{db: db}
}

// NewStockMovementRepository returns the repository for the current driver
func NewStockMovementRepository(db *sql.DB) StockMovementRepository {
	return &StockMovementRepositoryPostgres{db: db}
}

// scanHub maps a row into a hub. Column order must match the select list.
func scanHub(hub *Hub) []any {
	return []any{
		&hub.Id,
		&hub.Name,
		&hub.City,
		&hub.Street,
		&hub.Lat,
		&hub.Lng,
		&hub.RadiusM,
		&hub.Status,
		&hub.CreatedAt,
		&hub.UpdatedAt,
	}
}

// scanHubAssortment maps a row into a hub assortment link.
func scanHubAssortment(link *HubAssortment) []any {
	return []any{
		&link.Id,
		&link.HubId,
		&link.ProductId,
		&link.CreatedAt,
	}
}

// scanHubStock maps a row into hub stock. Column order must match the select list.
func scanHubStock(stock *HubStock) []any {
	return []any{
		&stock.Id,
		&stock.HubId,
		&stock.ProductId,
		&stock.Quantity,
		&stock.Reserved,
		&stock.ReorderPoint,
		&stock.CreatedAt,
		&stock.UpdatedAt,
	}
}

// scanStockMovement maps a row into a stock movement.
func scanStockMovement(movement *StockMovement) []any {
	return []any{
		&movement.Id,
		&movement.HubId,
		&movement.ProductId,
		&movement.QtyChange,
		&movement.Reason,
		&movement.Ref,
		&movement.CreatedAt,
	}
}

const hubColumns = `id, name, city, street, lat, lng, radius_m, status, created_at, updated_at`

const hubStockColumns = `id, hub_id, product_id, quantity, reserved, reorder_point, created_at, updated_at`

// --- Postgres ---

// Create inserts a row
func (r *HubRepositoryPostgres) Create(hub *Hub) error {
	if lo.IsEmpty(hub.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		hub.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO hubs (id, name, city, street, lat, lng, radius_m, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		hub.Id.String(),
		hub.Name,
		hub.City,
		hub.Street,
		hub.Lat,
		hub.Lng,
		hub.RadiusM,
		hub.Status,
	)

	return err
}

// GetById returns a hub by Id
func (r *HubRepositoryPostgres) GetById(id Id) (*Hub, error) {
	return QueryOne(
		r.db,
		`SELECT `+hubColumns+` FROM hubs WHERE id = $1`,
		scanHub,
		id.String(),
	)
}

// List returns hubs alphabetically
func (r *HubRepositoryPostgres) List(limit, offset int) ([]*Hub, error) {
	return QueryMany(
		r.db,
		`SELECT `+hubColumns+` FROM hubs ORDER BY city, name LIMIT $1 OFFSET $2`,
		scanHub,
		limit,
		offset,
	)
}

// Update rewrites a row
func (r *HubRepositoryPostgres) Update(hub *Hub) error {
	_, err := r.db.Exec(
		`UPDATE hubs
		SET name = $2, city = $3, street = $4, lat = $5, lng = $6, radius_m = $7,
			status = $8, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE id = $1`,
		hub.Id.String(),
		hub.Name,
		hub.City,
		hub.Street,
		hub.Lat,
		hub.Lng,
		hub.RadiusM,
		hub.Status,
	)

	return err
}

// Delete removes a row
func (r *HubRepositoryPostgres) Delete(id Id) error {
	_, err := r.db.Exec(`DELETE FROM hubs WHERE id = $1`, id.String())

	return err
}

// Add carries a product at a hub
func (r *HubAssortmentRepositoryPostgres) Add(hubId, productId Id) error {
	id, err := NewId()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		`INSERT INTO hub_assortment (id, hub_id, product_id) VALUES ($1, $2, $3)`,
		id.String(),
		hubId.String(),
		productId.String(),
	)

	return err
}

// Remove drops a product from a hub
func (r *HubAssortmentRepositoryPostgres) Remove(hubId, productId Id) error {
	_, err := r.db.Exec(
		`DELETE FROM hub_assortment WHERE hub_id = $1 AND product_id = $2`,
		hubId.String(),
		productId.String(),
	)

	return err
}

// ListByHub returns the assortment links of a hub
func (r *HubAssortmentRepositoryPostgres) ListByHub(hubId Id) ([]*HubAssortment, error) {
	return QueryMany(
		r.db,
		`SELECT id, hub_id, product_id, created_at FROM hub_assortment WHERE hub_id = $1`,
		scanHubAssortment,
		hubId.String(),
	)
}

// Create inserts a row
func (r *HubStockRepositoryPostgres) Create(stock *HubStock) error {
	if lo.IsEmpty(stock.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		stock.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO hub_stock (id, hub_id, product_id, quantity, reserved, reorder_point)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		stock.Id.String(),
		stock.HubId.String(),
		stock.ProductId.String(),
		stock.Quantity,
		stock.Reserved,
		stock.ReorderPoint,
	)

	return err
}

// GetByHubProduct returns the stock row of a product at a hub
func (r *HubStockRepositoryPostgres) GetByHubProduct(hubId, productId Id) (*HubStock, error) {
	return QueryOne(
		r.db,
		`SELECT `+hubStockColumns+` FROM hub_stock WHERE hub_id = $1 AND product_id = $2`,
		scanHubStock,
		hubId.String(),
		productId.String(),
	)
}

// ListByHub returns the stock rows of a hub
func (r *HubStockRepositoryPostgres) ListByHub(hubId Id) ([]*HubStock, error) {
	return QueryMany(
		r.db,
		`SELECT `+hubStockColumns+` FROM hub_stock WHERE hub_id = $1`,
		scanHubStock,
		hubId.String(),
	)
}

// ListLowStock returns rows at or below their reorder point
func (r *HubStockRepositoryPostgres) ListLowStock(hubId Id) ([]*HubStock, error) {
	return QueryMany(
		r.db,
		`SELECT `+hubStockColumns+` FROM hub_stock WHERE hub_id = $1 AND quantity <= reorder_point`,
		scanHubStock,
		hubId.String(),
	)
}

// Adjust shifts the quantity atomically, never below zero
func (r *HubStockRepositoryPostgres) Adjust(hubId, productId Id, delta int) error {
	_, err := r.db.Exec(
		`UPDATE hub_stock
		SET quantity = GREATEST(0, quantity + $3), updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE hub_id = $1 AND product_id = $2`,
		hubId.String(),
		productId.String(),
		delta,
	)

	return err
}

// Delete removes a row
func (r *HubStockRepositoryPostgres) Delete(id Id) error {
	_, err := r.db.Exec(`DELETE FROM hub_stock WHERE id = $1`, id.String())

	return err
}

// Record appends a movement to the ledger
func (r *StockMovementRepositoryPostgres) Record(movement *StockMovement) error {
	if lo.IsEmpty(movement.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		movement.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO stock_movements (id, hub_id, product_id, qty_change, reason, ref)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		movement.Id.String(),
		movement.HubId.String(),
		movement.ProductId.String(),
		movement.QtyChange,
		movement.Reason,
		movement.Ref,
	)

	return err
}

// ListByHubProduct returns the latest movements of a product at a hub
func (r *StockMovementRepositoryPostgres) ListByHubProduct(hubId, productId Id, limit int) ([]*StockMovement, error) {
	return QueryMany(
		r.db,
		`SELECT id, hub_id, product_id, qty_change, reason, ref, created_at
		FROM stock_movements WHERE hub_id = $1 AND product_id = $2 ORDER BY created_at DESC LIMIT $3`,
		scanStockMovement,
		hubId.String(),
		productId.String(),
		limit,
	)
}
