package link

import (
	"slack/events/base"
)

// A message was posted containing one or more links relevant to your application
// https://api.slack.com/events/link_shared
// Expected scopes: links:read
type Shared struct {
	base.Event
	Channel   string       `json:"channel"`
	User      string       `json:"user"`
	MessageTS string       `json:"message_ts"`
	ThreadTS  string       `json:"thread_ts"`
	Links     []SharedLink `json:"links"`
}

type SharedLink struct {
	Domain string `json:"domain"`
	URL    string `json:"url"`
}
