package team

import (
	"slack/events/base"
)

// The workspace domain has changed
// https://api.slack.com/events/team_domain_change
// Expected scopes: team:read
type DomainChange struct {
	base.Event
	URL    string `json:"url"`
	Domain string `json:"domain"`
}
