// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitShouldSkipAuth(t *testing.T) {
	cases := []struct {
		path string
		want bool
	}{
		{path: "/api/v1/public/_health", want: true},
		{path: "/assets/app.js", want: true},
		{path: "/api/v1/workspaces/ws-1/invites", want: false},
		{path: "/api/v1/me", want: false},
	}

	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			assert.Equal(t, tc.want, ShouldSkipAuth(tc.path))
		})
	}
}
