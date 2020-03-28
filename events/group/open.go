package group

import (
	"slack/events/base"
)

// You created a group DM
// https://api.slack.com/events/group_open
// Expected scopes: groups:read
type Open struct {
	base.Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}
