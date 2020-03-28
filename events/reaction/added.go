package reaction

import (
	"slack/events/base"
)

// A member has added an emoji reaction to an item
// https://api.slack.com/events/reaction_added
// Expected scopes: reactions:read
type Added struct {
	base.Event
	User     string `json:"user"`
	Reaction string `json:"reaction"`
	ItemUser string `json:"item_user"`
	Item     Item   `json:"item"`
	EventTS  string `json:"event_ts"`
}
