// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationPaymentRepository(t *testing.T) {
	database := openTestDB(t)
	payments := NewPaymentRepository(database)
	promotions := NewPromotionRepository(database)
	redemptions := NewVoucherRedemptionRepository(database)
	orders := NewOrderRepository(database)
	hubs := NewHubRepository(database)

	user := createTestUser(t, database)

	hub := &Hub{Name: "Wedding", City: "Berlin", Lat: 52.55, Lng: 13.36, Status: HubStatusOpen}
	require.NoError(t, hubs.Create(hub))
	t.Cleanup(func() { _ = hubs.Delete(hub.Id) })

	order := &Order{UserId: user.Id, HubId: hub.Id, Status: OrderStatusReceived, TotalCents: 1000}
	require.NoError(t, orders.Create(order))
	t.Cleanup(func() {
		_, _ = database.Exec(`DELETE FROM orders WHERE id = $1`, order.Id.String())
	})

	t.Run("payment", func(t *testing.T) {
		payment := &Payment{OrderId: order.Id, Provider: "stripe", AmountCents: 1000, Status: PaymentStatusPending}
		require.NoError(t, payments.Create(payment))

		got, err := payments.GetById(payment.Id)
		require.NoError(t, err)
		assert.Equal(t, PaymentStatusPending, got.Status)

		byOrder, err := payments.GetByOrder(order.Id)
		require.NoError(t, err)
		assert.Equal(t, payment.Id, byOrder.Id)

		require.NoError(t, payments.UpdateStatus(payment.Id, PaymentStatusPaid, timePtr(time.Now().UTC())))

		got, err = payments.GetById(payment.Id)
		require.NoError(t, err)
		assert.Equal(t, PaymentStatusPaid, got.Status)
		require.NotNil(t, got.PaidAt)
	})

	t.Run("promotion and redemption", func(t *testing.T) {
		promotion := &Promotion{Code: "WELCOME10", Kind: PromotionKindPercent, Value: 10, IsActive: true}
		require.NoError(t, promotions.Create(promotion))

		got, err := promotions.GetByCode("WELCOME10")
		require.NoError(t, err)
		assert.Equal(t, promotion.Id, got.Id)

		active, err := promotions.ListActive()
		require.NoError(t, err)
		assert.NotEmpty(t, active)

		redemption := &VoucherRedemption{PromotionId: promotion.Id, OrderId: order.Id, UserId: user.Id, AmountCents: 100}
		require.NoError(t, redemptions.Create(redemption))

		byOrder, err := redemptions.GetByOrder(order.Id)
		require.NoError(t, err)
		assert.Equal(t, 100, byOrder.AmountCents)

		byPromotion, err := redemptions.ListByPromotion(promotion.Id)
		require.NoError(t, err)
		require.Len(t, byPromotion, 1)
	})
}
