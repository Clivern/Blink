// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationAgentSkillRepository(t *testing.T) {
	database := openTestDB(t)
	workspace := createTestWorkspace(t, database)
	agents := NewAgentRepository(database)
	repo := NewAgentSkillRepository(database)

	createAgent := func(name, card string) *Agent {
		agent := &Agent{WorkspaceId: workspace.Id, Name: name, CardPath: "/.well-known/agent-card.json", Version: "1.0.0"}
		require.NoError(t, agents.Create(agent))

		stored, err := agents.SetCardIfMissing(agent.Id, card, "cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc")
		require.NoError(t, err)
		require.True(t, stored)

		return agent
	}

	createAgent("support-assistant", `{"skills":[{"id":"ticket.open"}]}`)
	directory := createAgent("user-directory", `{
		"defaultInputModes": ["application/json"],
		"skills": [
			{"id": "user.profile.get", "name": "Get user profile", "tags": ["profile", "users"]},
			{"id": "user.delete", "name": "Delete user", "tags": ["users"]}
		]
	}`)
	mirror := createAgent("user-mirror", `{"skills":[{"id":"user.profile.get","name":"Profile (mirror)","tags":["profile"]}]}`)
	createAgent("billing", `{"skills":[{"id":"invoice.get"}]}`)

	filter := func(change func(*ListCallableSkillsFilter)) *ListCallableSkillsFilter {
		f := &ListCallableSkillsFilter{Caller: "support-assistant", Tags: []string{}, Limit: 20}
		change(f)

		return f
	}

	t.Run("lists skills grouped by id, skipping the caller's own offers", func(t *testing.T) {
		skills, total, err := repo.ListCallable(workspace.Id, filter(func(*ListCallableSkillsFilter) {}))
		require.NoError(t, err)
		assert.Equal(t, int64(3), total)
		require.Len(t, skills, 3)
		assert.Equal(t, "user.profile.get", skills[2].SkillId)
		assert.Equal(t, "Get user profile", skills[2].Name)
		assert.Equal(t, []string{"profile", "users"}, skills[2].Tags)
		assert.Equal(t, []string{"application/json"}, skills[2].InputModes)
		assert.Equal(t, []string{}, skills[2].OutputModes)
		assert.Equal(t, 2, skills[2].Agents)
	})

	t.Run("filters by tag", func(t *testing.T) {
		skills, total, err := repo.ListCallable(workspace.Id, filter(func(f *ListCallableSkillsFilter) { f.Tags = []string{"billing"} }))
		require.NoError(t, err)
		assert.Equal(t, int64(0), total)
		assert.Empty(t, skills)
	})

	t.Run("passing only counts agents with a live instance", func(t *testing.T) {
		createTestInstance(t, database, mirror.Id, "mirror-1")
		down := createTestInstance(t, database, directory.Id, "directory-1")
		require.NoError(t, NewHealthCheckRepository(database).Create(&HealthCheck{
			InstanceId: down.Id,
			CheckId:    "http",
			Name:       "HTTP",
			Type:       HealthCheckTypeHTTP,
			Source:     "agent",
			Status:     HealthCheckStatusCritical,
		}))

		skills, _, err := repo.ListCallable(workspace.Id, filter(func(f *ListCallableSkillsFilter) { f.Passing = true }))
		require.NoError(t, err)
		require.Len(t, skills, 1)
		assert.Equal(t, 1, skills[0].Agents)
		assert.Equal(t, "Profile (mirror)", skills[0].Name)
	})

	t.Run("keeps the total when the page is past the end", func(t *testing.T) {
		skills, total, err := repo.ListCallable(workspace.Id, filter(func(f *ListCallableSkillsFilter) { f.Offset = 5 }))
		require.NoError(t, err)
		assert.Equal(t, int64(3), total)
		assert.Empty(t, skills)
	})

	t.Run("get skill describes it from the first agent by name and is nil when nobody offers it", func(t *testing.T) {
		skill, err := repo.GetSkill(workspace.Id, "user.profile.get")
		require.NoError(t, err)
		require.NotNil(t, skill)
		assert.Equal(t, "Get user profile", skill.Name)
		assert.Equal(t, []string{"profile", "users"}, skill.Tags)

		skill, err = repo.GetSkill(workspace.Id, "ghost.skill")
		require.NoError(t, err)
		assert.Nil(t, skill)
	})

	t.Run("list callable agents counts instances and filters passing", func(t *testing.T) {
		agentsFor := func(skill string, passing bool) ([]*CallableAgent, int64) {
			list, total, err := repo.ListCallableAgents(workspace.Id, skill, &ListCallableAgentsFilter{Caller: "support-assistant", Passing: passing, Limit: 20})
			require.NoError(t, err)

			return list, total
		}

		list, total := agentsFor("user.profile.get", false)
		assert.Equal(t, int64(2), total)
		require.Len(t, list, 2)
		assert.Equal(t, "user-directory", list[0].Name)
		assert.Equal(t, 1, list[0].Instances)
		assert.Equal(t, 0, list[0].PassingInstances)
		assert.Equal(t, "user-mirror", list[1].Name)
		assert.Equal(t, 1, list[1].PassingInstances)

		list, total = agentsFor("user.profile.get", true)
		assert.Equal(t, int64(1), total)
		require.Len(t, list, 1)
		assert.Equal(t, "user-mirror", list[0].Name)

		list, total = agentsFor("user.delete", false)
		assert.Equal(t, int64(1), total)
		require.Len(t, list, 1)
		assert.Equal(t, "user-directory", list[0].Name)
	})
}
