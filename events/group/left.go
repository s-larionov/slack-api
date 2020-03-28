package group

import (
	"slack/events/base"
)

// You left a private channel
// https://api.slack.com/events/group_left
// Expected scopes: groups:read
type Left struct {
	base.Event
	Channel string `json:"channel"`
}
