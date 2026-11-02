// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationOrderRepository(t *testing.T) {
	database := openTestDB(t)
	users := NewUserRepository(database)
	orders := NewOrderRepository(database)
	lines := NewOrderItemRepository(database)
	tasks := NewPickTaskRepository(database)
	hubs := NewHubRepository(database)
	products := NewProductRepository(database)
	addresses := NewAddressRepository(database)

	user := createTestUser(t, database)

	picker := createTestUser(t, database)
	picker.Role = UserRolePicker
	require.NoError(t, users.Update(picker))

	hub := &Hub{Name: "Neukolln", City: "Berlin", Lat: 52.47, Lng: 13.43, Status: HubStatusOpen}
	require.NoError(t, hubs.Create(hub))
	t.Cleanup(func() { _ = hubs.Delete(hub.Id) })

	product := &Product{SKU: "TEST-ORDER-001", Name: "Test Butter", IsActive: true}
	require.NoError(t, products.Create(product))
	t.Cleanup(func() { _ = products.Delete(product.Id) })

	address := &Address{UserId: user.Id, Street: "Karl-Marx-Str 1", Lat: 52.47, Lng: 13.43}
	require.NoError(t, addresses.Create(address))
	t.Cleanup(func() { _ = addresses.Delete(address.Id) })

	order := &Order{
		UserId:        user.Id,
		HubId:         hub.Id,
		AddressId:     stringPtr(address.Id.String()),
		Status:        OrderStatusReceived,
		SubtotalCents: 500,
		TotalCents:    500,
	}
	require.NoError(t, orders.Create(order))
	t.Cleanup(func() {
		_, _ = database.Exec(`DELETE FROM orders WHERE id = $1`, order.Id.String())
	})

	t.Run("get list and status", func(t *testing.T) {
		got, err := orders.GetById(order.Id)
		require.NoError(t, err)
		assert.Equal(t, OrderStatusReceived, got.Status)

		mine, err := orders.ListByUser(user.Id, 10, 0)
		require.NoError(t, err)
		require.Len(t, mine, 1)

		hubOrders, err := orders.ListByHub(hub.Id, OrderStatusReceived, 10, 0)
		require.NoError(t, err)
		require.Len(t, hubOrders, 1)

		require.NoError(t, orders.UpdateStatus(order.Id, OrderStatusPicking))

		got, err = orders.GetById(order.Id)
		require.NoError(t, err)
		assert.Equal(t, OrderStatusPicking, got.Status)
	})

	t.Run("items", func(t *testing.T) {
		item := &OrderItem{OrderId: order.Id, ProductId: product.Id, Quantity: 2, PriceCents: 250}
		require.NoError(t, lines.Create(item))

		items, err := lines.ListByOrder(order.Id)
		require.NoError(t, err)
		require.Len(t, items, 1)

		require.NoError(t, lines.UpdatePicked(item.Id, 1, SubstitutionSubstituted))

		items, err = lines.ListByOrder(order.Id)
		require.NoError(t, err)
		assert.Equal(t, 1, items[0].PickedQuantity)
		assert.Equal(t, SubstitutionSubstituted, items[0].Substitution)
	})

	t.Run("pick task", func(t *testing.T) {
		task := &PickTask{OrderId: order.Id, Status: PickTaskStatusPending}
		require.NoError(t, tasks.Create(task))

		got, err := tasks.GetByOrder(order.Id)
		require.NoError(t, err)
		assert.Equal(t, order.Id, got.OrderId)

		byId, err := tasks.GetById(task.Id)
		require.NoError(t, err)
		assert.Equal(t, task.Id, byId.Id)

		require.NoError(t, tasks.Assign(task.Id, picker.Id))

		got, err = tasks.GetById(task.Id)
		require.NoError(t, err)
		assert.Equal(t, PickTaskStatusPicking, got.Status)
		require.NotNil(t, got.PickerUserId)

		require.NoError(t, tasks.Complete(task.Id))

		got, err = tasks.GetById(task.Id)
		require.NoError(t, err)
		assert.Equal(t, PickTaskStatusDone, got.Status)
		require.NotNil(t, got.FinishedAt)
	})
}
