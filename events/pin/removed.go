package pin

import (
	"slack/events/base"
)

// A pin was removed from a channel
// https://api.slack.com/events/pin_removed
// Expected scopes: pins:read
type Removed struct {
	base.Event
	User      string      `json:"user"`
	ChannelID string      `json:"channel_id"`
	HasPins   bool        `json:"has_pins"`
	Item      interface{} `json:"item"` // FIXME
	EventTS   string      `json:"event_ts"`
}
