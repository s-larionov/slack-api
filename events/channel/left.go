package channel

import (
	"slack/events/base"
)

// You left a channel
// https://api.slack.com/events/channel_left
// Expected scopes: channels:read
type Left struct {
	base.Event
	Channel string `json:"channel"`
}
