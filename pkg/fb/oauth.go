// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package fb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	OauthAuthorizeURL = "https://www.facebook.com/v26.0/dialog/oauth"
	OauthTokenURL     = "https://graph.facebook.com/v26.0/oauth/access_token"
	OauthUserURL      = "https://graph.facebook.com/v26.0/me"
)

// OAuthConfig holds Facebook OAuth app credentials.
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

// OAuth exchanges authorization codes for access tokens.
type OAuth struct {
	cfg OAuthConfig
}

// Token is an OAuth access token response.
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// UserInfo is the authenticated Facebook user.
type UserInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NewOAuth returns an OAuth helper for the given app configuration.
func NewOAuth(cfg OAuthConfig) *OAuth {
	return &OAuth{cfg: cfg}
}

// AuthorizeURL builds the Facebook authorization URL for the provided state.
func (o *OAuth) AuthorizeURL(state string) string {
	q := url.Values{
		"client_id":    {o.cfg.ClientID},
		"redirect_uri": {o.cfg.RedirectURL},
		"state":        {state},
		"scope":        {strings.Join(o.cfg.Scopes, ",")},
	}

	return OauthAuthorizeURL + "?" + q.Encode()
}

// Exchange validates state and exchanges an authorization code for an access token.
func (o *OAuth) Exchange(ctx context.Context, code, state, expectedState string) (*Token, error) {
	if state != expectedState {
		return nil, ErrInvalidOAuthState
	}

	q := url.Values{
		"client_id":     {o.cfg.ClientID},
		"redirect_uri":  {o.cfg.RedirectURL},
		"client_secret": {o.cfg.ClientSecret},
		"code":          {code},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, OauthTokenURL+"?"+q.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("facebook oauth build request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("facebook oauth request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("facebook oauth read body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("facebook oauth token: status %d: %s", resp.StatusCode, body)
	}

	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, fmt.Errorf("facebook oauth decode body: %w", err)
	}

	return &token, nil
}

// User fetches the authenticated Facebook user.
func (o *OAuth) User(ctx context.Context, accessToken string) (*UserInfo, error) {
	q := url.Values{
		"fields":       {"id,name,email"},
		"access_token": {accessToken},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, OauthUserURL+"?"+q.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("facebook oauth user request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("facebook oauth user: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("facebook oauth user body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("facebook oauth user: status %d: %s", resp.StatusCode, body)
	}

	var user UserInfo
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("facebook oauth user decode: %w", err)
	}

	return &user, nil
}
