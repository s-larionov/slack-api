package subteam

import (
	"slack/events/base"
	"slack/subteam"
)

// An existing User Group has been updated or its members changed
// https://api.slack.com/events/subteam_updated
// Expected scopes: usergroups:read
type Updated struct {
	base.Event
	Subteam subteam.Subteam `json:"subteam"`
}
