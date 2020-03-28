package channel

import (
	"slack/events/base"
)

// A channel was created
// https://api.slack.com/events/channel_created
// Expected scopes: channels:read
type Created struct {
	base.Event
	Channel CreatedChannel
}

type CreatedChannel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created int    `json:"created"`
	Creator string `json:"creator"`
}
