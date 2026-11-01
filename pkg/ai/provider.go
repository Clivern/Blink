// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package ai

import (
	"github.com/clivern/blink/pkg/ai/openrouter"
	"github.com/spf13/viper"
)

const ProviderOpenRouter = "openrouter"

// IsConfigured reports whether an AI provider API key is set.
func IsConfigured() bool {
	return viper.GetString("app.ai.api_key") != ""
}

// NewChatClient returns the chat client for the app.ai.provider config.
func NewChatClient() ChatClient {
	apiKey := viper.GetString("app.ai.api_key")
	model := viper.GetString("app.ai.chat_model")

	return NewOpenRouterChat(openrouter.New(openrouter.Config{APIKey: apiKey, Model: model}))
}
