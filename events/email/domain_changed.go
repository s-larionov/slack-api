package email

import (
	"slack/events/base"
)

// The workspace email domain has changed
// https://api.slack.com/events/email_domain_changed
// Expected scopes: team:read
type DomainChanged struct {
	base.Event
	Domain  string `json:"email_domain"`
	EventTS string `json:"event_ts"`
}
