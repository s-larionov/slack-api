package file

import (
	"slack/events/base"
)

// A file was deleted
// https://api.slack.com/events/file_deleted
// Expected scopes: files:read
type Deleted struct {
	base.Event
	FileID  string `json:"file_id"`
	EventTS string `json:"event_ts"`
}
