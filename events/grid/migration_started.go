package grid

import (
	"slack/events/base"
)

// An enterprise grid migration has started on this workspace.
// https://api.slack.com/events/grid_migration_started
// Expected scopes: none
type MigrationStarted struct {
	base.Event
	EnterpriseID string `json:"enterprise_id"`
}
