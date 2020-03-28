package invite

import (
	"slack/events/base"
)

// User requested an invite
// https://api.slack.com/events/invite_requested
// Expected scopes: admin.invites:read
type Requested struct {
	base.Event
	Invite Request `json:"invite_request"`
}

type Request struct {
	ID            string      `json:"id"`
	Email         string      `json:"email"`
	DateCreated   int64       `json:"date_created"`
	RequesterIDs  []string    `json:"requester_ids"`
	ChannelIDs    []string    `json:"channel_ids"`
	Type          string      `json:"invite_type"`
	RealName      string      `json:"real_name"`
	DateExpire    int64       `json:"date_expire"`
	RequestReason string      `json:"request_reason"`
	Team          RequestTeam `json:"team"`
}

type RequestTeam struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}
