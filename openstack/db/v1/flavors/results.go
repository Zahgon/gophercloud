package flavors

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// GetResult temporarily holds the response from a Get call.
type GetResult struct {
	gophercloud.Result
}

// Extract provides access to the individual Flavor returned by the Get function.
func (r GetResult) Extract() (*Flavor, error) { _ = "STUB: not implemented"; return nil, nil }

// Flavor records represent (virtual) hardware configurations for server resources in a region.
type Flavor struct {
	// The flavor's unique identifier.
	// Contains 0 if the ID is not an integer.
	ID int `json:"id"`

	// The RAM capacity for the flavor.
	RAM int `json:"ram"`

	// The Name field provides a human-readable moniker for the flavor.
	Name string `json:"name"`

	// Links to access the flavor.
	Links []gophercloud.Link

	// The flavor's unique identifier as a string
	StrID string `json:"str_id"`
}

// FlavorPage contains a single page of the response from a List call.
type FlavorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines if a page contains any results.
func (page FlavorPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (page FlavorPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractFlavors provides access to the list of flavors in a page acquired from the List operation.
func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
