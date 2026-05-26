package executions

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// CreateResult is the response of a Post operations. Call its Extract method to interpret it as an Execution.
type CreateResult struct {
	commonResult
}

// GetResult is the response of Get operations. Call its Extract method to interpret it as an Execution.
type GetResult struct {
	commonResult
}

// Extract helps to get an Execution struct from a Get or a Create function.
func (r commonResult) Extract() (*Execution, error) { _ = "STUB: not implemented"; return nil, nil }

// DeleteResult is the result from a Delete operation. Call its ExtractErr method to determine the success of the call.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Execution represents a workflow execution on OpenStack mistral API.
type Execution struct {
	// ID is the execution's unique ID.
	ID string `json:"id"`

	// CreatedAt contains the execution creation date.
	CreatedAt time.Time `json:"-"`

	// UpdatedAt is the last update of the execution.
	UpdatedAt time.Time `json:"-"`

	// RootExecutionID is the parent execution ID.
	RootExecutionID *string `json:"root_execution_id"`

	// TaskExecutionID is the task execution ID.
	TaskExecutionID *string `json:"task_execution_id"`

	// Description is the description of the execution.
	Description string `json:"description"`

	// Input contains the workflow input values.
	Input map[string]any `json:"-"`

	// Ouput contains the workflow output values.
	Output map[string]any `json:"-"`

	// Params contains workflow type specific parameters.
	Params map[string]any `json:"-"`

	// ProjectID is the project id owner of the execution.
	ProjectID string `json:"project_id"`

	// State is the current state of the execution. State can be one of: IDLE, RUNNING, SUCCESS, ERROR, PAUSED, CANCELLED.
	State string `json:"state"`

	// StateInfo contains an optional state information string.
	StateInfo *string `json:"state_info"`

	// WorkflowID is the ID of the workflow linked to the execution.
	WorkflowID string `json:"workflow_id"`

	// WorkflowName is the name of the workflow linked to the execution.
	WorkflowName string `json:"workflow_name"`

	// WorkflowNamespace is the namespace of the workflow linked to the execution.
	WorkflowNamespace string `json:"workflow_namespace"`
}

// UnmarshalJSON implements unmarshalling custom types
func (r *Execution) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// ExecutionPage contains a single page of all executions from a List call.
type ExecutionPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks if an ExecutionPage contains any results.
func (r ExecutionPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL finds the next page URL in a page in order to navigate to the next page of results.
func (r ExecutionPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractExecutions get the list of executions from a page acquired from the List call.
func ExtractExecutions(r pagination.Page) ([]Execution, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
