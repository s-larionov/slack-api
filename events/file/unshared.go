package file

import (
	"slack/events/base"
)

// A file was unshared
// https://api.slack.com/events/file_unshared
// Expected scopes: files:read
type Unshared struct {
	base.Event
	FileID string       `json:"file_id"`
	File   UnsharedFile `json:"file"`
}

type UnsharedFile struct {
	ID string `json:"id"`
}
