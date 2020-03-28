package group

import (
	"slack/events/base"
)

// A private channel was unarchived
// https://api.slack.com/events/group_unarchive
// Expected scopes: groups:read
type Unarchive struct {
	base.Event
	Channel string `json:"channel"`
}
