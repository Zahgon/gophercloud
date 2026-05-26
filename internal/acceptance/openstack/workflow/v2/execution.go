package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/workflow/v2/executions"
	"github.com/gophercloud/gophercloud/v2/openstack/workflow/v2/workflows"
)

// CreateExecution creates an execution for the given workflow.
func CreateExecution(t *testing.T, client *gophercloud.ServiceClient, workflow *workflows.Workflow) (*executions.Execution, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteExecution deletes an execution.
func DeleteExecution(t *testing.T, client *gophercloud.ServiceClient, execution *executions.Execution) {
	_ = "STUB: not implemented"
	return
}

// ListExecutions lists the executions.
func ListExecutions(t *testing.T, client *gophercloud.ServiceClient, opts executions.ListOptsBuilder) ([]executions.Execution, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
