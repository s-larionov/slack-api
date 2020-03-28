package group

import (
	"slack/events/base"
)

// A private channel was deleted
// https://api.slack.com/events/group_deleted
// Expected scopes: groups:read
type Deleted struct {
	base.Event
	Channel string `json:"channel"`
}
