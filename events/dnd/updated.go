package dnd

import (
	"slack/events/base"
)

// Do not Disturb settings changed for the current user
// https://api.slack.com/events/dnd_updated
// Expected scopes: dnd:read
type Updated struct {
	base.Event
	User   string
	Status UpdatedStatus `json:"dnd_status"`
}

type UpdatedStatus struct {
	IsEnabled       bool  `json:"dnd_enabled"`
	NextStartTS     int64 `json:"next_dnd_start_ts"`
	NextEndTS       int64 `json:"next_dnd_end_ts"`
	IsSnoozeEnabled bool  `json:"snooze_enabled"`
	SnoozeEndtime   int64 `json:"snooze_endtime"`
}
