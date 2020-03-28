package subteam

import (
	"slack/events/base"
	"slack/subteam"
)

// A User Group has been added to the workspace
// https://api.slack.com/events/subteam_created
// Expected scopes: usergroups:read
type Created struct {
	base.Event
	Subteam subteam.Subteam `json:"subteam"`
}
