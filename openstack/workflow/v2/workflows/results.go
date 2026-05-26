package workflows

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateResult is the response of a Post operations. Call its Extract method to interpret it as a list of Workflows.
type CreateResult struct {
	gophercloud.Result
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr method to determine the success of the call.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Extract helps to get created Workflow struct from a Create function.
func (r CreateResult) Extract() ([]Workflow, error) { _ = "STUB: not implemented"; return nil, nil }

// GetResult is the response of Get operations. Call its Extract method to interpret it as a Workflow.
type GetResult struct {
	gophercloud.Result
}

// Extract helps to get a Workflow struct from a Get function.
func (r GetResult) Extract() (*Workflow, error) { _ = "STUB: not implemented"; return nil, nil }

// Workflow represents a workflow execution on OpenStack mistral API.
type Workflow struct {
	// ID is the workflow's unique ID.
	ID string `json:"id"`

	// Definition is the workflow definition in Mistral v2 DSL.
	Definition string `json:"definition"`

	// Name is the name of the workflow.
	Name string `json:"name"`

	// Namespace is the namespace of the workflow.
	Namespace string `json:"namespace"`

	// Input represents the needed input to execute the workflow.
	// This parameter is a list of each input, comma separated.
	Input string `json:"input"`

	// ProjectID is the project id owner of the workflow.
	ProjectID string `json:"project_id"`

	// Scope is the scope of the workflow.
	// Values can be "private" or "public".
	Scope string `json:"scope"`

	// Tags is a list of tags associated to the workflow.
	Tags []string `json:"tags"`

	// CreatedAt is the creation date of the workflow.
	CreatedAt time.Time `json:"-"`

	// UpdatedAt is the last update date of the workflow.
	UpdatedAt *time.Time `json:"-"`
}

// UnmarshalJSON implements unmarshalling custom types
func (r *Workflow) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// WorkflowPage contains a single page of all workflows from a List call.
type WorkflowPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks if an WorkflowPage contains any results.
func (r WorkflowPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL finds the next page URL in a page in order to navigate to the next page of results.
func (r WorkflowPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractWorkflows get the list of cron triggers from a page acquired from the List call.
func ExtractWorkflows(r pagination.Page) ([]Workflow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
