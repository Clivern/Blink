// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationProductRepository(t *testing.T) {
	database := openTestDB(t)
	repo := NewProductRepository(database)
	links := NewProductCategoryRepository(database)
	prices := NewPriceRepository(database)
	categories := NewCategoryRepository(database)

	category := &Category{Name: "Snacks", Slug: "snacks-cat", Position: 1, IsActive: true}
	require.NoError(t, categories.Create(category))
	t.Cleanup(func() { _ = categories.Delete(category.Id) })

	product := &Product{SKU: "TEST-001", Name: "Test Milk", Brand: stringPtr("Test"), IsActive: true}
	require.NoError(t, repo.Create(product))
	t.Cleanup(func() { _ = repo.Delete(product.Id) })

	t.Run("get by id and sku", func(t *testing.T) {
		got, err := repo.GetById(product.Id)
		require.NoError(t, err)
		assert.Equal(t, "Test Milk", got.Name)

		bySKU, err := repo.GetBySKU("TEST-001")
		require.NoError(t, err)
		assert.Equal(t, product.Id, bySKU.Id)
	})

	t.Run("list active and update", func(t *testing.T) {
		items, err := repo.ListActive(10, 0)
		require.NoError(t, err)
		assert.NotEmpty(t, items)

		product.Name = "Test Milk 1L"
		require.NoError(t, repo.Update(product))

		got, err := repo.GetById(product.Id)
		require.NoError(t, err)
		assert.Equal(t, "Test Milk 1L", got.Name)
	})

	t.Run("category links", func(t *testing.T) {
		require.NoError(t, links.Add(product.Id, category.Id))

		byProduct, err := links.ListByProduct(product.Id)
		require.NoError(t, err)
		require.Len(t, byProduct, 1)

		byCategory, err := links.ListByCategory(category.Id)
		require.NoError(t, err)
		require.Len(t, byCategory, 1)

		require.NoError(t, links.Remove(product.Id, category.Id))

		byProduct, err = links.ListByProduct(product.Id)
		require.NoError(t, err)
		assert.Empty(t, byProduct)
	})

	t.Run("price history", func(t *testing.T) {
		price := &Price{ProductId: product.Id, Cents: 199, ValidFrom: time.Now().UTC()}
		require.NoError(t, prices.Create(price))

		current, err := prices.GetCurrent(product.Id)
		require.NoError(t, err)
		assert.Equal(t, 199, current.Cents)

		history, err := prices.ListByProduct(product.Id)
		require.NoError(t, err)
		require.Len(t, history, 1)
	})
}
