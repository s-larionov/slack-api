package member

import (
	"slack/channel"
	"slack/events/base"
)

// A user left a public or private channel
// https://api.slack.com/events/member_left_channel
// Expected scopes: channels:read, groups:read
type LeftChannel struct {
	base.Event
	User        string       `json:"user"`
	Channel     string       `json:"channel"`
	ChannelType channel.Type `json:"channel_type"`
	Team        string       `json:"team"`
}
