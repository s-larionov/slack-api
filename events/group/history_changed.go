package group

import (
	"slack/events/base"
)

// Bulk updates were made to a private channel's history
// https://api.slack.com/events/group_history_changed
// Expected scopes: groups:history
type HistoryChanged struct {
	base.Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}
