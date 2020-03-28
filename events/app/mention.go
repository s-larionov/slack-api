package app

import (
	"slack/events/base"
)

// Subscribe to only the message events that mention your app or bot
// https://api.slack.com/events/app_mention
// Expected scopes: app_mentions:read
type Mention struct {
	base.Event
	User    string `json:"user"`
	Channel string `json:"channel"`
	EventTS int64  `json:"event_ts"`
	Text    string `json:"text"`
	TS      string `json:"ts"`
}
