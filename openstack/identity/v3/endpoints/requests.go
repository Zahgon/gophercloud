package endpoints

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type CreateOptsBuilder interface {
	ToEndpointCreateMap() (map[string]any, error)
}

// CreateOpts contains the subset of Endpoint attributes that should be used
// to create an Endpoint.
type CreateOpts struct {
	// Availability is the interface type of the Endpoint (admin, internal,
	// or public), referenced by the gophercloud.Availability type.
	Availability gophercloud.Availability `json:"interface" required:"true"`

	// Description is the description of the Endpoint.
	Description string `json:"description,omitempty"`

	// Enabled indicates whether the Endpoint appears in the service
	// or not.
	Enabled *bool `json:"enabled,omitempty"`

	// Name is the name of the Endpoint.
	Name string `json:"name,omitempty"`

	// Region is the region the Endpoint is located in.
	// This field can be omitted or left as a blank string.
	Region string `json:"region,omitempty"`

	// URL is the url of the Endpoint.
	URL string `json:"url" required:"true"`

	// ServiceID is the ID of the service the Endpoint refers to.
	ServiceID string `json:"service_id" required:"true"`
}

// ToEndpointCreateMap builds a request body from the Endpoint Create options.
func (opts CreateOpts) ToEndpointCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create inserts a new Endpoint into the service catalog.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToEndpointListParams() (string, error)
}

// ListOpts allows finer control over the endpoints returned by a List call.
// All fields are optional.
type ListOpts struct {
	// Availability is the interface type of the Endpoint (admin, internal,
	// or public), referenced by the gophercloud.Availability type.
	Availability gophercloud.Availability `q:"interface"`

	// ServiceID is the ID of the service the Endpoint refers to.
	ServiceID string `q:"service_id"`

	// RegionID is the ID of the region the Endpoint refers to.
	RegionID string `q:"region_id"`
}

// ToEndpointListParams builds a list request from the List options.
func (opts ListOpts) ToEndpointListParams() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List enumerates endpoints in a paginated collection, optionally filtered
// by ListOpts criteria.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single endpoint, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add parameters to the Update request.
type UpdateOptsBuilder interface {
	ToEndpointUpdateMap() (map[string]any, error)
}

// UpdateOpts contains the subset of Endpoint attributes that should be used to
// update an Endpoint.
type UpdateOpts struct {
	// Availability is the interface type of the Endpoint (admin, internal,
	// or public), referenced by the gophercloud.Availability type.
	Availability gophercloud.Availability `json:"interface,omitempty"`

	// Name is the name of the Endpoint.
	Name string `json:"name,omitempty"`

	// Enabled indicates whether the Endpoint appears in the service
	// or not.
	Enabled *bool `json:"enabled,omitempty"`

	// Region is the region the Endpoint is located in.
	// This field can be omitted or left as a blank string.
	Region string `json:"region,omitempty"`

	// URL is the url of the Endpoint.
	URL string `json:"url,omitempty"`

	// ServiceID is the ID of the service the Endpoint refers to.
	ServiceID string `json:"service_id,omitempty"`

	// Description is an updated description of the endpoint.
	Description *string `json:"description,omitempty"`
}

// ToEndpointUpdateMap builds an update request body from the Update options.
func (opts UpdateOpts) ToEndpointUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update changes an existing endpoint with new data.
func Update(ctx context.Context, client *gophercloud.ServiceClient, endpointID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete removes an endpoint from the service catalog.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, endpointID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
