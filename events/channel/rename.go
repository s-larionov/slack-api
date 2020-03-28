package channel

import (
	"slack/events/base"
)

// A channel was renamed
// https://api.slack.com/events/channel_rename
// Expected scopes: channels:read
type Rename struct {
	base.Event
	Channel RenamedChannel `json:"channel"`
}

type RenamedChannel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created int    `json:"created"`
}
