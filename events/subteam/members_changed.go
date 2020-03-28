package subteam

import (
	"slack/events/base"
)

// The membership of an existing User Group has changed
// https://api.slack.com/events/subteam_members_changed
// Expected scopes: usergroups:read
type MembersChanged struct {
	base.Event
	SubteamID          string   `json:"subteam_id"`
	TeamID             string   `json:"team_id"`
	DatePreviousUpdate int64    `json:"date_previous_update"`
	DateUpdate         int64    `json:"date_update"`
	AddedUsers         []string `json:"added_users"`
	AddedUsersCount    int32    `json:"added_users_count"`
	RemovedUsers       []string `json:"removed_users"`
	RemovedUsersCount  int32    `json:"removed_users_count"`
}
