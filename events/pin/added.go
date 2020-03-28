package pin

import (
	"slack/events/base"
)

// A pin was added to a channel
// https://api.slack.com/events/pin_added
// Expected scopes: pins:read
type Added struct {
	base.Event
	User      string      `json:"user"`
	ChannelID string      `json:"channel_id"`
	Item      interface{} `json:"item"` // FIXME
	EventTS   string      `json:"event_ts"`
}
