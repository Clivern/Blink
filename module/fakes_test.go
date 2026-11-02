// Copyright 2026 Blink. All rights reserved.
// License can be found in the LICENSE file.

package module

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/clivern/blink/db"
	"github.com/clivern/blink/pkg/ai"

	"github.com/samber/lo"
)

type fakeWorkspaces struct {
	db.WorkspaceRepository
	workspace  *db.Workspace
	list       []*db.Workspace
	total      int64
	membership *db.WorkspaceUser
	created    []*db.Workspace
	updated    []*db.Workspace
	deleted    []db.Id
}

func (f *fakeWorkspaces) GetById(id db.Id) (*db.Workspace, error) {
	return f.workspace, nil
}

func (f *fakeWorkspaces) GetByHandle(handle string) (*db.Workspace, error) {
	return nil, nil
}

func (f *fakeWorkspaces) Create(workspace *db.Workspace) error {
	workspace.Id = "ws-new"
	f.created = append(f.created, workspace)

	return nil
}

func (f *fakeWorkspaces) Update(workspace *db.Workspace) error {
	f.updated = append(f.updated, workspace)

	return nil
}

func (f *fakeWorkspaces) Delete(id db.Id) error {
	f.deleted = append(f.deleted, id)

	return nil
}

func (f *fakeWorkspaces) List(limit, offset int, userId db.Id) ([]*db.Workspace, error) {
	return f.list, nil
}

func (f *fakeWorkspaces) Count(userId db.Id) (int64, error) {
	return f.total, nil
}

func (f *fakeWorkspaces) GetWorkspaceMembership(workspaceId, userId db.Id) (*db.WorkspaceUser, error) {
	return f.membership, nil
}

// errTest is the error fakes return when a test wants a repository failure.
var errTest = errors.New("boom")

type fakeConfigs struct {
	db.ConfigRepository
	values  map[string]string
	created []string
	updated []string
	err     error
}

func (f *fakeConfigs) Get(key string) (*db.Config, error) {
	value, ok := f.values[key]
	if f.err != nil || !ok {
		return nil, f.err
	}

	return &db.Config{Key: key, Value: value}, nil
}

func (f *fakeConfigs) Create(key, value string) error {
	if f.values == nil {
		f.values = map[string]string{}
	}
	f.values[key] = value
	f.created = append(f.created, key)

	return f.err
}

func (f *fakeConfigs) Update(key, value string) error {
	if f.values == nil {
		f.values = map[string]string{}
	}
	f.values[key] = value
	f.updated = append(f.updated, key)

	return f.err
}

type fakeUsers struct {
	db.UserRepository
	user       *db.User
	byEmail    *db.User
	byProvider *db.User
	created    []*db.User
	updated    []*db.User
	lastLogin  []db.Id
	err        error
}

func (f *fakeUsers) GetById(id db.Id) (*db.User, error) {
	return f.user, f.err
}

func (f *fakeUsers) GetByEmail(email string) (*db.User, error) {
	return f.byEmail, f.err
}

func (f *fakeUsers) GetByProvider(provider, providerUserId string) (*db.User, error) {
	return f.byProvider, f.err
}

func (f *fakeUsers) Create(user *db.User) error {
	if user.Id == "" {
		user.Id = "user-new"
	}
	f.created = append(f.created, user)

	return f.err
}

func (f *fakeUsers) Update(user *db.User) error {
	f.updated = append(f.updated, user)

	return f.err
}

func (f *fakeUsers) UpdateLastLogin(id db.Id) error {
	f.lastLogin = append(f.lastLogin, id)

	return f.err
}

type fakeKeyValues struct {
	db.KeyValueRepository
	item     *db.KeyValue
	upserted []*db.KeyValue
	deleted  []string
	expired  int64
	err      error
}

func (f *fakeKeyValues) Get(key string) (*db.KeyValue, error) {
	return f.item, f.err
}

func (f *fakeKeyValues) Upsert(item *db.KeyValue) error {
	f.upserted = append(f.upserted, item)

	return f.err
}

func (f *fakeKeyValues) Delete(key string) error {
	f.deleted = append(f.deleted, key)

	return f.err
}

func (f *fakeKeyValues) DeleteExpired() (int64, error) {
	return f.expired, f.err
}

type fakeWorkspaceStats struct {
	db.WorkspaceStatsRepository
	err error
}

func (f *fakeWorkspaceStats) GetByWorkspaceId(workspaceId db.Id) (*db.WorkspaceStats, error) {
	return &db.WorkspaceStats{}, f.err
}

type fakeWorkspaceKeyValues struct {
	db.WorkspaceKeyValueRepository
	item     *db.WorkspaceKeyValue
	items    []*db.WorkspaceKeyValue
	upserted []*db.WorkspaceKeyValue
	deleted  []string
	prefixes []string
	err      error
}

func (f *fakeWorkspaceKeyValues) Get(workspaceId db.Id, key string) (*db.WorkspaceKeyValue, error) {
	return f.item, f.err
}

func (f *fakeWorkspaceKeyValues) Upsert(item *db.WorkspaceKeyValue) error {
	f.upserted = append(f.upserted, item)

	return f.err
}

func (f *fakeWorkspaceKeyValues) Delete(workspaceId db.Id, key string) error {
	f.deleted = append(f.deleted, key)

	return f.err
}

func (f *fakeWorkspaceKeyValues) ListByPrefix(workspaceId db.Id, prefix string) ([]*db.WorkspaceKeyValue, error) {
	f.prefixes = append(f.prefixes, prefix)

	return f.items, f.err
}

type fakeAPIKeys struct {
	db.APIKeyRepository
	item    *db.APIKey
	list    []*db.APIKey
	total   int64
	expired int64
	created []*db.APIKey
	deleted []db.Id
	err     error
}

func (f *fakeAPIKeys) GetById(id db.Id) (*db.APIKey, error) {
	return f.item, f.err
}

func (f *fakeAPIKeys) GetByKey(key string) (*db.APIKey, error) {
	return f.item, f.err
}

func (f *fakeAPIKeys) Create(apiKey *db.APIKey) error {
	apiKey.Id = "key-new"
	f.created = append(f.created, apiKey)

	return f.err
}

func (f *fakeAPIKeys) Delete(id db.Id) error {
	f.deleted = append(f.deleted, id)

	return f.err
}

func (f *fakeAPIKeys) DeleteExpired() (int64, error) {
	return f.expired, f.err
}

func (f *fakeAPIKeys) ListByUserId(userId db.Id, limit, offset int) ([]*db.APIKey, error) {
	return f.list, f.err
}

func (f *fakeAPIKeys) CountByUserId(userId db.Id) (int64, error) {
	return f.total, f.err
}

type fakeAccessKeys struct {
	db.AccessKeyRepository
	item    *db.AccessKey
	list    []*db.AccessKey
	total   int64
	expired int64
	created []*db.AccessKey
	deleted []db.Id
	err     error
}

func (f *fakeAccessKeys) GetById(id db.Id) (*db.AccessKey, error) {
	return f.item, f.err
}

func (f *fakeAccessKeys) GetByKey(key string) (*db.AccessKey, error) {
	return f.item, f.err
}

func (f *fakeAccessKeys) Create(accessKey *db.AccessKey) error {
	accessKey.Id = "access-new"
	f.created = append(f.created, accessKey)

	return f.err
}

func (f *fakeAccessKeys) Delete(id db.Id) error {
	f.deleted = append(f.deleted, id)

	return f.err
}

func (f *fakeAccessKeys) DeleteExpired() (int64, error) {
	return f.expired, f.err
}

func (f *fakeAccessKeys) ListByWorkspaceId(workspaceId db.Id, limit, offset int) ([]*db.AccessKey, error) {
	return f.list, f.err
}

func (f *fakeAccessKeys) CountByWorkspaceId(workspaceId db.Id) (int64, error) {
	return f.total, f.err
}

type fakeWorkspaceUsers struct {
	db.WorkspaceUserRepository
	membership *db.WorkspaceUser
	list       []*db.WorkspaceUser
	total      int64
	created    []*db.WorkspaceUser
	updated    []*db.WorkspaceUser
	deleted    []db.Id
	err        error
}

func (f *fakeWorkspaceUsers) GetByWorkspaceAndUser(workspaceId, userId db.Id) (*db.WorkspaceUser, error) {
	return f.membership, f.err
}

func (f *fakeWorkspaceUsers) Create(member *db.WorkspaceUser) error {
	f.created = append(f.created, member)

	return f.err
}

func (f *fakeWorkspaceUsers) Update(member *db.WorkspaceUser) error {
	f.updated = append(f.updated, member)

	return f.err
}

func (f *fakeWorkspaceUsers) Delete(id db.Id) error {
	f.deleted = append(f.deleted, id)

	return f.err
}

func (f *fakeWorkspaceUsers) ListByWorkspaceId(workspaceId db.Id, limit, offset int) ([]*db.WorkspaceUser, error) {
	return f.list, f.err
}

func (f *fakeWorkspaceUsers) CountByWorkspaceId(workspaceId db.Id) (int64, error) {
	return f.total, f.err
}

type fakeIntegrations struct {
	db.IntegrationRepository
	item    *db.Integration
	list    []*db.Integration
	total   int64
	created []*db.Integration
	updated []*db.Integration
	deleted []db.Id
	err     error
}

func (f *fakeIntegrations) GetById(id db.Id) (*db.Integration, error) {
	return f.item, f.err
}

func (f *fakeIntegrations) Create(integration *db.Integration) error {
	integration.Id = "integration-new"
	f.created = append(f.created, integration)

	return f.err
}

func (f *fakeIntegrations) Update(integration *db.Integration) error {
	f.updated = append(f.updated, integration)

	return f.err
}

func (f *fakeIntegrations) Delete(id db.Id) error {
	f.deleted = append(f.deleted, id)

	return f.err
}

func (f *fakeIntegrations) ListByWorkspaceId(workspaceId db.Id, limit, offset int) ([]*db.Integration, error) {
	return f.list, f.err
}

func (f *fakeIntegrations) CountByWorkspaceId(workspaceId db.Id) (int64, error) {
	return f.total, f.err
}

type fakeSessions struct {
	db.SessionRepository
	session     *db.Session
	list        []*db.Session
	created     []*db.Session
	deleted     []db.Id
	deletedUser []db.Id
	expired     int64
	err         error
}

func (f *fakeSessions) Create(session *db.Session) error {
	f.created = append(f.created, session)

	return f.err
}

func (f *fakeSessions) GetByToken(token string) (*db.Session, error) {
	return f.session, f.err
}

func (f *fakeSessions) Delete(id db.Id) error {
	f.deleted = append(f.deleted, id)

	return f.err
}

func (f *fakeSessions) DeleteByUserId(userId db.Id) error {
	f.deletedUser = append(f.deletedUser, userId)

	return f.err
}

func (f *fakeSessions) GetByUserId(userId db.Id) ([]*db.Session, error) {
	return f.list, f.err
}

func (f *fakeSessions) DeleteExpired() (int64, error) {
	return f.expired, f.err
}

type fakeUserInvites struct {
	db.UserInviteRepository
	item          *db.UserInvite
	list          []*db.UserInvite
	pending       []*db.UserInvite
	total         int64
	pendingCount  int64
	expired       int64
	created       []*db.UserInvite
	deleted       []db.Id
	statusUpdates []string
	err           error
}

func (f *fakeUserInvites) GetById(id db.Id) (*db.UserInvite, error) {
	return f.item, f.err
}

func (f *fakeUserInvites) Create(invite *db.UserInvite) error {
	invite.Id = "invite-new"
	f.created = append(f.created, invite)

	return f.err
}

func (f *fakeUserInvites) Delete(id db.Id) error {
	f.deleted = append(f.deleted, id)

	return f.err
}

func (f *fakeUserInvites) ListByWorkspaceId(workspaceId db.Id, limit, offset int) ([]*db.UserInvite, error) {
	return f.list, f.err
}

func (f *fakeUserInvites) CountByWorkspaceId(workspaceId db.Id) (int64, error) {
	return f.total, f.err
}

func (f *fakeUserInvites) ListPendingByEmail(email string) ([]*db.UserInvite, error) {
	return f.pending, f.err
}

func (f *fakeUserInvites) UpdateStatus(id db.Id, status string, acceptedAt *time.Time) error {
	f.statusUpdates = append(f.statusUpdates, status)

	return f.err
}

func (f *fakeUserInvites) MarkExpiredAsExpired() (int64, error) {
	return f.expired, f.err
}

func (f *fakeUserInvites) CountPendingByEmailInWorkspace(workspaceId db.Id, email string) (int64, error) {
	return f.pendingCount, f.err
}

type fakeMailer struct {
	sent []string
	err  error
}

func (f *fakeMailer) SendInviteEmail(to, inviteLink, pname string) error {
	f.sent = append(f.sent, to)

	return f.err
}

func newWorkspaceModule(
	workspaces *fakeWorkspaces,
	members *fakeWorkspaceUsers,
) *Workspace {
	return NewWorkspace(workspaces, members, &fakeUsers{})
}

func googleIdentity() *OAuthIdentity {
	return &OAuthIdentity{
		Provider:       "google",
		ProviderUserID: "g-123",
		Email:          "ada@example.com",
		Name:           "Ada",
	}
}

// newOAuthAuth wires an Auth module for OAuth logins.
func newOAuthAuth(users *fakeUsers, sessions *fakeSessions, configs *fakeConfigs) *Auth {
	return NewAuth(users, sessions, configs)
}

// newPerm builds a checker without a database, with an in-memory workspace repository reporting membership.
func newPerm(membership *db.WorkspaceUser) *Perm {
	perm := NewPerm(nil)
	perm.WorkspaceRepository = &fakeWorkspaces{
		membership: membership,
		workspace:  &db.Workspace{Id: "ws-1"},
	}

	return perm
}

func newInviteModule(
	invites *fakeUserInvites,
	users *fakeUsers,
	members *fakeWorkspaceUsers,
	mailer *fakeMailer,
) *Invite {
	return NewInvite(
		invites,
		users,
		&fakeConfigs{},
		&fakeWorkspaces{workspace: &db.Workspace{Id: "ws-1"}},
		members,
		mailer,
	)
}

type fakeNotifications struct {
	db.NotificationRepository
	list      []*db.Notification
	unread    int64
	created   []*db.Notification
	read      []db.Id
	dismissed []db.Id
	admin     bool
	hidden    bool
	err       error
}

func (f *fakeNotifications) Create(notification *db.Notification) error {
	notification.Id = db.Id(fmt.Sprintf("notification-%d", len(f.created)+1))
	f.created = append(f.created, notification)
	return f.err
}

func (f *fakeNotifications) GetLatestMatching(match *db.Notification) (*db.Notification, error) {
	for i := len(f.created) - 1; i >= 0; i-- {
		item := f.created[i]
		if item.Audience == match.Audience &&
			lo.FromPtr(item.WorkspaceId) == lo.FromPtr(match.WorkspaceId) &&
			lo.FromPtr(item.UserId) == lo.FromPtr(match.UserId) &&
			item.Type == match.Type &&
			lo.FromPtr(item.ResourceType) == lo.FromPtr(match.ResourceType) &&
			lo.FromPtr(item.ResourceId) == lo.FromPtr(match.ResourceId) {
			return item, f.err
		}
	}

	return nil, f.err
}

func (f *fakeNotifications) ListForUser(workspaceId, userId db.Id, admin bool, limit, offset int) ([]*db.Notification, error) {
	f.admin = admin
	return f.list, f.err
}

func (f *fakeNotifications) CountUnreadForUser(workspaceId, userId db.Id, admin bool) (int64, error) {
	f.admin = admin
	return f.unread, f.err
}

func (f *fakeNotifications) MarkRead(id, workspaceId, userId db.Id, admin bool) (bool, error) {
	f.admin = admin
	if f.hidden {
		return false, f.err
	}

	f.read = append(f.read, id)
	return true, f.err
}

func (f *fakeNotifications) MarkAllRead(workspaceId, userId db.Id, admin bool) error {
	f.admin = admin
	return f.err
}

func (f *fakeNotifications) Dismiss(id, workspaceId, userId db.Id, admin bool) (bool, error) {
	f.admin = admin
	if f.hidden {
		return false, f.err
	}

	f.dismissed = append(f.dismissed, id)
	return true, f.err
}

// recordingTrigger captures the events a module triggers.
type recordingTrigger struct {
	types    []string
	requests []EventRequest
}

func (f *recordingTrigger) Trigger(workspaceId db.Id, event string, request EventRequest) {
	f.types = append(f.types, event)
	f.requests = append(f.requests, request)
}

type fakeEvents struct {
	db.EventRepository
	item  *db.Event
	list  []*db.Event
	total int64
	err   error
}

func (f *fakeEvents) GetById(id db.Id) (*db.Event, error) {
	return f.item, f.err
}

func (f *fakeEvents) ListByWorkspaceId(workspaceId db.Id, limit, offset int) ([]*db.Event, error) {
	return f.list, f.err
}

func (f *fakeEvents) CountByWorkspaceId(workspaceId db.Id) (int64, error) {
	return f.total, f.err
}

type fakeChats struct {
	db.ChatRepository
	chat     *db.Chat
	messages []*db.ChatMessage
}

func (f *fakeChats) GetChat(workspaceId, userId, id db.Id) (*db.Chat, error) {
	return f.chat, nil
}

func (f *fakeChats) UpdateChat(chat *db.Chat) error {
	f.chat = chat
	return nil
}

func (f *fakeChats) DeleteChat(workspaceId, userId, id db.Id) (bool, error) {
	deleted := f.chat != nil
	f.chat = nil
	return deleted, nil
}

func (f *fakeChats) CreateMessage(message *db.ChatMessage) error {
	message.Id = db.Id(fmt.Sprintf("message-%d", len(f.messages)+1))
	f.messages = append(f.messages, message)
	return nil
}

func (f *fakeChats) ListMessages(chatId db.Id) ([]*db.ChatMessage, error) {
	return f.messages, nil
}

// fakeModel replies with the next scripted reply and records what it was asked.
type fakeModel struct {
	replies []string
	asked   [][]ai.Message
}

func (f *fakeModel) Complete(ctx context.Context, messages []ai.Message) (string, ai.Usage, error) {
	f.asked = append(f.asked, messages)
	reply := f.replies[0]
	f.replies = f.replies[1:]
	return reply, ai.Usage{}, nil
}

// Stream replies like Complete, one word at a time.
func (f *fakeModel) Stream(ctx context.Context, messages []ai.Message, onDelta func(string)) (string, ai.Usage, error) {
	reply, usage, err := f.Complete(ctx, messages)
	for _, word := range strings.SplitAfter(reply, " ") {
		onDelta(word)
	}
	return reply, usage, err
}
