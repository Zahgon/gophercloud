package conductors

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToConductorListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the conductor attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListOpts struct {
	// One or more fields to be returned in the response.
	Fields []string `q:"fields" format:"comma-separated"`

	// Requests a page size of items.
	Limit int `q:"limit"`

	// The ID of the last-seen item.
	Marker string `q:"marker"`

	// Sorts the response by the requested sort direction.
	SortDir string `q:"sort_dir"`

	// Sorts the response by the this attribute value.
	SortKey string `q:"sort_key"`

	// Provide additional information for the BIOS Settings
	Detail bool `q:"detail"`
}

// ToConductorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToConductorListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request against the API to list conductors accessible to you.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get requests details on a single conductor by hostname
func Get(ctx context.Context, client *gophercloud.ServiceClient, name string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
