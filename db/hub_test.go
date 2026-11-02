// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationHubRepository(t *testing.T) {
	database := openTestDB(t)
	hubs := NewHubRepository(database)
	assortment := NewHubAssortmentRepository(database)
	stocks := NewHubStockRepository(database)
	movements := NewStockMovementRepository(database)
	products := NewProductRepository(database)

	hub := &Hub{Name: "Mitte", City: "Berlin", Lat: 52.52, Lng: 13.40, RadiusM: 3000, Status: HubStatusOpen}
	require.NoError(t, hubs.Create(hub))
	t.Cleanup(func() { _ = hubs.Delete(hub.Id) })

	product := &Product{SKU: "TEST-HUB-001", Name: "Test Bread", IsActive: true}
	require.NoError(t, products.Create(product))
	t.Cleanup(func() { _ = products.Delete(product.Id) })

	t.Run("get list and update", func(t *testing.T) {
		got, err := hubs.GetById(hub.Id)
		require.NoError(t, err)
		assert.Equal(t, "Mitte", got.Name)

		items, err := hubs.List(10, 0)
		require.NoError(t, err)
		assert.NotEmpty(t, items)

		hub.Status = HubStatusClosed
		require.NoError(t, hubs.Update(hub))

		got, err = hubs.GetById(hub.Id)
		require.NoError(t, err)
		assert.Equal(t, HubStatusClosed, got.Status)
	})

	t.Run("assortment", func(t *testing.T) {
		require.NoError(t, assortment.Add(hub.Id, product.Id))

		items, err := assortment.ListByHub(hub.Id)
		require.NoError(t, err)
		require.Len(t, items, 1)

		require.NoError(t, assortment.Remove(hub.Id, product.Id))

		items, err = assortment.ListByHub(hub.Id)
		require.NoError(t, err)
		assert.Empty(t, items)
	})

	t.Run("stock and movements", func(t *testing.T) {
		stock := &HubStock{HubId: hub.Id, ProductId: product.Id, Quantity: 10, ReorderPoint: 5}
		require.NoError(t, stocks.Create(stock))
		t.Cleanup(func() { _ = stocks.Delete(stock.Id) })

		got, err := stocks.GetByHubProduct(hub.Id, product.Id)
		require.NoError(t, err)
		assert.Equal(t, 10, got.Quantity)

		require.NoError(t, stocks.Adjust(hub.Id, product.Id, -20))

		got, err = stocks.GetByHubProduct(hub.Id, product.Id)
		require.NoError(t, err)
		assert.Equal(t, 0, got.Quantity)

		low, err := stocks.ListLowStock(hub.Id)
		require.NoError(t, err)
		assert.NotEmpty(t, low)

		all, err := stocks.ListByHub(hub.Id)
		require.NoError(t, err)
		assert.NotEmpty(t, all)

		require.NoError(t, movements.Record(&StockMovement{
			HubId:     hub.Id,
			ProductId: product.Id,
			QtyChange: -10,
			Reason:    StockMovementSale,
		}))

		ledger, err := movements.ListByHubProduct(hub.Id, product.Id, 10)
		require.NoError(t, err)
		require.Len(t, ledger, 1)
		assert.Equal(t, -10, ledger[0].QtyChange)
	})
}
