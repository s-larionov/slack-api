package reaction

import (
	"slack/events/base"
)

// A member removed an emoji reaction
// https://api.slack.com/events/reaction_removed
// Expected scopes: reactions:read
type Removed struct {
	base.Event
	User     string `json:"user"`
	Reaction string `json:"reaction"`
	ItemUser string `json:"item_user"`
	Item     Item   `json:"item"`
	EventTS  string `json:"event_ts"`
}
