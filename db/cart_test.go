// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationCartRepository(t *testing.T) {
	database := openTestDB(t)
	carts := NewCartRepository(database)
	items := NewCartItemRepository(database)
	hubs := NewHubRepository(database)
	products := NewProductRepository(database)

	user := createTestUser(t, database)

	hub := &Hub{Name: "Kreuzberg", City: "Berlin", Lat: 52.49, Lng: 13.42, Status: HubStatusOpen}
	require.NoError(t, hubs.Create(hub))
	t.Cleanup(func() { _ = hubs.Delete(hub.Id) })

	product := &Product{SKU: "TEST-CART-001", Name: "Test Eggs", IsActive: true}
	require.NoError(t, products.Create(product))
	t.Cleanup(func() { _ = products.Delete(product.Id) })

	cart := &Cart{UserId: user.Id}
	require.NoError(t, carts.Create(cart))
	t.Cleanup(func() { _ = carts.Delete(cart.Id) })

	t.Run("get by id and user", func(t *testing.T) {
		got, err := carts.GetById(cart.Id)
		require.NoError(t, err)
		assert.Equal(t, user.Id, got.UserId)

		byUser, err := carts.GetByUser(user.Id)
		require.NoError(t, err)
		assert.Equal(t, cart.Id, byUser.Id)
	})

	t.Run("items", func(t *testing.T) {
		require.NoError(t, items.Add(&CartItem{CartId: cart.Id, ProductId: product.Id, HubId: hub.Id, Quantity: 2}))
		require.NoError(t, items.Add(&CartItem{CartId: cart.Id, ProductId: product.Id, HubId: hub.Id, Quantity: 1}))

		lines, err := items.ListByCart(cart.Id)
		require.NoError(t, err)
		require.Len(t, lines, 1)
		assert.Equal(t, 3, lines[0].Quantity)

		require.NoError(t, items.SetQuantity(cart.Id, product.Id, 5))

		lines, err = items.ListByCart(cart.Id)
		require.NoError(t, err)
		assert.Equal(t, 5, lines[0].Quantity)

		require.NoError(t, items.Remove(cart.Id, product.Id))

		lines, err = items.ListByCart(cart.Id)
		require.NoError(t, err)
		assert.Empty(t, lines)

		require.NoError(t, items.Add(&CartItem{CartId: cart.Id, ProductId: product.Id, HubId: hub.Id, Quantity: 1}))
		require.NoError(t, items.ClearCart(cart.Id))

		lines, err = items.ListByCart(cart.Id)
		require.NoError(t, err)
		assert.Empty(t, lines)
	})
}
