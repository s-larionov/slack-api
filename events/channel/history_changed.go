package channel

import (
	"slack/events/base"
)

// Bulk updates were made to a channel's history
// https://api.slack.com/events/channel_history_changed
// Expected scopes: channels:history, groups:history, mpim:history
type HistoryChanged struct {
	base.Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}
