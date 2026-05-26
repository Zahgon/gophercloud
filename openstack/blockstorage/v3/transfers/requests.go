package transfers

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToCreateMap() (map[string]any, error)
}

// CreateOpts contains options for a Volume transfer.
type CreateOpts struct {
	// The ID of the volume to transfer.
	VolumeID string `json:"volume_id" required:"true"`

	// The name of the volume transfer
	Name string `json:"name,omitempty"`
}

// ToCreateMap assembles a request body based on the contents of a
// TransferOpts.
func (opts CreateOpts) ToCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create will create a volume tranfer request based on the values in CreateOpts.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// AcceptOptsBuilder allows extensions to add additional parameters to the
// Accept request.
type AcceptOptsBuilder interface {
	ToAcceptMap() (map[string]any, error)
}

// AcceptOpts contains options for a Volume transfer accept reqeust.
type AcceptOpts struct {
	// The auth key of the volume transfer to accept.
	AuthKey string `json:"auth_key" required:"true"`
}

// ToAcceptMap assembles a request body based on the contents of a
// AcceptOpts.
func (opts AcceptOpts) ToAcceptMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Accept will accept a volume tranfer request based on the values in AcceptOpts.
func Accept(ctx context.Context, client *gophercloud.ServiceClient, id string, opts AcceptOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a volume transfer.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToTransferListQuery() (string, error)
}

// ListOpts holds options for listing Transfers. It is passed to the transfers.List
// function.
type ListOpts struct {
	// AllTenants will retrieve transfers of all tenants/projects.
	AllTenants bool `q:"all_tenants"`

	// Comma-separated list of sort keys and optional sort directions in the
	// form of <key>[:<direction>].
	Sort string `q:"sort"`

	// Requests a page size of items.
	Limit int `q:"limit"`

	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`

	// The ID of the last-seen item.
	Marker string `q:"marker"`
}

// ToTransferListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTransferListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns Transfers optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves the Transfer with the provided ID. To extract the Transfer object
// from the response, call the Extract method on the GetResult.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
