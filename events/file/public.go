package file

import (
	"slack/events/base"
)

// A file was made public
// https://api.slack.com/events/file_public
// Expected scopes: files:read
type Public struct {
	base.Event
	FileID string     `json:"file_id"`
	File   PublicFile `json:"file"`
}

type PublicFile struct {
	ID string `json:"id"`
}
