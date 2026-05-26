package volumetypes

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// VolumeType contains all the information associated with an OpenStack Volume Type.
type VolumeType struct {
	// Unique identifier for the volume type.
	ID string `json:"id"`
	// Human-readable display name for the volume type.
	Name string `json:"name"`
	// Human-readable description for the volume type.
	Description string `json:"description"`
	// Arbitrary key-value pairs defined by the user.
	ExtraSpecs map[string]string `json:"extra_specs"`
	// Whether the volume type is publicly visible.
	IsPublic bool `json:"is_public"`
	// Qos Spec ID
	QosSpecID string `json:"qos_specs_id"`
	// Volume Type access public attribute
	PublicAccess bool `json:"os-volume-type-access:is_public"`
}

// VolumeTypePage is a pagination.pager that is returned from a call to the List function.
type VolumeTypePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no Volume Types.
func (r VolumeTypePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (page VolumeTypePage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractVolumeTypes extracts and returns Volumes. It is used while iterating over a volumetypes.List call.
func ExtractVolumeTypes(r pagination.Page) ([]VolumeType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Volume Type object out of the commonResult object.
func (r commonResult) Extract() (*VolumeType, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractInto converts our response data into a volume type struct
func (r commonResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

// ExtractVolumeTypesInto similar to ExtractInto but operates on a `list` of volume types
func ExtractVolumeTypesInto(r pagination.Page, v any) error { _ = "STUB: not implemented"; return nil }

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

// extraSpecsResult contains the result of a call for (potentially) multiple
// key-value pairs. Call its Extract method to interpret it as a
// map[string]interface.
type extraSpecsResult struct {
	gophercloud.Result
}

// ListExtraSpecsResult contains the result of a Get operation. Call its Extract
// method to interpret it as a map[string]interface.
type ListExtraSpecsResult struct {
	extraSpecsResult
}

// CreateExtraSpecsResult contains the result of a Create operation. Call its
// Extract method to interpret it as a map[string]interface.
type CreateExtraSpecsResult struct {
	extraSpecsResult
}

// Extract interprets any extraSpecsResult as ExtraSpecs, if possible.
func (r extraSpecsResult) Extract() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extraSpecResult contains the result of a call for individual a single
// key-value pair.
type extraSpecResult struct {
	gophercloud.Result
}

// GetExtraSpecResult contains the result of a Get operation. Call its Extract
// method to interpret it as a map[string]interface.
type GetExtraSpecResult struct {
	extraSpecResult
}

// UpdateExtraSpecResult contains the result of an Update operation. Call its
// Extract method to interpret it as a map[string]interface.
type UpdateExtraSpecResult struct {
	extraSpecResult
}

// DeleteExtraSpecResult contains the result of a Delete operation. Call its
// ExtractErr method to determine if the call succeeded or failed.
type DeleteExtraSpecResult struct {
	gophercloud.ErrResult
}

// Extract interprets any extraSpecResult as an ExtraSpec, if possible.
func (r extraSpecResult) Extract() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VolumeTypeAccess represents an ACL of project access to a specific Volume Type.
type VolumeTypeAccess struct {
	// VolumeTypeID is the unique ID of the volume type.
	VolumeTypeID string `json:"volume_type_id"`

	// ProjectID is the unique ID of the project.
	ProjectID string `json:"project_id"`
}

// AccessPage contains a single page of all VolumeTypeAccess entries for a volume type.
type AccessPage struct {
	pagination.SinglePageBase
}

// IsEmpty indicates whether an AccessPage is empty.
func (page AccessPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractAccesses interprets a page of results as a slice of VolumeTypeAccess.
func ExtractAccesses(r pagination.Page) ([]VolumeTypeAccess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddAccessResult is the response from a AddAccess request. Call its
// ExtractErr method to determine if the request succeeded or failed.
type AddAccessResult struct {
	gophercloud.ErrResult
}

// RemoveAccessResult is the response from a RemoveAccess request. Call its
// ExtractErr method to determine if the request succeeded or failed.
type RemoveAccessResult struct {
	gophercloud.ErrResult
}

type EncryptionType struct {
	// Unique identifier for the volume type.
	VolumeTypeID string `json:"volume_type_id"`
	// Notional service where encryption is performed.
	ControlLocation string `json:"control_location"`
	// Unique identifier for encryption type.
	EncryptionID string `json:"encryption_id"`
	// Size of encryption key.
	KeySize int `json:"key_size"`
	// Class that provides encryption support.
	Provider string `json:"provider"`
	// The encryption algorithm or mode.
	Cipher string `json:"cipher"`
}

type encryptionResult struct {
	gophercloud.Result
}

func (r encryptionResult) Extract() (*EncryptionType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractInto converts our response data into a volume type struct
func (r encryptionResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

type CreateEncryptionResult struct {
	encryptionResult
}

// UpdateResult contains the response body and error from an UpdateEncryption request.
type UpdateEncryptionResult struct {
	encryptionResult
}

// DeleteEncryptionResult contains the response body and error from a DeleteEncryprion request.
type DeleteEncryptionResult struct {
	gophercloud.ErrResult
}

type GetEncryptionType struct {
	// Unique identifier for the volume type.
	VolumeTypeID string `json:"volume_type_id"`
	// Notional service where encryption is performed.
	ControlLocation string `json:"control_location"`
	// Shows if the resource is deleted or Notional
	Deleted bool `json:"deleted"`
	// Shows the date and time the resource was created.
	CreatedAt string `json:"created_at"`
	// Shows the date and time when resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Unique identifier for encryption type.
	EncryptionID string `json:"encryption_id"`
	// Size of encryption key.
	KeySize int `json:"key_size"`
	// Class that provides encryption support.
	Provider string `json:"provider"`
	// Shows the date and time the reousrce was deleted.
	DeletedAt string `json:"deleted_at"`
	// The encryption algorithm or mode.
	Cipher string `json:"cipher"`
}

type encryptionShowResult struct {
	gophercloud.Result
}

// Extract interprets any extraSpecResult as an ExtraSpec, if possible.
func (r encryptionShowResult) Extract() (*GetEncryptionType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GetEncryptionResult struct {
	encryptionShowResult
}

type encryptionShowSpecResult struct {
	gophercloud.Result
}

// Extract interprets any empty interface Result as an empty interface.
func (r encryptionShowSpecResult) Extract() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GetEncryptionSpecResult struct {
	encryptionShowSpecResult
}
