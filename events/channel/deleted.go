package channel

import (
	"slack/events/base"
)

// A channel was deleted
// https://api.slack.com/events/channel_deleted
// Expected scopes: channels:read
type Deleted struct {
	base.Event
	Channel string `json:"channel"`
}
