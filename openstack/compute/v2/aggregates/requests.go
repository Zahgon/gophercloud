package aggregates

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List makes a request against the API to list aggregates.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAggregatesCreateMap() (map[string]any, error)
}

type CreateOpts struct {
	// The name of the host aggregate.
	Name string `json:"name" required:"true"`

	// The availability zone of the host aggregate.
	// You should use a custom availability zone rather than
	// the default returned by the os-availability-zone API.
	// The availability zone must not include ‘:’ in its name.
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

func (opts CreateOpts) ToAggregatesCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create makes a request against the API to create an aggregate.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete makes a request against the API to delete an aggregate.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, aggregateID int) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Get makes a request against the API to get details for a specific aggregate.
func Get(ctx context.Context, client *gophercloud.ServiceClient, aggregateID int) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToAggregatesUpdateMap() (map[string]any, error)
}

type UpdateOpts struct {
	// The name of the host aggregate.
	Name *string `json:"name,omitempty"`

	// The availability zone of the host aggregate.
	// You should use a custom availability zone rather than
	// the default returned by the os-availability-zone API.
	// The availability zone must not include ‘:’ in its name.
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (opts UpdateOpts) ToAggregatesUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update makes a request against the API to update a specific aggregate.
func Update(ctx context.Context, client *gophercloud.ServiceClient, aggregateID int, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

type AddHostOpts struct {
	// The name of the host.
	Host string `json:"host" required:"true"`
}

func (opts AddHostOpts) ToAggregatesAddHostMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddHost makes a request against the API to add host to a specific aggregate.
func AddHost(ctx context.Context, client *gophercloud.ServiceClient, aggregateID int, opts AddHostOpts) (r ActionResult) {
	_ = "STUB: not implemented"
	return *new(ActionResult)
}

type RemoveHostOpts struct {
	// The name of the host.
	Host string `json:"host" required:"true"`
}

func (opts RemoveHostOpts) ToAggregatesRemoveHostMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveHost makes a request against the API to remove host from a specific aggregate.
func RemoveHost(ctx context.Context, client *gophercloud.ServiceClient, aggregateID int, opts RemoveHostOpts) (r ActionResult) {
	_ = "STUB: not implemented"
	return *new(ActionResult)
}

type SetMetadataOpts struct {
	Metadata map[string]any `json:"metadata" required:"true"`
}

func (opts SetMetadataOpts) ToSetMetadataMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetMetadata makes a request against the API to set metadata to a specific aggregate.
func SetMetadata(ctx context.Context, client *gophercloud.ServiceClient, aggregateID int, opts SetMetadataOpts) (r ActionResult) {
	_ = "STUB: not implemented"
	return *new(ActionResult)
}
