package group

import (
	"slack/events/base"
)

// You closed a private channel
// https://api.slack.com/events/group_close
// Expected scopes: groups:read
type Close struct {
	base.Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}
