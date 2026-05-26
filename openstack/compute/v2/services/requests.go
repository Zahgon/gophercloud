package services

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request.
type ListOptsBuilder interface {
	ToServicesListQuery() (string, error)
}

// ListOpts represents options to list services.
type ListOpts struct {
	Binary string `q:"binary"`
	Host   string `q:"host"`
}

// ToServicesListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToServicesListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request against the API to list services.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

type ServiceStatus string

const (
	// ServiceEnabled is used to mark a service as being enabled.
	ServiceEnabled ServiceStatus = "enabled"

	// ServiceDisabled is used to mark a service as being disabled.
	ServiceDisabled ServiceStatus = "disabled"
)

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToServiceUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies the base attributes that may be updated on a service.
type UpdateOpts struct {
	// Status represents the new service status. One of enabled or disabled.
	Status ServiceStatus `json:"status,omitempty"`

	// DisabledReason represents the reason for disabling a service.
	DisabledReason string `json:"disabled_reason,omitempty"`

	// ForcedDown is a manual override to tell nova that the service in question
	// has been fenced manually by the operations team.
	ForcedDown *bool `json:"forced_down,omitempty"`
}

// ToServiceUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToServiceUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update requests that various attributes of the indicated service be changed.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete will delete the existing service with the provided ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
