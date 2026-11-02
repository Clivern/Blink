// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/clivern/blink/migration"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	testDBOnce sync.Once
	testDBConn *sql.DB
	testDBErr  error
)

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()

	testDBOnce.Do(func() {
		port := 5432
		if raw := os.Getenv("BLINK_DATABASE_PORT"); raw != "" {
			parsed, err := strconv.Atoi(raw)
			if err != nil {
				testDBErr = fmt.Errorf("invalid BLINK_DATABASE_PORT: %w", err)
				return
			}

			port = parsed
		}

		config := DatabaseConfig{
			Driver:   envOr("BLINK_DATABASE_DRIVER", "postgres"),
			Host:     envOr("BLINK_DATABASE_HOST", "localhost"),
			Port:     port,
			Username: envOr("BLINK_DATABASE_USERNAME", "postgres"),
			Password: envOr("BLINK_DATABASE_PASSWORD", "postgres"),
			Database: envOr("BLINK_DATABASE_NAME", "blink"),
		}

		var conn *Connection
		var err error
		for attempt := 1; attempt <= 5; attempt++ {
			conn, err = NewConnection(config)
			if err == nil {
				break
			}

			time.Sleep(500 * time.Millisecond)
		}
		if err != nil {
			testDBErr = err
			return
		}

		mgr := migration.NewManager(conn.DB)
		for _, item := range migration.GetAll() {
			mgr.Register(item)
		}
		if err := mgr.Up(); err != nil {
			_ = conn.Close()
			testDBErr = err
			return
		}

		testDBConn = conn.DB
	})

	if testDBErr != nil {
		t.Fatalf("postgres unavailable: %v", testDBErr)
	}

	return testDBConn
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func stringPtr(value string) *string {
	return &value
}

func timePtr(value time.Time) *time.Time {
	return &value
}

func createTestWorkspace(t *testing.T, database *sql.DB) *Workspace {
	t.Helper()

	workspace := &Workspace{
		Name:   "test",
		Handle: fmt.Sprintf("t-%s", uuid.NewString()),
	}
	require.NoError(t, NewWorkspaceRepository(database).Create(workspace))
	t.Cleanup(func() {
		_ = NewWorkspaceRepository(database).Delete(workspace.Id)
	})

	return workspace
}

func createTestUser(t *testing.T, database *sql.DB) *User {
	t.Helper()

	user := &User{
		Name:        "tester",
		Email:       fmt.Sprintf("u-%s@t.test", uuid.NewString()[:8]),
		Password:    "hash",
		Provider:    UserProviderLocal,
		Role:        UserRoleRegular,
		IsActive:    true,
		Language:    UserLanguageEN,
		Theme:       UserThemeDefault,
		LastLoginAt: time.Now().UTC(),
	}
	require.NoError(t, NewUserRepository(database).Create(user))
	t.Cleanup(func() {
		_ = NewUserRepository(database).Delete(user.Id)
	})

	return user
}
