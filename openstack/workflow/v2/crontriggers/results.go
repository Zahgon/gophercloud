package crontriggers

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// CreateResult is the response of a Post operations. Call its Extract method to interpret it as a CronTrigger.
type CreateResult struct {
	commonResult
}

// GetResult is the response of Get operations. Call its Extract method to interpret it as a CronTrigger.
type GetResult struct {
	commonResult
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr method to determine the success of the call.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Extract helps to get a CronTrigger struct from a Get or a Create function.
func (r commonResult) Extract() (*CronTrigger, error) { _ = "STUB: not implemented"; return nil, nil }

// CronTrigger represents a workflow cron trigger on OpenStack mistral API.
type CronTrigger struct {
	// ID is the cron trigger's unique ID.
	ID string `json:"id"`

	// Name is the name of the cron trigger.
	Name string `json:"name"`

	// Pattern is the cron-like style pattern to execute the workflow.
	// Example of value: "* * * * *"
	Pattern string `json:"pattern"`

	// ProjectID is the project id owner of the cron trigger.
	ProjectID string `json:"project_id"`

	// RemainingExecutions is the number of remaining executions of this trigger.
	RemainingExecutions int `json:"remaining_executions"`

	// Scope is the scope of the trigger.
	// Values can be "private" or "public".
	Scope string `json:"scope"`

	// WorkflowID is the ID of the workflow linked to the trigger.
	WorkflowID string `json:"workflow_id"`

	// WorkflowName is the name of the workflow linked to the trigger.
	WorkflowName string `json:"workflow_name"`

	// WorkflowInput contains the workflow input values.
	WorkflowInput map[string]any `json:"-"`

	// WorkflowParams contains workflow type specific parameters.
	WorkflowParams map[string]any `json:"-"`

	// CreatedAt contains the cron trigger creation date.
	CreatedAt time.Time `json:"-"`

	// FirstExecutionTime is the date of the first execution of the trigger.
	FirstExecutionTime *time.Time `json:"-"`

	// NextExecutionTime is the date of the next execution of the trigger.
	NextExecutionTime *time.Time `json:"-"`
}

// UnmarshalJSON implements unmarshalling custom types
func (r *CronTrigger) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// CronTriggerPage contains a single page of all cron triggers from a List call.
type CronTriggerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks if an CronTriggerPage contains any results.
func (r CronTriggerPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL finds the next page URL in a page in order to navigate to the next page of results.
func (r CronTriggerPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractCronTriggers get the list of cron triggers from a page acquired from the List call.
func ExtractCronTriggers(r pagination.Page) ([]CronTrigger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
