package dnd

import (
	"slack/events/base"
)

// Do not Disturb settings changed for a member
// https://api.slack.com/events/dnd_updated_user
// Expected scopes: dnd:read
type UpdatedUser struct {
	base.Event
	User   string
	Status UpdatedUserStatus `json:"dnd_status"`
}

type UpdatedUserStatus struct {
	IsEnabled   bool  `json:"dnd_enabled"`
	NextStartTS int64 `json:"next_dnd_start_ts"`
	NextEndTS   int64 `json:"next_dnd_end_ts"`
}
