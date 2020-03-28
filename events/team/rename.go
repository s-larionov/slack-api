package team

import (
	"slack/events/base"
)

// The workspace name has changed
// https://api.slack.com/events/team_rename
// Expected scopes: team:read
type Rename struct {
	base.Event
	Name string `json:"name"`
}
