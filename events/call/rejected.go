package call

import (
	"slack/events/base"
)

// A call was rejected
// https://api.slack.com/events/call_rejected
// Expected scopes: calls:read
type Rejected struct {
	base.Event
}
