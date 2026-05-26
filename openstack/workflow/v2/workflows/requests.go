package workflows

import (
	"context"
	"io"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extension to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToWorkflowCreateParams() (io.Reader, string, error)
}

// CreateOpts specifies parameters used to create a cron trigger.
type CreateOpts struct {
	// Scope is the scope of the workflow.
	// Allowed values are "private" and "public".
	Scope string `q:"scope"`

	// Namespace will define the namespace of the workflow.
	Namespace string `q:"namespace"`

	// Definition is the workflow definition written in Mistral Workflow Language v2.
	Definition io.Reader
}

// ToWorkflowCreateParams constructs a request query string from CreateOpts.
func (opts CreateOpts) ToWorkflowCreateParams() (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

// Create requests the creation of a new execution.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Drop default JSON Accept header

// Delete deletes the specified execution.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Get retrieves details of a single execution.
// Use Extract to convert its result into an Workflow.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ListOptsBuilder allows extension to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToWorkflowListQuery() (string, error)
}

// ListOpts filters the result returned by the List() function.
type ListOpts struct {
	// Scope filters by the workflow's scope.
	// Values can be "private" or "public".
	Scope string `q:"scope"`
	// CreatedAt allows to filter by workflow creation date.
	CreatedAt *ListDateFilter `q:"-"`
	// UpdatedAt allows to filter by last execution update date.
	UpdatedAt *ListDateFilter `q:"-"`
	// Name allows to filter by workflow name.
	Name *ListFilter `q:"-"`
	// Tags allows to filter by tags.
	Tags []string
	// Definition allows to filter by workflow definition.
	Definition *ListFilter `q:"-"`
	// Namespace allows to filter by workflow namespace.
	Namespace *ListFilter `q:"-"`
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
	// ProjectID allows to filter by given project id. Admin required.
	ProjectID string `q:"project_id"`
	// AllProjects requests to get executions of all projects. Admin required.
	AllProjects int `q:"all_projects"`
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

// ToWorkflowListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToWorkflowListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List performs a call to list cron triggers.
// You may provide options to filter the results.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
