package request

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToTransferRequestListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned.
// https://developer.openstack.org/api-ref/dns/
type ListOpts struct {
	Status string `q:"status"`
}

// ToTransferRequestListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTransferRequestListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List implements a transfer request List request.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get returns information about a transfer request, given its ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, transferRequestID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional attributes to the
// Create request.
type CreateOptsBuilder interface {
	ToTransferRequestCreateMap() (map[string]any, error)
}

// CreateOpts specifies the attributes used to create a transfer request.
type CreateOpts struct {
	// TargetProjectID is ID that the request will be limited to. No other project
	// will be allowed to accept this request.
	TargetProjectID string `json:"target_project_id,omitempty"`

	// Description of the transfer request.
	Description string `json:"description,omitempty"`
}

// ToTransferRequestCreateMap formats an CreateOpts structure into a request body.
func (opts CreateOpts) ToTransferRequestCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create implements a transfer request create request.
func Create(ctx context.Context, client *gophercloud.ServiceClient, zoneID string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToTransferRequestUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies the attributes to update a transfer request.
type UpdateOpts struct {
	// TargetProjectID is ID that the request will be limited to. No other project
	// will be allowed to accept this request.
	TargetProjectID string `json:"target_project_id,omitempty"`

	// Description of the transfer request.
	Description string `json:"description,omitempty"`
}

// ToTransferRequestUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToTransferRequestUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update implements a transfer request update request.
func Update(ctx context.Context, client *gophercloud.ServiceClient, transferID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete implements a transfer request delete request.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, transferID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
