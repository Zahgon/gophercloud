package traits

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request.
type ListOptsBuilder interface {
	ToTraitListQuery() (string, error)
}

// ListOpts allows the filtering of traits. Filtering is achieved by passing in struct
// field values that map to the trait attributes you want to see returned.
type ListOpts struct {
	// Name is a string used to filter traits by name.
	// It supports startswith operator to filter the traits whose name begins with
	// a specific prefix, e.g. name=startswith:CUSTOM
	// in operator filters the traits whose name is in the specified list,
	// e.g. name=in:HW_CPU_X86_AVX,HW_CPU_X86_SSE,HW_CPU_X86_INVALID_FEATURE
	Name string `q:"name"`

	// Associated is a boolean used to filter traits by whether they are associated with
	// at least one resource provider.
	Associated *bool `q:"associated"`
}

// ToTraitListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTraitListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List retrieves a list of traits.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get confirms the existence of a trait.
func Get(ctx context.Context, client *gophercloud.ServiceClient, traitName string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Create creates a new trait.
func Create(ctx context.Context, client *gophercloud.ServiceClient, traitName string) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes the trait specified by name.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, traitName string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
