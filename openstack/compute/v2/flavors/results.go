package flavors

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// CreateResult is the response of a Get operations. Call its Extract method to
// interpret it as a Flavor.
type CreateResult struct {
	commonResult
}

// UpdateResult is the response of a Put operation. Call its Extract method to
// interpret it as a Flavor.
type UpdateResult struct {
	commonResult
}

// GetResult is the response of a Get operations. Call its Extract method to
// interpret it as a Flavor.
type GetResult struct {
	commonResult
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Extract provides access to the individual Flavor returned by the Get and
// Create functions.
func (r commonResult) Extract() (*Flavor, error) { _ = "STUB: not implemented"; return nil, nil }

// Flavor represent (virtual) hardware configurations for server resources
// in a region.
type Flavor struct {
	// ID is the flavor's unique ID.
	ID string `json:"id"`

	// Disk is the amount of root disk, measured in GB.
	Disk int `json:"disk"`

	// RAM is the amount of memory, measured in MB.
	RAM int `json:"ram"`

	// Name is the name of the flavor.
	Name string `json:"name"`

	// RxTxFactor describes bandwidth alterations of the flavor.
	RxTxFactor float64 `json:"rxtx_factor"`

	// Swap is the amount of swap space, measured in MB.
	Swap int `json:"-"`

	// VCPUs indicates how many (virtual) CPUs are available for this flavor.
	VCPUs int `json:"vcpus"`

	// IsPublic indicates whether the flavor is public.
	IsPublic bool `json:"os-flavor-access:is_public"`

	// Ephemeral is the amount of ephemeral disk space, measured in GB.
	Ephemeral int `json:"OS-FLV-EXT-DATA:ephemeral"`

	// Description is a free form description of the flavor. Limited to
	// 65535 characters in length. Only printable characters are allowed.
	// New in version 2.55
	Description string `json:"description"`

	// Properties is a dictionary of the flavor’s extra-specs key-and-value
	// pairs. This will only be included if the user is allowed by policy to
	// index flavor extra_specs
	// New in version 2.61
	ExtraSpecs map[string]string `json:"extra_specs"`
}

func (r *Flavor) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// FlavorPage contains a single page of all flavors from a ListDetails call.
type FlavorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines if a FlavorPage contains any results.
func (page FlavorPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (page FlavorPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractFlavors provides access to the list of flavors in a page acquired
// from the ListDetail operation.
func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccessPage contains a single page of all FlavorAccess entries for a flavor.
type AccessPage struct {
	pagination.SinglePageBase
}

// IsEmpty indicates whether an AccessPage is empty.
func (page AccessPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractAccesses interprets a page of results as a slice of FlavorAccess.
func ExtractAccesses(r pagination.Page) ([]FlavorAccess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type accessResult struct {
	gophercloud.Result
}

// AddAccessResult is the response of an AddAccess operation. Call its
// Extract method to interpret it as a slice of FlavorAccess.
type AddAccessResult struct {
	accessResult
}

// RemoveAccessResult is the response of a RemoveAccess operation. Call its
// Extract method to interpret it as a slice of FlavorAccess.
type RemoveAccessResult struct {
	accessResult
}

// Extract provides access to the result of an access create or delete.
// The result will be all accesses that the flavor has.
func (r accessResult) Extract() ([]FlavorAccess, error) { _ = "STUB: not implemented"; return nil, nil }

// FlavorAccess represents an ACL of tenant access to a specific Flavor.
type FlavorAccess struct {
	// FlavorID is the unique ID of the flavor.
	FlavorID string `json:"flavor_id"`

	// TenantID is the unique ID of the tenant.
	TenantID string `json:"tenant_id"`
}

// Extract interprets any extraSpecsResult as ExtraSpecs, if possible.
func (r extraSpecsResult) Extract() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
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

// CreateExtraSpecResult contains the result of a Create operation. Call its
// Extract method to interpret it as a map[string]interface.
type CreateExtraSpecsResult struct {
	extraSpecsResult
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
