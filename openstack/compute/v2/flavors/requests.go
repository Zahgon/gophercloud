package flavors

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFlavorListQuery() (string, error)
}

/*
AccessType maps to OpenStack's Flavor.is_public field. Although the is_public
field is boolean, the request options are ternary, which is why AccessType is
a string. The following values are allowed:

The AccessType arguement is optional, and if it is not supplied, OpenStack
returns the PublicAccess flavors.
*/
type AccessType string

const (
	// PublicAccess returns public flavors and private flavors associated with
	// that project.
	PublicAccess AccessType = "true"

	// PrivateAccess (admin only) returns private flavors, across all projects.
	PrivateAccess AccessType = "false"

	// AllAccess (admin only) returns public and private flavors across all
	// projects.
	AllAccess AccessType = "None"
)

/*
ListOpts filters the results returned by the List() function.
For example, a flavor with a minDisk field of 10 will not be returned if you
specify MinDisk set to 20.

Typically, software will use the last ID of the previous call to List to set
the Marker for the current call.
*/
type ListOpts struct {
	// ChangesSince, if provided, instructs List to return only those things which
	// have changed since the timestamp provided.
	ChangesSince string `q:"changes-since"`

	// MinDisk and MinRAM, if provided, elides flavors which do not meet your
	// criteria.
	MinDisk int `q:"minDisk"`
	MinRAM  int `q:"minRam"`

	// SortDir allows to select sort direction.
	// It can be "asc" or "desc" (default).
	SortDir string `q:"sort_dir"`

	// SortKey allows to sort by one of the flavors attributes.
	// Default is flavorid.
	SortKey string `q:"sort_key"`

	// Marker and Limit control paging.
	// Marker instructs List where to start listing from.
	Marker string `q:"marker"`

	// Limit instructs List to refrain from sending excessively large lists of
	// flavors.
	Limit int `q:"limit"`

	// AccessType, if provided, instructs List which set of flavors to return.
	// If IsPublic not provided, flavors for the current project are returned.
	AccessType AccessType `q:"is_public"`
}

// ToFlavorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFlavorListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ListDetail instructs OpenStack to provide a list of flavors.
// You may provide criteria by which List curtails its results for easier
// processing.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

type CreateOptsBuilder interface {
	ToFlavorCreateMap() (map[string]any, error)
}

// CreateOpts specifies parameters used for creating a flavor.
type CreateOpts struct {
	// Name is the name of the flavor.
	Name string `json:"name" required:"true"`

	// RAM is the memory of the flavor, measured in MB.
	RAM int `json:"ram" required:"true"`

	// VCPUs is the number of vcpus for the flavor.
	VCPUs int `json:"vcpus" required:"true"`

	// Disk the amount of root disk space, measured in GB.
	Disk *int `json:"disk" required:"true"`

	// ID is a unique ID for the flavor.
	ID string `json:"id,omitempty"`

	// Swap is the amount of swap space for the flavor, measured in MB.
	Swap *int `json:"swap,omitempty"`

	// RxTxFactor alters the network bandwidth of a flavor.
	RxTxFactor float64 `json:"rxtx_factor,omitempty"`

	// IsPublic flags a flavor as being available to all projects or not.
	IsPublic *bool `json:"os-flavor-access:is_public,omitempty"`

	// Ephemeral is the amount of ephemeral disk space, measured in GB.
	Ephemeral *int `json:"OS-FLV-EXT-DATA:ephemeral,omitempty"`

	// Description is a free form description of the flavor. Limited to
	// 65535 characters in length. Only printable characters are allowed.
	// New in version 2.55
	Description string `json:"description,omitempty"`
}

// ToFlavorCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToFlavorCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new flavor.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

type UpdateOptsBuilder interface {
	ToFlavorUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies parameters used for updating a flavor.
type UpdateOpts struct {
	// Description is a free form description of the flavor. Limited to
	// 65535 characters in length. Only printable characters are allowed.
	// New in version 2.55
	Description *string `json:"description,omitempty"`
}

// ToFlavorUpdateMap constructs a request body from UpdateOpts.
func (opts UpdateOpts) ToFlavorUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update requests the update of a new flavor.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Get retrieves details of a single flavor. Use Extract to convert its
// result into a Flavor.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Delete deletes the specified flavor ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListAccesses retrieves the tenants which have access to a flavor.
func ListAccesses(client *gophercloud.ServiceClient, id string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// AddAccessOptsBuilder allows extensions to add additional parameters to the
// AddAccess requests.
type AddAccessOptsBuilder interface {
	ToFlavorAddAccessMap() (map[string]any, error)
}

// AddAccessOpts represents options for adding access to a flavor.
type AddAccessOpts struct {
	// Tenant is the project/tenant ID to grant access.
	Tenant string `json:"tenant"`
}

// ToFlavorAddAccessMap constructs a request body from AddAccessOpts.
func (opts AddAccessOpts) ToFlavorAddAccessMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddAccess grants a tenant/project access to a flavor.
func AddAccess(ctx context.Context, client *gophercloud.ServiceClient, id string, opts AddAccessOptsBuilder) (r AddAccessResult) {
	_ = "STUB: not implemented"
	return *new(AddAccessResult)
}

// RemoveAccessOptsBuilder allows extensions to add additional parameters to the
// RemoveAccess requests.
type RemoveAccessOptsBuilder interface {
	ToFlavorRemoveAccessMap() (map[string]any, error)
}

// RemoveAccessOpts represents options for removing access to a flavor.
type RemoveAccessOpts struct {
	// Tenant is the project/tenant ID to grant access.
	Tenant string `json:"tenant"`
}

// ToFlavorRemoveAccessMap constructs a request body from RemoveAccessOpts.
func (opts RemoveAccessOpts) ToFlavorRemoveAccessMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveAccess removes/revokes a tenant/project access to a flavor.
func RemoveAccess(ctx context.Context, client *gophercloud.ServiceClient, id string, opts RemoveAccessOptsBuilder) (r RemoveAccessResult) {
	_ = "STUB: not implemented"
	return *new(RemoveAccessResult)
}

// ExtraSpecs requests all the extra-specs for the given flavor ID.
func ListExtraSpecs(ctx context.Context, client *gophercloud.ServiceClient, flavorID string) (r ListExtraSpecsResult) {
	_ = "STUB: not implemented"
	return *new(ListExtraSpecsResult)
}

func GetExtraSpec(ctx context.Context, client *gophercloud.ServiceClient, flavorID string, key string) (r GetExtraSpecResult) {
	_ = "STUB: not implemented"
	return *new(GetExtraSpecResult)
}

// CreateExtraSpecsOptsBuilder allows extensions to add additional parameters to the
// CreateExtraSpecs requests.
type CreateExtraSpecsOptsBuilder interface {
	ToFlavorExtraSpecsCreateMap() (map[string]any, error)
}

// ExtraSpecsOpts is a map that contains key-value pairs.
type ExtraSpecsOpts map[string]string

// ToFlavorExtraSpecsCreateMap assembles a body for a Create request based on
// the contents of ExtraSpecsOpts.
func (opts ExtraSpecsOpts) ToFlavorExtraSpecsCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateExtraSpecs will create or update the extra-specs key-value pairs for
// the specified Flavor.
func CreateExtraSpecs(ctx context.Context, client *gophercloud.ServiceClient, flavorID string, opts CreateExtraSpecsOptsBuilder) (r CreateExtraSpecsResult) {
	_ = "STUB: not implemented"
	return *new(CreateExtraSpecsResult)
}

// UpdateExtraSpecOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateExtraSpecOptsBuilder interface {
	ToFlavorExtraSpecUpdateMap() (map[string]string, string, error)
}

// ToFlavorExtraSpecUpdateMap assembles a body for an Update request based on
// the contents of a ExtraSpecOpts.
func (opts ExtraSpecsOpts) ToFlavorExtraSpecUpdateMap() (map[string]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// UpdateExtraSpec will updates the value of the specified flavor's extra spec
// for the key in opts.
func UpdateExtraSpec(ctx context.Context, client *gophercloud.ServiceClient, flavorID string, opts UpdateExtraSpecOptsBuilder) (r UpdateExtraSpecResult) {
	_ = "STUB: not implemented"
	return *new(UpdateExtraSpecResult)
}

// DeleteExtraSpec will delete the key-value pair with the given key for the given
// flavor ID.
func DeleteExtraSpec(ctx context.Context, client *gophercloud.ServiceClient, flavorID, key string) (r DeleteExtraSpecResult) {
	_ = "STUB: not implemented"
	return *new(DeleteExtraSpecResult)
}
