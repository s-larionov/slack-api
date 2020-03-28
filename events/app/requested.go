package app

import (
	"slack/events/base"
)

// User requested an app
// https://api.slack.com/events/app_requested
// Expected scopes: admin.apps:read
type Requested struct {
	base.Event
	Request Request `json:"request"`
}

type Request struct {
	ID                 string     `json:"id"`
	App                App        `json:"app"`
	PreviousResolution Resolution `json:"previous_resolution"`
	User               User       `json:"user"`
	Team               Team       `json:"team"`
	Scopes             []Scope    `json:"scopes"`
	Message            string     `json:"message"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Team struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type App struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	HelpURL                string `json:"help_url"`
	PrivacyPolicyURL       string `json:"privacy_policy_url"`
	AppHomepageURL         string `json:"app_homepage_url"`
	AppDirectoryURL        string `json:"app_directory_url"`
	IsAppDirectoryApproved bool   `json:"is_app_directory_approved"`
	IsInternal             bool   `json:"is_internal"`
	AdditionalInfo         string `json:"additional_info"`
}

const (
	ResolutionStatusApproved   = "approved"
	ResolutionStatusRestricted = "restricted"
)

type ResolutionStatus string

type Resolution struct {
	Status ResolutionStatus `json:"status"`
	Scopes []Scope          `json:"scopes"`
}

type Scope struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSensitive bool   `json:"is_sensitive"`
	TokenType   string `json:"token_type"`
}
