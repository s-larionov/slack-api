package channel

import (
	"slack/events/base"
)

// A channel was archived
// https://api.slack.com/events/channel_archive
// Expected scopes: channels:read
type Archive struct {
	base.Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}
