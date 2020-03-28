package star

import (
	"slack/events/base"
)

// A member has starred an item
// https://api.slack.com/events/star_added
// Expected scopes: stars:read
type Added struct {
	base.Event
	User    string `json:"user"`
	Item    Item   `json:"item"`
	EventTS string `json:"event_ts"`
}
