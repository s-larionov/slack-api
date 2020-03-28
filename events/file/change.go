package file

import (
	"slack/events/base"
)

// A file was changed
// https://api.slack.com/events/file_change
// Expected scopes: files:read
type Change struct {
	base.Event
	FileID string      `json:"file_id"`
	File   ChangedFile `json:"file"`
}

type ChangedFile struct {
	ID string `json:"id"`
}
