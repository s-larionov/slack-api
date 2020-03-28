package app

import (
	"slack/events/base"
)

// Your Slack app was uninstalled.
// https://api.slack.com/events/app_uninstalled
// Expected scopes: none
type Uninstalled struct {
	base.Event
}
