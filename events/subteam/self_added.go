package subteam

import (
	"slack/events/base"
)

// You have been added to a User Group
// https://api.slack.com/events/subteam_self_added
// Expected scopes: usergroups:read
type SelfAdded struct {
	base.Event
	SubteamID string `json:"subteam_id"`
}
