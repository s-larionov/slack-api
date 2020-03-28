package events

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fastjson"

	"slack/events/app"
	"slack/events/base"
	"slack/events/call"
	"slack/events/channel"
	"slack/events/dnd"
	"slack/events/email"
	"slack/events/emoji"
	"slack/events/file"
	"slack/events/grid"
	"slack/events/group"
	"slack/events/im"
	"slack/events/invite"
	"slack/events/link"
	"slack/events/member"
	"slack/events/message"
	"slack/events/pin"
	"slack/events/reaction"
	"slack/events/star"
	"slack/events/subteam"
	"slack/events/team"
	"slack/events/tokens"
)

var events = map[base.Type]func() interface{}{
	base.TypeAppHomeOpened:         func() interface{} { return &app.HomeOpened{} },
	base.TypeAppMention:            func() interface{} { return &app.Mention{} },
	base.TypeAppRequested:          func() interface{} { return &app.Requested{} },
	base.TypeAppUninstalled:        func() interface{} { return &app.Uninstalled{} },
	base.TypeCallRejected:          func() interface{} { return &call.Rejected{} },
	base.TypeChannelArchive:        func() interface{} { return &channel.Archive{} },
	base.TypeChannelCreated:        func() interface{} { return &channel.Created{} },
	base.TypeChannelDeleted:        func() interface{} { return &channel.Deleted{} },
	base.TypeChannelHistoryChanged: func() interface{} { return &channel.HistoryChanged{} },
	base.TypeChannelLeft:           func() interface{} { return &channel.Left{} },
	base.TypeChannelShared:         func() interface{} { return &channel.Shared{} },
	base.TypeChannelUnarchive:      func() interface{} { return &channel.Unarchive{} },
	base.TypeChannelUnshared:       func() interface{} { return &channel.Unshared{} },
	base.TypeDndUpdated:            func() interface{} { return &dnd.Updated{} },
	base.TypeDndUpdatedUser:        func() interface{} { return &dnd.UpdatedUser{} },
	base.TypeEmailDomainChanged:    func() interface{} { return &email.DomainChanged{} },
	base.TypeEmojiChanged:          func() interface{} { return &emoji.Changed{} },
	base.TypeFileChange:            func() interface{} { return &file.Change{} },
	base.TypeFileCreated:           func() interface{} { return &file.Created{} },
	base.TypeFileDeleted:           func() interface{} { return &file.Deleted{} },
	base.TypeFilePublic:            func() interface{} { return &file.Public{} },
	base.TypeFileShared:            func() interface{} { return &file.Shared{} },
	base.TypeFileUnshared:          func() interface{} { return &file.Unshared{} },
	base.TypeGridMigrationFinished: func() interface{} { return &grid.MigrationFinished{} },
	base.TypeGridMigrationStarted:  func() interface{} { return &grid.MigrationStarted{} },
	base.TypeGroupArchive:          func() interface{} { return &group.Archive{} },
	base.TypeGroupClose:            func() interface{} { return &group.Close{} },
	base.TypeGroupDeleted:          func() interface{} { return &group.Deleted{} },
	base.TypeGroupHistoryChanged:   func() interface{} { return &group.HistoryChanged{} },
	base.TypeGroupLeft:             func() interface{} { return &group.Left{} },
	base.TypeGroupOpen:             func() interface{} { return &group.Open{} },
	base.TypeGroupRename:           func() interface{} { return &group.Rename{} },
	base.TypeGroupUnarchive:        func() interface{} { return &group.Unarchive{} },
	base.TypeIMClose:               func() interface{} { return &im.Close{} },
	base.TypeIMCreated:             func() interface{} { return &im.Created{} },
	base.TypeIMHistoryChanged:      func() interface{} { return &im.HistoryChanged{} },
	base.TypeIMOpen:                func() interface{} { return &im.Open{} },
	base.TypeInviteRequested:       func() interface{} { return &invite.Requested{} },
	base.TypeLinkShared:            func() interface{} { return &link.Shared{} },
	base.TypeMemberJoinedChannel:   func() interface{} { return &member.JoinedChannel{} },
	base.TypeMemberLeftChannel:     func() interface{} { return &member.LeftChannel{} },
	base.TypeMessage:               func() interface{} { return &message.Message{} },
	base.TypePinAdded:              func() interface{} { return &pin.Added{} },
	base.TypePinRemoved:            func() interface{} { return &pin.Removed{} },
	base.TypeReactionAdded:         func() interface{} { return &reaction.Added{} },
	base.TypeReactionRemoved:       func() interface{} { return &reaction.Removed{} },
	base.TypeStarAdded:             func() interface{} { return &star.Added{} },
	base.TypeStarRemoved:           func() interface{} { return &star.Removed{} },
	base.TypeSubteamCreated:        func() interface{} { return &subteam.Created{} },
	base.TypeSubteamMembersChanged: func() interface{} { return &subteam.MembersChanged{} },
	base.TypeSubteamSelfAdded:      func() interface{} { return &subteam.SelfAdded{} },
	base.TypeSubteamSelfRemoved:    func() interface{} { return &subteam.SelfRemoved{} },
	base.TypeSubteamUpdated:        func() interface{} { return &subteam.Updated{} },
	base.TypeTeamDomainChange:      func() interface{} { return &team.DomainChange{} },
	base.TypeTeamJoin:              func() interface{} { return &team.Join{} },
	base.TypeTeamRename:            func() interface{} { return &team.Rename{} },
	base.TypeTokensRevoked:         func() interface{} { return &tokens.Revoked{} },
}

var (
	ErrUnknownEvent = errors.New("unknown event type")
)

type Wrapper struct {
	Type        RequestType `json:"type"`
	Token       string      `json:"token"`
	TeamID      string      `json:"team_id"`
	APIAppID    string      `json:"api_app_id"`
	Event       interface{} `json:"-"` // FIXME
	EventID     string      `json:"event_id"`
	EventTime   int64       `json:"event_time"`
	AuthedUsers []string    `json:"authed_users"`
}

type wrapper Wrapper

func (w *Wrapper) UnmarshalJSON(body []byte) error {
	var wr wrapper
	err := json.Unmarshal(body, &wr)
	if err != nil {
		return err
	}

	w.Type = wr.Type
	w.Token = wr.Token
	w.TeamID = wr.TeamID
	w.APIAppID = wr.APIAppID
	w.EventID = wr.EventID
	w.EventTime = wr.EventTime
	w.AuthedUsers = wr.AuthedUsers

	// FIXME Might be we shouldn't use fastjson
	parsed, err := fastjson.Parse(string(body))
	if err != nil {
		return err
	}
	if !parsed.Exists("event") {
		return errors.New("unable to parse event field")
	}

	var eventBody []byte
	eventBody = parsed.Get("event").MarshalTo(eventBody)

	eventType := base.Type(fastjson.GetString(eventBody, "type"))
	if len(eventType) == 0 {
		return errors.New("unknown event type")
	}

	generator, ok := events[eventType]
	if !ok {
		return ErrUnknownEvent
	}
	event := generator()
	err = json.Unmarshal(eventBody, &event)
	if err != nil {
		return err
	}
	w.Event = event

	return nil
}
