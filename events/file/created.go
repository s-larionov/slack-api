package file

import (
	"slack/events/base"
)

// A file was created
// https://api.slack.com/events/file_created
// Expected scopes: files:read
type Created struct {
	base.Event
	FileID string      `json:"file_id"`
	File   CreatedFile `json:"file"`
}

type CreatedFile struct {
	ID string `json:"id"`
}
