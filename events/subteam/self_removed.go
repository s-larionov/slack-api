package subteam

import (
	"slack/events/base"
)

// You have been removed from a User Group
// https://api.slack.com/events/subteam_self_removed
// Expected scopes: usergroups:read
type SelfRemoved struct {
	base.Event
	SubteamID string `json:"subteam_id"`
}
