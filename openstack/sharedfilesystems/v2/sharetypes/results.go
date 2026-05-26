package sharetypes

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ShareType contains all the information associated with an OpenStack
// ShareType.
type ShareType struct {
	// The Share Type ID
	ID string `json:"id"`
	// The Share Type name
	Name string `json:"name"`
	// Indicates whether a share type is publicly accessible
	IsPublic bool `json:"os-share-type-access:is_public"`
	// The required extra specifications for the share type
	RequiredExtraSpecs map[string]any `json:"required_extra_specs"`
	// The extra specifications for the share type
	ExtraSpecs map[string]any `json:"extra_specs"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract will get the ShareType object out of the commonResult object.
func (r commonResult) Extract() (*ShareType, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ShareTypePage is a pagination.pager that is returned from a call to the List function.
type ShareTypePage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no ShareTypes.
func (r ShareTypePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractShareTypes extracts and returns ShareTypes. It is used while
// iterating over a sharetypes.List call.
func ExtractShareTypes(r pagination.Page) ([]ShareType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDefaultResult contains the response body and error from a Get Default request.
type GetDefaultResult struct {
	commonResult
}

// ExtraSpecs contains all the information associated with extra specifications
// for an Openstack ShareType.
type ExtraSpecs map[string]any

type extraSpecsResult struct {
	gophercloud.Result
}

// Extract will get the ExtraSpecs object out of the commonResult object.
func (r extraSpecsResult) Extract() (ExtraSpecs, error) {
	_ = "STUB: not implemented"
	return *new(ExtraSpecs), nil
}

// GetExtraSpecsResult contains the response body and error from a Get Extra Specs request.
type GetExtraSpecsResult struct {
	extraSpecsResult
}

// SetExtraSpecsResult contains the response body and error from a Set Extra Specs request.
type SetExtraSpecsResult struct {
	extraSpecsResult
}

// UnsetExtraSpecsResult contains the response body and error from a Unset Extra Specs request.
type UnsetExtraSpecsResult struct {
	gophercloud.ErrResult
}

// ShareTypeAccess contains all the information associated with an OpenStack
// ShareTypeAccess.
type ShareTypeAccess struct {
	// The share type ID of the member.
	ShareTypeID string `json:"share_type_id"`
	// The UUID of the project for which access to the share type is granted.
	ProjectID string `json:"project_id"`
}

type shareTypeAccessResult struct {
	gophercloud.Result
}

// ShowAccessResult contains the response body and error from a Show access request.
type ShowAccessResult struct {
	shareTypeAccessResult
}

// Extract will get the ShareTypeAccess objects out of the shareTypeAccessResult object.
func (r ShowAccessResult) Extract() ([]ShareTypeAccess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddAccessResult contains the response body and error from a Add Access request.
type AddAccessResult struct {
	gophercloud.ErrResult
}

// RemoveAccessResult contains the response body and error from a Remove Access request.
type RemoveAccessResult struct {
	gophercloud.ErrResult
}
