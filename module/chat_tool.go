// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package module

import (
	"github.com/clivern/blink/db"

	"github.com/samber/lo"
)

const (
	ChatToolLimit    = 50
	ChatSummaryLimit = 1000
)

// ChatTool is a dashboard action the assistant may run for the user, gated by the user's permission.
type ChatTool struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Arguments   map[string]string `json:"arguments,omitempty"`
	Permission  string            `json:"-"`
	Run         ChatToolRun       `json:"-"`
}

// ChatToolRun runs a tool in a workspace, for the chat's user, with the arguments the model chose.
type ChatToolRun func(workspaceId, userId db.Id, args map[string]string) (any, error)

// ChatTools are the modules the dashboard tools read from.
type ChatTools struct {
	Workspace   *Workspace
	KeyValue    *KeyValue
	Event       *Event
	Integration *Integration
}

// ChatWorkspaceSummary is the workspace with counters of what it holds, for questions about the workspace.
type ChatWorkspaceSummary struct {
	Workspace *WorkspaceResponse `json:"workspace"`
	KeyValues int64              `json:"kvKeys"`
}

// ChatIntegration is an integration without its URL, site or index
type ChatIntegration struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Events    []string `json:"events"`
	IsActive  bool     `json:"isActive"`
	HasSecret bool     `json:"hasSecret"`
}

// NewChatTools returns the dashboard tools backed by modules.
func NewChatTools(
	workspace *Workspace,
	keyValue *KeyValue,
	event *Event,
	integration *Integration,
) *ChatTools {
	return &ChatTools{
		Workspace:   workspace,
		KeyValue:    keyValue,
		Event:       event,
		Integration: integration,
	}
}

// List returns every dashboard tool the assistant knows.
func (t *ChatTools) List() []*ChatTool {
	return []*ChatTool{
		{
			Name:        "get_workspace",
			Description: "Get this workspace's name, id, handle, member count and the user's own role, with the KV key count.",
			Permission:  CanGetWorkspace,
			Run:         t.WorkspaceSummary,
		},
		{
			Name:        "list_members",
			Description: "List the workspace members with their name, email and role.",
			Permission:  CanListWorkspaceMembers,
			Run:         t.ListMembers,
		},
		{
			Name:        "list_kv",
			Description: "List the workspace key-value config, optionally under a key prefix.",
			Arguments:   map[string]string{"prefix": "optional key prefix"},
			Permission:  CanListKeyValue,
			Run:         t.ListKeyValue,
		},
		{
			Name:        "list_integrations",
			Description: "List the integrations that deliver workspace events, such as Slack or webhooks, with the events each sends and whether it is active.",
			Permission:  CanListIntegrations,
			Run:         t.ListIntegrations,
		},
		{
			Name:        "list_events",
			Description: "List the latest workspace events.",
			Permission:  CanListEvents,
			Run:         t.ListEvents,
		},
	}
}

// WorkspaceSummary returns the workspace with its KV key count.
func (t *ChatTools) WorkspaceSummary(workspaceId, userId db.Id, _ map[string]string) (any, error) {
	workspace, err := t.Workspace.GetWorkspace(workspaceId, &db.User{Id: userId})
	if err != nil {
		return nil, err
	}

	keyValues, err := t.KeyValue.ListKeyValue(workspaceId, "")
	if err != nil {
		return nil, err
	}

	summary := &ChatWorkspaceSummary{
		Workspace: workspace,
		KeyValues: keyValues.Total,
	}

	return summary, nil
}

// ListMembers returns the workspace members with their name, email and role.
func (t *ChatTools) ListMembers(workspaceId, _ db.Id, _ map[string]string) (any, error) {
	return t.Workspace.ListWorkspaceMembers(workspaceId, ChatToolLimit, 0)
}

// ListKeyValue returns the workspace keys, under the prefix argument when given.
func (t *ChatTools) ListKeyValue(workspaceId, _ db.Id, args map[string]string) (any, error) {
	return t.KeyValue.ListKeyValue(workspaceId, args["prefix"])
}

// ListIntegrations returns the workspace integrations without their URL, site or index.
func (t *ChatTools) ListIntegrations(workspaceId, _ db.Id, _ map[string]string) (any, error) {
	result, err := t.Integration.ListIntegrations(workspaceId, ChatToolLimit, 0)
	if err != nil {
		return nil, err
	}

	return lo.Map(result.Integrations, func(item *IntegrationResponse, _ int) *ChatIntegration {
		return &ChatIntegration{
			Name:      item.Name,
			Type:      item.Type,
			Events:    item.Events,
			IsActive:  item.IsActive,
			HasSecret: item.HasSecret,
		}
	}), nil
}

// ListEvents returns the latest workspace events.
func (t *ChatTools) ListEvents(workspaceId, _ db.Id, _ map[string]string) (any, error) {
	return t.Event.ListEvents(workspaceId, ChatToolLimit, 0)
}
