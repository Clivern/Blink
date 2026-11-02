// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationAddressRepository(t *testing.T) {
	database := openTestDB(t)
	users := NewUserRepository(database)
	repo := NewAddressRepository(database)

	user := createTestUser(t, database)
	user.Role = UserRoleCustomer
	require.NoError(t, users.Update(user))

	address := &Address{
		UserId: user.Id,
		Label:  stringPtr("home"),
		Street: "Torstr 1",
		City:   stringPtr("Berlin"),
		Lat:    52.52,
		Lng:    13.40,
	}
	require.NoError(t, repo.Create(address))
	t.Cleanup(func() { _ = repo.Delete(address.Id) })

	t.Run("get and list", func(t *testing.T) {
		got, err := repo.GetById(address.Id)
		require.NoError(t, err)
		assert.Equal(t, "Torstr 1", got.Street)

		items, err := repo.ListByUser(user.Id)
		require.NoError(t, err)
		require.Len(t, items, 1)
	})

	t.Run("update and delete", func(t *testing.T) {
		address.Street = "Torstr 2"
		require.NoError(t, repo.Update(address))

		got, err := repo.GetById(address.Id)
		require.NoError(t, err)
		assert.Equal(t, "Torstr 2", got.Street)

		require.NoError(t, repo.Delete(address.Id))

		items, err := repo.ListByUser(user.Id)
		require.NoError(t, err)
		assert.Empty(t, items)
	})
}
