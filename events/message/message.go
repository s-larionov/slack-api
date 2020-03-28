package message

import (
	"slack/channel"
	"slack/events/base"
	"slack/message"
)

// A message was sent to a channel
// https://api.slack.com/events/message
// Expected scopes: channels:history, groups:history, im:history, mpim:history
type Message struct {
	base.Event
	Channel     string               `json:"channel"`
	ChannelType channel.Type         `json:"channel_type"`
	User        string               `json:"user"`
	Text        string               `json:"text"`
	TS          string               `json:"ts"`
	SubType     message.SubType      `json:"sub_type"`
	Attachments []message.Attachment `json:"attachments"`
	Edited      message.Edited       `json:"edited"`
	// The is_starred property is present and true if the calling user has starred the message.
	IsStarred bool   `json:"is_starred"`
	IsHidden  bool   `json:"hidden"`
	DeletedTS string `json:"deleted_ts"`
	EventTS   string `json:"event_ts"`
	// Contains the IDs of any channels to which the message is currently pinned.
	PinnedTo  []string   `json:"pinned_to"`
	Reactions []Reaction `json:"reactions"`
}

// FIXME
type Reaction struct {
	Name  string   `json:"name"`
	Count int64    `json:"count"`
	Users []string `json:"users"`
}
