package availabilityzoneprofiles

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type AvailabilityZoneProfile struct {
	// The unique ID for the AvailabilityZoneProfile
	ID string `json:"id"`

	// Human-readable name for the AvailabilityZoneProfile.
	// Does not have to be unique.
	Name string `json:"name"`

	// Name of the provider
	ProviderName string `json:"provider_name"`

	// Availability zone data
	AvailabilityZoneData string `json:"availability_zone_data"`
}

// AvailabilityZoneProfile is the page returned by a pager when traversing
// over a collection of profiles.
type AvailabilityZoneProfilePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of profiles has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r AvailabilityZoneProfilePage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsEmpty checks whether a AvailabilityZoneProfilePage struct is empty.
func (r AvailabilityZoneProfilePage) IsEmpty() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ExtractAvailabilityZoneProfiles accepts a Page struct, specifically a
// AvailabilityZoneProfilePage struct, and extracts the elements into a slice
// of Flavor structs. In other words, a generic collection is mapped into a
// relevant slice.
func ExtractAvailabilityZoneProfiles(r pagination.Page) ([]AvailabilityZoneProfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a flavor.
func (r commonResult) Extract() (*AvailabilityZoneProfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Flavor.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Flavor.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Flavor.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
