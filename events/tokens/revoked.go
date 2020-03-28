package tokens

import (
	"slack/events/base"
)

// API tokens for your app were revoked.
// https://api.slack.com/events/tokens_revoked
// Expected scopes: none
type Revoked struct {
	base.Event
	Tokens RevokedTokens `json:"tokens"`
}

type RevokedTokens struct {
	OAuth []string `json:"oauth"`
	Bot   []string `json:"bot"`
}
