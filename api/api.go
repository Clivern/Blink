// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package api

import (
	"github.com/clivern/blink/db"
	"github.com/clivern/blink/module"
	"github.com/clivern/blink/pkg/fb"
	"github.com/clivern/blink/pkg/github"
	"github.com/clivern/blink/pkg/google"
	"github.com/clivern/blink/pkg/resend"

	"github.com/spf13/viper"
)

// Instance is the process-wide API wired at server startup.
var Instance *API

// API holds shared application modules for HTTP handlers.
type API struct {
	Auth      *module.Auth
	Profile   *module.Profile
	Setup     *module.Setup
	Settings  *module.Settings
	APIKey    *module.APIKey
	Me        *module.Me
	Workspace *module.Workspace
	Invite    *module.Invite
	Access    *module.Access
	Stats     *module.Stats
	Chat      *module.Chat
	KeyValue  *module.KeyValue

	GitHubOAuth   *github.OAuth
	GoogleOAuth   *google.OAuth
	FacebookOAuth *fb.OAuth
}

// New wires repositories and modules once for the process.
func New() *API {
	conn := db.GetDB()

	users := db.NewUserRepository(conn)
	sessions := db.NewSessionRepository(conn)
	configs := db.NewConfigRepository(conn)
	workspaces := db.NewWorkspaceRepository(conn)
	workspaceUsers := db.NewWorkspaceUserRepository(conn)
	invites := db.NewUserInviteRepository(conn)
	apiKeys := db.NewAPIKeyRepository(conn)
	accessKeys := db.NewAccessKeyRepository(conn)
	stats := db.NewWorkspaceStatsRepository(conn)
	integrations := db.NewIntegrationRepository(conn)
	events := db.NewEventRepository(conn)
	kv := db.NewWorkspaceKeyValueRepository(conn)

	a := &API{
		Auth:      module.NewAuth(users, sessions, configs),
		Profile:   module.NewProfile(users),
		Setup:     module.NewSetup(configs, users),
		Settings:  module.NewSettings(configs),
		APIKey:    module.NewAPIKey(apiKeys),
		Me:        module.NewMe(apiKeys, users, accessKeys),
		Workspace: module.NewWorkspace(workspaces, workspaceUsers, users),
		Invite:    module.NewInvite(invites, users, configs, workspaces, workspaceUsers, resend.NewMailer()),
		Access:    module.NewAccess(accessKeys, workspaces),
		Stats:     module.NewStats(workspaces, stats),
		KeyValue:  module.NewKeyValue(kv, workspaces),
		Chat: module.NewChat(
			db.NewChatRepository(conn),
			workspaces,
			module.NewChatTools(
				module.NewWorkspace(workspaces, workspaceUsers, users),
				module.NewKeyValue(kv, workspaces),
				module.NewEvent(events),
				module.NewIntegration(integrations),
			),
		),
		GitHubOAuth: github.NewOAuth(github.OAuthConfig{
			ClientID:     viper.GetString("app.oauth.github.client_id"),
			ClientSecret: viper.GetString("app.oauth.github.client_secret"),
			RedirectURL:  viper.GetString("app.oauth.github.redirect_url"),
			Scopes:       []string{"read:user", "user:email"},
			AllowSignup:  true,
		}),
		GoogleOAuth: google.NewOAuth(google.OAuthConfig{
			ClientID:     viper.GetString("app.oauth.google.client_id"),
			ClientSecret: viper.GetString("app.oauth.google.client_secret"),
			RedirectURL:  viper.GetString("app.oauth.google.redirect_url"),
			Scopes:       []string{"openid", "email", "profile"},
		}),
		FacebookOAuth: fb.NewOAuth(fb.OAuthConfig{
			ClientID:     viper.GetString("app.oauth.facebook.client_id"),
			ClientSecret: viper.GetString("app.oauth.facebook.client_secret"),
			RedirectURL:  viper.GetString("app.oauth.facebook.redirect_url"),
			Scopes:       []string{"email", "public_profile"},
		}),
	}

	Instance = a

	return a
}
