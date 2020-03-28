package im

import (
	"slack/events/base"
)

// Bulk updates were made to a DM's history
// https://api.slack.com/events/im_history_changed
// Expected scopes: im:history
type HistoryChanged struct {
	base.Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}
