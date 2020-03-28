package channel

import (
	"slack/events/base"
)

// A channel was unarchived
// https://api.slack.com/events/channel_unarchive
// Expected scopes: channels:read
type Unarchive struct {
	base.Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}
