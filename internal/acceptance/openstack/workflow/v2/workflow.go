package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/workflow/v2/workflows"
)

// GetEchoWorkflowDefinition returns a simple workflow definition that does nothing except a simple "echo" command.
func GetEchoWorkflowDefinition(workflowName string) string { _ = "STUB: not implemented"; return "" }

// CreateWorkflow creates a workflow on Mistral API.
// The created workflow is a dummy workflow that performs a simple echo.
func CreateWorkflow(t *testing.T, client *gophercloud.ServiceClient) (*workflows.Workflow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteWorkflow deletes the given workflow.
func DeleteWorkflow(t *testing.T, client *gophercloud.ServiceClient, workflow *workflows.Workflow) {
	_ = "STUB: not implemented"
	return
}

// GetWorkflow gets a workflow.
func GetWorkflow(t *testing.T, client *gophercloud.ServiceClient, id string) (*workflows.Workflow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListWorkflows lists the workflows.
func ListWorkflows(t *testing.T, client *gophercloud.ServiceClient, opts workflows.ListOptsBuilder) ([]workflows.Workflow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
