// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationCategoryRepository(t *testing.T) {
	database := openTestDB(t)
	repo := NewCategoryRepository(database)

	root := &Category{Name: "Fresh", Slug: "fresh-cat", Position: 1, IsActive: true}
	require.NoError(t, repo.Create(root))

	child := &Category{Name: "Dairy", Slug: "dairy-cat", Position: 1, IsActive: true}
	child.ParentId = stringPtr(root.Id.String())
	require.NoError(t, repo.Create(child))

	t.Run("get by id and slug", func(t *testing.T) {
		got, err := repo.GetById(child.Id)
		require.NoError(t, err)
		assert.Equal(t, "Dairy", got.Name)
		require.NotNil(t, got.ParentId)
		assert.Equal(t, root.Id.String(), *got.ParentId)

		bySlug, err := repo.GetBySlug("dairy-cat")
		require.NoError(t, err)
		assert.Equal(t, child.Id, bySlug.Id)
	})

	t.Run("list by parent", func(t *testing.T) {
		roots, err := repo.ListByParent(nil)
		require.NoError(t, err)
		assert.NotEmpty(t, roots)

		children, err := repo.ListByParent(stringPtr(root.Id.String()))
		require.NoError(t, err)
		require.Len(t, children, 1)
		assert.Equal(t, child.Id, children[0].Id)
	})

	t.Run("update and delete", func(t *testing.T) {
		child.Position = 5
		require.NoError(t, repo.Update(child))

		got, err := repo.GetById(child.Id)
		require.NoError(t, err)
		assert.Equal(t, 5, got.Position)

		require.NoError(t, repo.Delete(child.Id))
		require.NoError(t, repo.Delete(root.Id))
	})
}
