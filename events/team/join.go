package team

import (
	"slack/events/base"
)

// A new member has joined
// https://api.slack.com/events/team_join
// Expected scopes: users:read
type Join struct {
	base.Event
	User interface{} `json:"user"` // TODO
}
