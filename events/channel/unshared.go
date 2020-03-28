package channel

import (
	"slack/events/base"
)

// A channel has been unshared with an external workspace
// https://api.slack.com/events/channel_unshared
// Expected scopes: channels:read, groups:read
type Unshared struct {
	base.Event
	PreviouslyConnectedTeamID string `json:"previously_connected_team_id"`
	Channel                   string `json:"channel"`
	IsExtShared               bool   `json:"is_ext_shared"`
	EventTS                   string `json:"event_ts"`
}
