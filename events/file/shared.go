package file

import (
	"slack/events/base"
)

// A file was shared
// https://api.slack.com/events/file_shared
// Expected scopes: files:read
type Shared struct {
	base.Event
	FileID string     `json:"file_id"`
	File   SharedFile `json:"file"`
}

type SharedFile struct {
	ID string `json:"id"`
}
