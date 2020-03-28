package group

import (
	"slack/events/base"
)

// A private channel was archived
// https://api.slack.com/events/group_archive
// Expected scopes: groups:read
type Archive struct {
	base.Event
	Channel string `json:"channel"`
}
