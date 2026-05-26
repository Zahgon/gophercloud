package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/workflow/v2/crontriggers"
	"github.com/gophercloud/gophercloud/v2/openstack/workflow/v2/workflows"
)

// CreateCronTrigger creates a cron trigger for the given workflow.
func CreateCronTrigger(t *testing.T, client *gophercloud.ServiceClient, workflow *workflows.Workflow) (*crontriggers.CronTrigger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteCronTrigger deletes a cron trigger.
func DeleteCronTrigger(t *testing.T, client *gophercloud.ServiceClient, crontrigger *crontriggers.CronTrigger) {
	_ = "STUB: not implemented"
	return
}

// GetCronTrigger gets a cron trigger.
func GetCronTrigger(t *testing.T, client *gophercloud.ServiceClient, id string) (*crontriggers.CronTrigger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListCronTriggers lists cron triggers.
func ListCronTriggers(t *testing.T, client *gophercloud.ServiceClient, opts crontriggers.ListOptsBuilder) ([]crontriggers.CronTrigger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
