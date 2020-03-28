package star

import (
	"slack/events/base"
)

// A member removed a star
// https://api.slack.com/events/star_removed
// Expected scopes: stars:read
type Removed struct {
	base.Event
	User    string `json:"user"`
	Item    Item   `json:"item"`
	EventTS string `json:"event_ts"`
}
