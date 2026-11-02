// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationDeliveryRepository(t *testing.T) {
	database := openTestDB(t)
	users := NewUserRepository(database)
	deliveries := NewDeliveryRepository(database)
	orders := NewOrderRepository(database)
	hubs := NewHubRepository(database)

	user := createTestUser(t, database)

	rider := createTestUser(t, database)
	rider.Role = UserRoleRider
	require.NoError(t, users.Update(rider))

	hub := &Hub{Name: "Moabit", City: "Berlin", Lat: 52.52, Lng: 13.33, Status: HubStatusOpen}
	require.NoError(t, hubs.Create(hub))
	t.Cleanup(func() { _ = hubs.Delete(hub.Id) })

	order := &Order{UserId: user.Id, HubId: hub.Id, Status: OrderStatusPacked, TotalCents: 800}
	require.NoError(t, orders.Create(order))
	t.Cleanup(func() {
		_, _ = database.Exec(`DELETE FROM orders WHERE id = $1`, order.Id.String())
	})

	delivery := &Delivery{OrderId: order.Id, Status: DeliveryStatusAssigned}
	require.NoError(t, deliveries.Create(delivery))

	t.Run("get by id and order", func(t *testing.T) {
		got, err := deliveries.GetById(delivery.Id)
		require.NoError(t, err)
		assert.Equal(t, order.Id, got.OrderId)

		byOrder, err := deliveries.GetByOrder(order.Id)
		require.NoError(t, err)
		assert.Equal(t, delivery.Id, byOrder.Id)
	})

	t.Run("assign and status", func(t *testing.T) {
		require.NoError(t, deliveries.Assign(delivery.Id, rider.Id))

		got, err := deliveries.GetById(delivery.Id)
		require.NoError(t, err)
		require.NotNil(t, got.RiderUserId)
		assert.Equal(t, rider.Id.String(), *got.RiderUserId)
		require.NotNil(t, got.AssignedAt)

		require.NoError(t, deliveries.UpdateStatus(delivery.Id, DeliveryStatusDelivered, timePtr(time.Now().UTC())))

		got, err = deliveries.GetById(delivery.Id)
		require.NoError(t, err)
		assert.Equal(t, DeliveryStatusDelivered, got.Status)
		require.NotNil(t, got.DeliveredAt)

		rides, err := deliveries.ListByRider(rider.Id, 10, 0)
		require.NoError(t, err)
		require.Len(t, rides, 1)
	})
}
