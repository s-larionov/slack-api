package im

import (
	"slack/events/base"
)

// You closed a DM
// https://api.slack.com/events/im_close
// Expected scopes: im:read
type Close struct {
	base.Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}
