package app

import (
	"slack/events/base"
)

// User clicked into your App Home
// https://api.slack.com/events/app_home_opened
// Expected scopes: none
type HomeOpened struct {
	base.Event
	User    string `json:"user"`
	Channel string `json:"channel"`
	EventTS int64  `json:"event_ts"`
	Tab     string `json:"tab"`
	View    View   `json:"view"`
}

type View struct {
	ID                 string `json:"id"`
	TeamID             string `json:"team_id"`
	Type               string `json:"type"`
	PrivateMetadata    string `json:"private_metadata"`
	CallbackID         string `json:"callback_id"`
	Hash               string `json:"hash"`
	ClearOnClose       bool   `json:"clear_on_close"`
	NotifyOnClose      bool   `json:"notify_on_close"`
	RootViewID         string `json:"root_view_id"`
	AppID              string `json:"app_id"`
	ExternalID         string `json:"external_id"`
	AppInstalledTeamID string `json:"app_installed_team_id"`
	BotID              string `json:"bot_id"`
	// Blocks []Block `json:"blocks"`
	// State State `json:"state"`
}
