// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"time"

	"github.com/samber/lo"
)

// Category is a node in the catalog taxonomy. Top-level categories
// have a nil ParentId; products link to leaf categories.
type Category struct {
	Id        Id
	ParentId  *string
	Name      string
	Slug      string
	Position  int
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CategoryRepository is the interface for category CRUD.
type CategoryRepository interface {
	Create(category *Category) error
	GetById(id Id) (*Category, error)
	GetBySlug(slug string) (*Category, error)
	ListByParent(parentId *string) ([]*Category, error)
	Update(category *Category) error
	Delete(id Id) error
}

// CategoryRepositoryPostgres implements CategoryRepository for PostgreSQL
type CategoryRepositoryPostgres struct {
	db *sql.DB
}

// NewCategoryRepository returns the repository for the current driver
func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryPostgres{db: db}
}

// scanCategory maps a row into a category. Column order must match the select list.
func scanCategory(category *Category) []any {
	return []any{
		&category.Id,
		&category.ParentId,
		&category.Name,
		&category.Slug,
		&category.Position,
		&category.IsActive,
		&category.CreatedAt,
		&category.UpdatedAt,
	}
}

const categoryColumns = `id, parent_id, name, slug, position, is_active, created_at, updated_at`

// --- Postgres ---

// Create inserts a row
func (r *CategoryRepositoryPostgres) Create(category *Category) error {
	if lo.IsEmpty(category.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		category.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO categories (id, parent_id, name, slug, position, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		category.Id.String(),
		category.ParentId,
		category.Name,
		category.Slug,
		category.Position,
		category.IsActive,
	)

	return err
}

// GetById returns a category by Id
func (r *CategoryRepositoryPostgres) GetById(id Id) (*Category, error) {
	return QueryOne(
		r.db,
		`SELECT `+categoryColumns+` FROM categories WHERE id = $1`,
		scanCategory,
		id.String(),
	)
}

// GetBySlug returns a category by slug
func (r *CategoryRepositoryPostgres) GetBySlug(slug string) (*Category, error) {
	return QueryOne(
		r.db,
		`SELECT `+categoryColumns+` FROM categories WHERE slug = $1`,
		scanCategory,
		slug,
	)
}

// ListByParent returns categories with the given parent, ordered for display
func (r *CategoryRepositoryPostgres) ListByParent(parentId *string) ([]*Category, error) {
	if parentId == nil {
		return QueryMany(
			r.db,
			`SELECT `+categoryColumns+` FROM categories WHERE parent_id IS NULL ORDER BY position, name`,
			scanCategory,
		)
	}

	return QueryMany(
		r.db,
		`SELECT `+categoryColumns+` FROM categories WHERE parent_id = $1 ORDER BY position, name`,
		scanCategory,
		*parentId,
	)
}

// Update rewrites a row
func (r *CategoryRepositoryPostgres) Update(category *Category) error {
	_, err := r.db.Exec(
		`UPDATE categories
		SET parent_id = $2, name = $3, slug = $4, position = $5, is_active = $6, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE id = $1`,
		category.Id.String(),
		category.ParentId,
		category.Name,
		category.Slug,
		category.Position,
		category.IsActive,
	)

	return err
}

// Delete removes a row
func (r *CategoryRepositoryPostgres) Delete(id Id) error {
	_, err := r.db.Exec(`DELETE FROM categories WHERE id = $1`, id.String())

	return err
}
