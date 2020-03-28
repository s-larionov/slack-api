package grid

import (
	"slack/events/base"
)

// An enterprise grid migration has finished on this workspace.
// https://api.slack.com/events/grid_migration_finished
// Expected scopes: none
type MigrationFinished struct {
	base.Event
	EnterpriseID string `json:"enterprise_id"`
}
