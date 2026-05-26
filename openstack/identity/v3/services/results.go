package services

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type serviceResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a concrete
// Service. An error is returned if the original call or the extraction failed.
func (r serviceResult) Extract() (*Service, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateResult is the response from a Create request. Call its Extract method
// to interpret it as a Service.
type CreateResult struct {
	serviceResult
}

// GetResult is the response from a Get request. Call its Extract method
// to interpret it as a Service.
type GetResult struct {
	serviceResult
}

// UpdateResult is the response from an Update request. Call its Extract method
// to interpret it as a Service.
type UpdateResult struct {
	serviceResult
}

// DeleteResult is the response from a Delete request. Call its ExtractErr
// method to interpret it as a Service.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Service represents an OpenStack Service.
type Service struct {
	// ID is the unique ID of the service.
	ID string `json:"id"`

	// Name is the name of the service.
	Name string `json:"name"`

	// Description is the description of the service.
	Description string `json:"description"`

	// Type is the type of the service.
	Type string `json:"type"`

	// Enabled is whether or not the service is enabled.
	Enabled bool `json:"enabled"`

	// Links contains referencing links to the service.
	Links map[string]any `json:"links"`

	// Extra is a collection of miscellaneous key/values.
	Extra map[string]any `json:"-"`
}

func (r *Service) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// the following code is required for backward compatibility with the
// old behavior, when both description and name were in extra

// ServicePage is a single page of Service results.
type ServicePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the ServicePage contains no results.
func (p ServicePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r ServicePage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractServices extracts a slice of Services from a Collection acquired
// from List.
func ExtractServices(r pagination.Page) ([]Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
