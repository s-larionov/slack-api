package im

import (
	"slack/events/base"
)

// You opened a DM
// https://api.slack.com/events/im_open
// Expected scopes: im:read
type Open struct {
	base.Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}
