package volumetypes

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeTypeCreateMap() (map[string]any, error)
}

// CreateOpts contains options for creating a Volume Type. This object is passed to
// the volumetypes.Create function. For more information about these parameters,
// see the Volume Type object.
type CreateOpts struct {
	// The name of the volume type
	Name string `json:"name" required:"true"`
	// The volume type description
	Description string `json:"description,omitempty"`
	// the ID of the existing volume snapshot
	IsPublic *bool `json:"os-volume-type-access:is_public,omitempty"`
	// Extra spec key-value pairs defined by the user.
	ExtraSpecs map[string]string `json:"extra_specs,omitempty"`
}

// ToVolumeTypeCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToVolumeTypeCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create will create a new Volume Type based on the values in CreateOpts. To extract
// the Volume Type object from the response, call the Extract method on the
// CreateResult.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete will delete the existing Volume Type with the provided ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Get retrieves the Volume Type with the provided ID. To extract the Volume Type object
// from the response, call the Extract method on the GetResult.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeTypeListQuery() (string, error)
}

// ListOpts holds options for listing Volume Types. It is passed to the volumetypes.List
// function.
type ListOpts struct {
	// Name will filter by the specified volume type name.
	Name string `q:"name"`
	// Description will filter by the specified volume type description.
	Description string `q:"description"`
	// Specifies whether the query should include public or private Volume Types.
	// By default, it queries both types.
	IsPublic visibility `q:"is_public"`
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

type visibility string

const (
	// VisibilityDefault enables querying both public and private Volume Types.
	VisibilityDefault visibility = "None"
	// VisibilityPublic restricts the query to only public Volume Types.
	VisibilityPublic visibility = "true"
	// VisibilityPrivate restricts the query to only private Volume Types.
	VisibilityPrivate visibility = "false"
)

// ToVolumeTypeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeTypeListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns Volume types.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVolumeTypeUpdateMap() (map[string]any, error)
}

// UpdateOpts contain options for updating an existing Volume Type. This object is passed
// to the volumetypes.Update function. For more information about the parameters, see
// the Volume Type object.
type UpdateOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPublic    *bool   `json:"is_public,omitempty"`
}

// ToVolumeTypeUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToVolumeTypeUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update will update the Volume Type with provided information. To extract the updated
// Volume Type from the response, call the Extract method on the UpdateResult.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// ListExtraSpecs requests all the extra-specs for the given volume type ID.
func ListExtraSpecs(ctx context.Context, client *gophercloud.ServiceClient, volumeTypeID string) (r ListExtraSpecsResult) {
	_ = "STUB: not implemented"
	return *new(ListExtraSpecsResult)
}

// GetExtraSpec requests an extra-spec specified by key for the given volume type ID
func GetExtraSpec(ctx context.Context, client *gophercloud.ServiceClient, volumeTypeID string, key string) (r GetExtraSpecResult) {
	_ = "STUB: not implemented"
	return *new(GetExtraSpecResult)
}

// CreateExtraSpecsOptsBuilder allows extensions to add additional parameters to the
// CreateExtraSpecs requests.
type CreateExtraSpecsOptsBuilder interface {
	ToVolumeTypeExtraSpecsCreateMap() (map[string]any, error)
}

// ExtraSpecsOpts is a map that contains key-value pairs.
type ExtraSpecsOpts map[string]string

// ToVolumeTypeExtraSpecsCreateMap assembles a body for a Create request based on
// the contents of ExtraSpecsOpts.
func (opts ExtraSpecsOpts) ToVolumeTypeExtraSpecsCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateExtraSpecs will create or update the extra-specs key-value pairs for
// the specified volume type.
func CreateExtraSpecs(ctx context.Context, client *gophercloud.ServiceClient, volumeTypeID string, opts CreateExtraSpecsOptsBuilder) (r CreateExtraSpecsResult) {
	_ = "STUB: not implemented"
	return *new(CreateExtraSpecsResult)
}

// UpdateExtraSpecOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateExtraSpecOptsBuilder interface {
	ToVolumeTypeExtraSpecUpdateMap() (map[string]string, string, error)
}

// ToVolumeTypeExtraSpecUpdateMap assembles a body for an Update request based on
// the contents of a ExtraSpecOpts.
func (opts ExtraSpecsOpts) ToVolumeTypeExtraSpecUpdateMap() (map[string]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// UpdateExtraSpec will updates the value of the specified volume type's extra spec
// for the key in opts.
func UpdateExtraSpec(ctx context.Context, client *gophercloud.ServiceClient, volumeTypeID string, opts UpdateExtraSpecOptsBuilder) (r UpdateExtraSpecResult) {
	_ = "STUB: not implemented"
	return *new(UpdateExtraSpecResult)
}

// DeleteExtraSpec will delete the key-value pair with the given key for the given
// volume type ID.
func DeleteExtraSpec(ctx context.Context, client *gophercloud.ServiceClient, volumeTypeID, key string) (r DeleteExtraSpecResult) {
	_ = "STUB: not implemented"
	return *new(DeleteExtraSpecResult)
}

// ListAccesses retrieves the tenants which have access to a volume type.
func ListAccesses(client *gophercloud.ServiceClient, id string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// AddAccessOptsBuilder allows extensions to add additional parameters to the
// AddAccess requests.
type AddAccessOptsBuilder interface {
	ToVolumeTypeAddAccessMap() (map[string]any, error)
}

// AddAccessOpts represents options for adding access to a volume type.
type AddAccessOpts struct {
	// Project is the project/tenant ID to grant access.
	Project string `json:"project"`
}

// ToVolumeTypeAddAccessMap constructs a request body from AddAccessOpts.
func (opts AddAccessOpts) ToVolumeTypeAddAccessMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddAccess grants a tenant/project access to a volume type.
func AddAccess(ctx context.Context, client *gophercloud.ServiceClient, id string, opts AddAccessOptsBuilder) (r AddAccessResult) {
	_ = "STUB: not implemented"
	return *new(AddAccessResult)
}

// RemoveAccessOptsBuilder allows extensions to add additional parameters to the
// RemoveAccess requests.
type RemoveAccessOptsBuilder interface {
	ToVolumeTypeRemoveAccessMap() (map[string]any, error)
}

// RemoveAccessOpts represents options for removing access to a volume type.
type RemoveAccessOpts struct {
	// Project is the project/tenant ID to remove access.
	Project string `json:"project"`
}

// ToVolumeTypeRemoveAccessMap constructs a request body from RemoveAccessOpts.
func (opts RemoveAccessOpts) ToVolumeTypeRemoveAccessMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveAccess removes/revokes a tenant/project access to a volume type.
func RemoveAccess(ctx context.Context, client *gophercloud.ServiceClient, id string, opts RemoveAccessOptsBuilder) (r RemoveAccessResult) {
	_ = "STUB: not implemented"
	return *new(RemoveAccessResult)
}

// CreateEncryptionOptsBuilder allows extensions to add additional parameters to the
// Create Encryption request.
type CreateEncryptionOptsBuilder interface {
	ToEncryptionCreateMap() (map[string]any, error)
}

// CreateEncryptionOpts contains options for creating an Encryption Type object.
// This object is passed to the volumetypes.CreateEncryption function.
// For more information about these parameters,see the Encryption Type object.
type CreateEncryptionOpts struct {
	// The size of the encryption key.
	KeySize int `json:"key_size"`
	// The class of that provides the encryption support.
	Provider string `json:"provider" required:"true"`
	// Notional service where encryption is performed.
	ControlLocation string `json:"control_location"`
	// The encryption algorithm or mode.
	Cipher string `json:"cipher"`
}

// ToEncryptionCreateMap assembles a request body based on the contents of a
// CreateEncryptionOpts.
func (opts CreateEncryptionOpts) ToEncryptionCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateEncryption will creates an Encryption Type object based on the CreateEncryptionOpts.
// To extract the Encryption Type object from the response, call the Extract method on the
// EncryptionCreateResult.
func CreateEncryption(ctx context.Context, client *gophercloud.ServiceClient, id string, opts CreateEncryptionOptsBuilder) (r CreateEncryptionResult) {
	_ = "STUB: not implemented"
	return *new(CreateEncryptionResult)
}

// Delete will delete an encryption type for an existing Volume Type with the provided ID.
func DeleteEncryption(ctx context.Context, client *gophercloud.ServiceClient, id, encryptionID string) (r DeleteEncryptionResult) {
	_ = "STUB: not implemented"
	return *new(DeleteEncryptionResult)
}

// GetEncryption retrieves the encryption type for an existing VolumeType with the provided ID.
func GetEncryption(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetEncryptionResult) {
	_ = "STUB: not implemented"
	return *new(GetEncryptionResult)
}

// GetEncryptionSpecs retrieves the encryption type specs for an existing VolumeType with the provided ID.
func GetEncryptionSpec(ctx context.Context, client *gophercloud.ServiceClient, id, key string) (r GetEncryptionSpecResult) {
	_ = "STUB: not implemented"
	return *new(GetEncryptionSpecResult)
}

// UpdateEncryptionOptsBuilder allows extensions to add additional parameters to the
// Update encryption request.
type UpdateEncryptionOptsBuilder interface {
	ToUpdateEncryptionMap() (map[string]any, error)
}

// Update Encryption Opts contains options for creating an Update Encryption Type. This object is passed to
// the volumetypes.UpdateEncryption function. For more information about these parameters,
// see the Update Encryption Type object.
type UpdateEncryptionOpts struct {
	// The size of the encryption key.
	KeySize int `json:"key_size"`
	// The class of that provides the encryption support.
	Provider string `json:"provider"`
	// Notional service where encryption is performed.
	ControlLocation string `json:"control_location"`
	// The encryption algorithm or mode.
	Cipher string `json:"cipher"`
}

// ToEncryptionCreateMap assembles a request body based on the contents of a
// UpdateEncryptionOpts.
func (opts UpdateEncryptionOpts) ToUpdateEncryptionMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update will update an existing encryption for a Volume Type based on the values in UpdateEncryptionOpts.
// To extract the UpdateEncryption Type object from the response, call the Extract method on the
// UpdateEncryptionResult.
func UpdateEncryption(ctx context.Context, client *gophercloud.ServiceClient, id, encryptionID string, opts UpdateEncryptionOptsBuilder) (r UpdateEncryptionResult) {
	_ = "STUB: not implemented"
	return *new(UpdateEncryptionResult)
}
