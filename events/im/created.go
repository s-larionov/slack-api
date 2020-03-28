package im

import (
	"slack/events/base"
)

// A DM was created
// https://api.slack.com/events/im_created
// Expected scopes: im:read
type Created struct {
	base.Event
	User    string         `json:"user"`
	Channel CreatedChannel `json:"channel"`
}

// TODO
type CreatedChannel struct {
}
