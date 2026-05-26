package crontriggers

import (
	"context"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extension to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToCronTriggerCreateMap() (map[string]any, error)
}

// CreateOpts specifies parameters used to create a cron trigger.
type CreateOpts struct {
	// Name is the cron trigger name.
	Name string `json:"name"`

	// Pattern is a Unix crontab patterns format to execute the workflow.
	Pattern string `json:"pattern"`

	// RemainingExecutions sets the number of executions for the trigger.
	RemainingExecutions int `json:"remaining_executions,omitempty"`

	// WorkflowID is the unique id of the workflow.
	WorkflowID string `json:"workflow_id,omitempty" or:"WorkflowName"`

	// WorkflowName is the name of the workflow.
	// It is recommended to refer to workflow by the WorkflowID parameter instead of WorkflowName.
	WorkflowName string `json:"workflow_name,omitempty" or:"WorkflowID"`

	// WorkflowParams defines workflow type specific parameters.
	WorkflowParams map[string]any `json:"workflow_params,omitempty"`

	// WorkflowInput defines workflow input values.
	WorkflowInput map[string]any `json:"workflow_input,omitempty"`

	// FirstExecutionTime defines the first execution time of the trigger.
	FirstExecutionTime *time.Time `json:"-"`
}

// ToCronTriggerCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToCronTriggerCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new cron trigger.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes the specified cron trigger.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Get retrieves details of a single cron trigger.
// Use Extract to convert its result into an CronTrigger.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ListOptsBuilder allows extension to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToCronTriggerListQuery() (string, error)
}

// ListOpts filters the result returned by the List() function.
type ListOpts struct {
	// WorkflowName allows to filter by workflow name.
	WorkflowName *ListFilter `q:"-"`
	// WorkflowID allows to filter by workflow id.
	WorkflowID string `q:"workflow_id"`
	// WorkflowInput allows to filter by specific workflow inputs.
	WorkflowInput map[string]any `q:"-"`
	// WorkflowParams allows to filter by specific workflow parameters.
	WorkflowParams map[string]any `q:"-"`
	// Scope filters by the trigger's scope.
	// Values can be "private" or "public".
	Scope string `q:"scope"`
	// Name allows to filter by trigger name.
	Name *ListFilter `q:"-"`
	// Pattern allows to filter by pattern.
	Pattern *ListFilter `q:"-"`
	// RemainingExecutions allows to filter by remaining executions.
	RemainingExecutions *ListIntFilter `q:"-"`
	// FirstExecutionTime allows to filter by first execution time.
	FirstExecutionTime *ListDateFilter `q:"-"`
	// NextExecutionTime allows to filter by next execution time.
	NextExecutionTime *ListDateFilter `q:"-"`
	// CreatedAt allows to filter by trigger creation date.
	CreatedAt *ListDateFilter `q:"-"`
	// UpdatedAt allows to filter by trigger last update date.
	UpdatedAt *ListDateFilter `q:"-"`
	// ProjectID allows to filter by given project id. Admin required.
	ProjectID string `q:"project_id"`
	// AllProjects requests to get executions of all projects. Admin required.
	AllProjects int `q:"all_projects"`
	// SortDirs allows to select sort direction.
	// It can be "asc" or "desc" (default).
	SortDirs string `q:"sort_dirs"`
	// SortKeys allows to sort by one of the cron trigger attributes.
	SortKeys string `q:"sort_keys"`
	// Marker and Limit control paging.
	// Marker instructs List where to start listing from.
	Marker string `q:"marker"`
	// Limit instructs List to refrain from sending excessively large lists of
	// cron triggers.
	Limit int `q:"limit"`
}

// ListFilter allows to filter string parameters with different filters.
// Empty value for Filter checks for equality.
type ListFilter struct {
	Filter FilterType
	Value  string
}

func (l ListFilter) String() string { _ = "STUB: not implemented"; return "" }

// ListDateFilter allows to filter date parameters with different filters.
// Empty value for Filter checks for equality.
type ListDateFilter struct {
	Filter FilterType
	Value  time.Time
}

func (l ListDateFilter) String() string { _ = "STUB: not implemented"; return "" }

// ListIntFilter allows to filter integer parameters with different filters.
// Empty value for Filter checks for equality.
type ListIntFilter struct {
	Filter FilterType
	Value  int
}

func (l ListIntFilter) String() string { _ = "STUB: not implemented"; return "" }

// FilterType represents a valid filter to use for filtering executions.
type FilterType string

const (
	// FilterEQ checks equality.
	FilterEQ = "eq"
	// FilterNEQ checks non equality.
	FilterNEQ = "neq"
	// FilterIN checks for belonging in a list, comma separated.
	FilterIN = "in"
	// FilterNIN checks for values that does not belong from a list, comma separated.
	FilterNIN = "nin"
	// FilterGT checks for values strictly greater.
	FilterGT = "gt"
	// FilterGTE checks for values greater or equal.
	FilterGTE = "gte"
	// FilterLT checks for values strictly lower.
	FilterLT = "lt"
	// FilterLTE checks for values lower or equal.
	FilterLTE = "lte"
	// FilterHas checks for values that contains the requested parameter.
	FilterHas = "has"
)

// ToCronTriggerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToCronTriggerListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List performs a call to list cron triggers.
// You may provide options to filter the results.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
