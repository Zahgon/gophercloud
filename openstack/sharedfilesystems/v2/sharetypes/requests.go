package sharetypes

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToShareTypeCreateMap() (map[string]any, error)
}

// CreateOpts contains options for creating a ShareType. This object is
// passed to the sharetypes.Create function. For more information about
// these parameters, see the ShareType object.
type CreateOpts struct {
	// The share type name
	Name string `json:"name" required:"true"`
	// Indicates whether a share type is publicly accessible
	IsPublic bool `json:"os-share-type-access:is_public"`
	// The extra specifications for the share type
	ExtraSpecs ExtraSpecsOpts `json:"extra_specs" required:"true"`
}

// ExtraSpecsOpts represent the extra specifications that can be selected for a share type
type ExtraSpecsOpts struct {
	// An extra specification that defines the driver mode for share server, or storage, life cycle management
	DriverHandlesShareServers bool `json:"driver_handles_share_servers" required:"true"`
	// An extra specification that filters back ends by whether they do or do not support share snapshots
	SnapshotSupport *bool `json:"snapshot_support,omitempty"`
}

// ToShareTypeCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToShareTypeCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create will create a new ShareType based on the values in CreateOpts. To
// extract the ShareType object from the response, call the Extract method
// on the CreateResult.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete will delete the existing ShareType with the provided ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToShareTypeListQuery() (string, error)
}

// ListOpts holds options for listing ShareTypes. It is passed to the
// sharetypes.List function.
type ListOpts struct {
	// Select if public types, private types, or both should be listed
	IsPublic string `q:"is_public"`
}

// ToShareTypeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToShareTypeListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns ShareTypes optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetDefault will retrieve the default ShareType.
func GetDefault(ctx context.Context, client *gophercloud.ServiceClient) (r GetDefaultResult) {
	_ = "STUB: not implemented"
	return *new(GetDefaultResult)
}

// GetExtraSpecs will retrieve the extra specifications for a given ShareType.
func GetExtraSpecs(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetExtraSpecsResult) {
	_ = "STUB: not implemented"
	return *new(GetExtraSpecsResult)
}

// SetExtraSpecsOptsBuilder allows extensions to add additional parameters to the
// SetExtraSpecs request.
type SetExtraSpecsOptsBuilder interface {
	ToShareTypeSetExtraSpecsMap() (map[string]any, error)
}

type SetExtraSpecsOpts struct {
	// A list of all extra specifications to be added to a ShareType
	ExtraSpecs map[string]any `json:"extra_specs" required:"true"`
}

// ToShareTypeSetExtraSpecsMap assembles a request body based on the contents of a
// SetExtraSpecsOpts.
func (opts SetExtraSpecsOpts) ToShareTypeSetExtraSpecsMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetExtraSpecs will set new specifications for a ShareType based on the values
// in SetExtraSpecsOpts. To extract the extra specifications object from the response,
// call the Extract method on the SetExtraSpecsResult.
func SetExtraSpecs(ctx context.Context, client *gophercloud.ServiceClient, id string, opts SetExtraSpecsOptsBuilder) (r SetExtraSpecsResult) {
	_ = "STUB: not implemented"
	return *new(SetExtraSpecsResult)
}

// UnsetExtraSpecs will unset an extra specification for an existing ShareType.
func UnsetExtraSpecs(ctx context.Context, client *gophercloud.ServiceClient, id string, key string) (r UnsetExtraSpecsResult) {
	_ = "STUB: not implemented"
	return *new(UnsetExtraSpecsResult)
}

// ShowAccess will show access details for an existing ShareType.
func ShowAccess(ctx context.Context, client *gophercloud.ServiceClient, id string) (r ShowAccessResult) {
	_ = "STUB: not implemented"
	return *new(ShowAccessResult)
}

// AddAccessOptsBuilder allows extensions to add additional parameters to the
// AddAccess
type AddAccessOptsBuilder interface {
	ToAddAccessMap() (map[string]any, error)
}

type AccessOpts struct {
	// The UUID of the project to which access to the share type is granted.
	Project string `json:"project"`
}

// ToAddAccessMap assembles a request body based on the contents of a
// AccessOpts.
func (opts AccessOpts) ToAddAccessMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddAccess will add access to a ShareType based on the values
// in AccessOpts.
func AddAccess(ctx context.Context, client *gophercloud.ServiceClient, id string, opts AddAccessOptsBuilder) (r AddAccessResult) {
	_ = "STUB: not implemented"
	return *new(AddAccessResult)
}

// RemoveAccessOptsBuilder allows extensions to add additional parameters to the
// RemoveAccess
type RemoveAccessOptsBuilder interface {
	ToRemoveAccessMap() (map[string]any, error)
}

// ToRemoveAccessMap assembles a request body based on the contents of a
// AccessOpts.
func (opts AccessOpts) ToRemoveAccessMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveAccess will remove access to a ShareType based on the values
// in AccessOpts.
func RemoveAccess(ctx context.Context, client *gophercloud.ServiceClient, id string, opts RemoveAccessOptsBuilder) (r RemoveAccessResult) {
	_ = "STUB: not implemented"
	return *new(RemoveAccessResult)
}
