package channel

import (
	"slack/events/base"
)

// A channel has been shared with an external workspace
// https://api.slack.com/events/channel_shared
// Expected scopes: channels:read, groups:read
type Shared struct {
	base.Event
	ConnectedTeamID string `json:"connected_team_id"`
	Channel         string `json:"channel"`
	EventTS         string `json:"event_ts"`
}
