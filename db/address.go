// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"time"

	"github.com/samber/lo"
)

// Address is a customer delivery address. HubId is the dark
// store resolved to serve the door at save time.
type Address struct {
	Id         Id
	UserId     Id
	Label      *string
	Street     string
	City       *string
	PostalCode *string
	Lat        float64
	Lng        float64
	HubId      *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// AddressRepository is the interface for address CRUD.
type AddressRepository interface {
	Create(address *Address) error
	GetById(id Id) (*Address, error)
	ListByUser(userId Id) ([]*Address, error)
	Update(address *Address) error
	Delete(id Id) error
}

// AddressRepositoryPostgres implements AddressRepository for PostgreSQL
type AddressRepositoryPostgres struct {
	db *sql.DB
}

// NewAddressRepository returns the repository for the current driver
func NewAddressRepository(db *sql.DB) AddressRepository {
	return &AddressRepositoryPostgres{db: db}
}

// scanAddress maps a row into an address. Column order must match the select list.
func scanAddress(address *Address) []any {
	return []any{
		&address.Id,
		&address.UserId,
		&address.Label,
		&address.Street,
		&address.City,
		&address.PostalCode,
		&address.Lat,
		&address.Lng,
		&address.HubId,
		&address.CreatedAt,
		&address.UpdatedAt,
	}
}

const addressColumns = `id, user_id, label, street, city, postal_code, lat, lng, hub_id, created_at, updated_at`

// --- Postgres ---

// Create inserts a row
func (r *AddressRepositoryPostgres) Create(address *Address) error {
	if lo.IsEmpty(address.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		address.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO addresses (id, user_id, label, street, city, postal_code, lat, lng, hub_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		address.Id.String(),
		address.UserId.String(),
		address.Label,
		address.Street,
		address.City,
		address.PostalCode,
		address.Lat,
		address.Lng,
		address.HubId,
	)

	return err
}

// GetById returns an address by Id
func (r *AddressRepositoryPostgres) GetById(id Id) (*Address, error) {
	return QueryOne(
		r.db,
		`SELECT `+addressColumns+` FROM addresses WHERE id = $1`,
		scanAddress,
		id.String(),
	)
}

// ListByUser returns the addresses of a user
func (r *AddressRepositoryPostgres) ListByUser(userId Id) ([]*Address, error) {
	return QueryMany(
		r.db,
		`SELECT `+addressColumns+` FROM addresses WHERE user_id = $1 ORDER BY created_at`,
		scanAddress,
		userId.String(),
	)
}

// Update rewrites a row
func (r *AddressRepositoryPostgres) Update(address *Address) error {
	_, err := r.db.Exec(
		`UPDATE addresses
		SET label = $2, street = $3, city = $4, postal_code = $5, lat = $6, lng = $7,
			hub_id = $8, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE id = $1`,
		address.Id.String(),
		address.Label,
		address.Street,
		address.City,
		address.PostalCode,
		address.Lat,
		address.Lng,
		address.HubId,
	)

	return err
}

// Delete removes a row
func (r *AddressRepositoryPostgres) Delete(id Id) error {
	_, err := r.db.Exec(`DELETE FROM addresses WHERE id = $1`, id.String())

	return err
}
