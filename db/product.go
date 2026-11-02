// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"time"

	"github.com/samber/lo"
)

// Product is a sellable grocery item in the central catalog.
type Product struct {
	Id        Id
	SKU       string
	EAN       *string
	Brand     *string
	Name      string
	PackSize  *string
	Unit      *string
	VATClass  *string
	ImageURL  *string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ProductCategory links a product to a leaf category.
type ProductCategory struct {
	Id         Id
	ProductId  Id
	CategoryId Id
	CreatedAt  time.Time
}

// Price is a product price valid for a time window. The current
// price is the row with the latest valid_from and no valid_to.
type Price struct {
	Id        Id
	ProductId Id
	Cents     int
	ValidFrom time.Time
	ValidTo   *time.Time
	CreatedAt time.Time
}

// ProductRepository is the interface for product CRUD.
type ProductRepository interface {
	Create(product *Product) error
	GetById(id Id) (*Product, error)
	GetBySKU(sku string) (*Product, error)
	ListActive(limit, offset int) ([]*Product, error)
	Update(product *Product) error
	Delete(id Id) error
}

// ProductCategoryRepository is the interface for product/category links.
type ProductCategoryRepository interface {
	Add(productId, categoryId Id) error
	Remove(productId, categoryId Id) error
	ListByProduct(productId Id) ([]*ProductCategory, error)
	ListByCategory(categoryId Id) ([]*ProductCategory, error)
}

// PriceRepository is the interface for product price history.
type PriceRepository interface {
	Create(price *Price) error
	GetCurrent(productId Id) (*Price, error)
	ListByProduct(productId Id) ([]*Price, error)
}

// ProductRepositoryPostgres implements ProductRepository for PostgreSQL
type ProductRepositoryPostgres struct {
	db *sql.DB
}

// ProductCategoryRepositoryPostgres implements ProductCategoryRepository for PostgreSQL
type ProductCategoryRepositoryPostgres struct {
	db *sql.DB
}

// PriceRepositoryPostgres implements PriceRepository for PostgreSQL
type PriceRepositoryPostgres struct {
	db *sql.DB
}

// NewProductRepository returns the repository for the current driver
func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryPostgres{db: db}
}

// NewProductCategoryRepository returns the repository for the current driver
func NewProductCategoryRepository(db *sql.DB) ProductCategoryRepository {
	return &ProductCategoryRepositoryPostgres{db: db}
}

// NewPriceRepository returns the repository for the current driver
func NewPriceRepository(db *sql.DB) PriceRepository {
	return &PriceRepositoryPostgres{db: db}
}

// scanProduct maps a row into a product. Column order must match the select list.
func scanProduct(product *Product) []any {
	return []any{
		&product.Id,
		&product.SKU,
		&product.EAN,
		&product.Brand,
		&product.Name,
		&product.PackSize,
		&product.Unit,
		&product.VATClass,
		&product.ImageURL,
		&product.IsActive,
		&product.CreatedAt,
		&product.UpdatedAt,
	}
}

// scanProductCategory maps a row into a product/category link.
func scanProductCategory(link *ProductCategory) []any {
	return []any{
		&link.Id,
		&link.ProductId,
		&link.CategoryId,
		&link.CreatedAt,
	}
}

// scanPrice maps a row into a price. Column order must match the select list.
func scanPrice(price *Price) []any {
	return []any{
		&price.Id,
		&price.ProductId,
		&price.Cents,
		&price.ValidFrom,
		&price.ValidTo,
		&price.CreatedAt,
	}
}

const productColumns = `id, sku, ean, brand, name, pack_size, unit, vat_class, image_url, is_active, created_at, updated_at`

// --- Postgres ---

// Create inserts a row
func (r *ProductRepositoryPostgres) Create(product *Product) error {
	if lo.IsEmpty(product.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		product.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO products (id, sku, ean, brand, name, pack_size, unit, vat_class, image_url, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		product.Id.String(),
		product.SKU,
		product.EAN,
		product.Brand,
		product.Name,
		product.PackSize,
		product.Unit,
		product.VATClass,
		product.ImageURL,
		product.IsActive,
	)

	return err
}

// GetById returns a product by Id
func (r *ProductRepositoryPostgres) GetById(id Id) (*Product, error) {
	return QueryOne(
		r.db,
		`SELECT `+productColumns+` FROM products WHERE id = $1`,
		scanProduct,
		id.String(),
	)
}

// GetBySKU returns a product by SKU
func (r *ProductRepositoryPostgres) GetBySKU(sku string) (*Product, error) {
	return QueryOne(
		r.db,
		`SELECT `+productColumns+` FROM products WHERE sku = $1`,
		scanProduct,
		sku,
	)
}

// ListActive returns active products alphabetically
func (r *ProductRepositoryPostgres) ListActive(limit, offset int) ([]*Product, error) {
	return QueryMany(
		r.db,
		`SELECT `+productColumns+` FROM products WHERE is_active = true ORDER BY name LIMIT $1 OFFSET $2`,
		scanProduct,
		limit,
		offset,
	)
}

// Update rewrites a row
func (r *ProductRepositoryPostgres) Update(product *Product) error {
	_, err := r.db.Exec(
		`UPDATE products
		SET sku = $2, ean = $3, brand = $4, name = $5, pack_size = $6, unit = $7,
			vat_class = $8, image_url = $9, is_active = $10, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
		WHERE id = $1`,
		product.Id.String(),
		product.SKU,
		product.EAN,
		product.Brand,
		product.Name,
		product.PackSize,
		product.Unit,
		product.VATClass,
		product.ImageURL,
		product.IsActive,
	)

	return err
}

// Delete removes a row
func (r *ProductRepositoryPostgres) Delete(id Id) error {
	_, err := r.db.Exec(`DELETE FROM products WHERE id = $1`, id.String())

	return err
}

// Add links a product to a category
func (r *ProductCategoryRepositoryPostgres) Add(productId, categoryId Id) error {
	id, err := NewId()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		`INSERT INTO product_categories (id, product_id, category_id) VALUES ($1, $2, $3)`,
		id.String(),
		productId.String(),
		categoryId.String(),
	)

	return err
}

// Remove unlinks a product from a category
func (r *ProductCategoryRepositoryPostgres) Remove(productId, categoryId Id) error {
	_, err := r.db.Exec(
		`DELETE FROM product_categories WHERE product_id = $1 AND category_id = $2`,
		productId.String(),
		categoryId.String(),
	)

	return err
}

// ListByProduct returns the category links of a product
func (r *ProductCategoryRepositoryPostgres) ListByProduct(productId Id) ([]*ProductCategory, error) {
	return QueryMany(
		r.db,
		`SELECT id, product_id, category_id, created_at FROM product_categories WHERE product_id = $1`,
		scanProductCategory,
		productId.String(),
	)
}

// ListByCategory returns the product links of a category
func (r *ProductCategoryRepositoryPostgres) ListByCategory(categoryId Id) ([]*ProductCategory, error) {
	return QueryMany(
		r.db,
		`SELECT id, product_id, category_id, created_at FROM product_categories WHERE category_id = $1`,
		scanProductCategory,
		categoryId.String(),
	)
}

// Create inserts a row
func (r *PriceRepositoryPostgres) Create(price *Price) error {
	if lo.IsEmpty(price.Id) {
		id, err := NewId()
		if err != nil {
			return err
		}

		price.Id = id
	}

	_, err := r.db.Exec(
		`INSERT INTO prices (id, product_id, price_cents, valid_from, valid_to)
		VALUES ($1, $2, $3, $4, $5)`,
		price.Id.String(),
		price.ProductId.String(),
		price.Cents,
		price.ValidFrom,
		price.ValidTo,
	)

	return err
}

// GetCurrent returns the latest open price of a product
func (r *PriceRepositoryPostgres) GetCurrent(productId Id) (*Price, error) {
	return QueryOne(
		r.db,
		`SELECT id, product_id, price_cents, valid_from, valid_to, created_at
		FROM prices WHERE product_id = $1 AND valid_to IS NULL ORDER BY valid_from DESC LIMIT 1`,
		scanPrice,
		productId.String(),
	)
}

// ListByProduct returns the price history of a product, newest first
func (r *PriceRepositoryPostgres) ListByProduct(productId Id) ([]*Price, error) {
	return QueryMany(
		r.db,
		`SELECT id, product_id, price_cents, valid_from, valid_to, created_at
		FROM prices WHERE product_id = $1 ORDER BY valid_from DESC`,
		scanPrice,
		productId.String(),
	)
}
