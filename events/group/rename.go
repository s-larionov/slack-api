package group

import (
	"slack/events/base"
)

// A private channel was renamed
// https://api.slack.com/events/group_rename
// Expected scopes: groups:read
type Rename struct {
	base.Event
	Channel RenameChannel `json:"channel"`
}

type RenameChannel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created int64  `json:"created"`
}
