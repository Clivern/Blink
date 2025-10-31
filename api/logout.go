// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/clivern/blink/locale"
	"github.com/clivern/blink/pkg/util"

	"github.com/rs/zerolog/log"
)

// LogoutAction logs the user out and revokes their session.
func (a *API) LogoutAction(w http.ResponseWriter, r *http.Request) {
	util.DeleteCookie(w, "_blink_session")

	user, ok := a.GetUser(r)
	if !ok {
		log.Info().
			Msg("New logout request")
		a.WriteJSON(w, http.StatusOK, map[string]any{
			"successMessage": locale.TR(r, "logout_successful"),
		})
		return
	}

	log.Info().
		Str("userId", user.Id.String()).
		Msg("New logout request")

	err := a.Auth.Logout(user.Id)
	if err != nil {
		log.Error().
			Str("userId", user.Id.String()).
			Err(err).
			Msg("Failed to revoke session")
	}

	a.WriteJSON(w, http.StatusOK, map[string]any{
		"successMessage": locale.TR(r, "logout_successful"),
	})
}
