package emoji

import (
	"slack/events/base"
)

const (
	ChangedSubTypeAdd    = "add"
	ChangedSubTypeRemove = "remove"
)

// A custom emoji has been added or changed
// https://api.slack.com/events/emoji_changed
// Expected scopes: emoji:read
type Changed struct {
	base.Event
	SubType ChangedSubType `json:"subtype"`
	EventTS string         `json:"event_ts"`
}

type ChangedSubType string

type ChangedAdd struct {
	Changed
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ChangedRemove struct {
	Changed
	Names []string `json:"names"`
}
