package user

import (
	"slack/events/base"
)

// A member's data has changed
// https://api.slack.com/events/user_change
// Expected scopes: users:read
type Change struct {
	base.Event
	User interface{} `json:"user"` // FIXME
}
