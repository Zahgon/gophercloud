package conductors

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type conductorResult struct {
	gophercloud.Result
}

// Extract interprets any conductorResult as a Conductor, if possible.
func (r conductorResult) Extract() (*Conductor, error) { _ = "STUB: not implemented"; return nil, nil }

func (r conductorResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

func ExtractConductorInto(r pagination.Page, v any) error { _ = "STUB: not implemented"; return nil }

// Conductor represents a conductor in the OpenStack Bare Metal API.
type Conductor struct {
	// Whether or not this Conductor is alive or not
	Alive bool `json:"alive"`

	// Hostname of this conductor
	Hostname string `json:"hostname"`

	// Array of drivers for this conductor.
	Drivers []string `json:"drivers"`

	// Conductor group for a conductor. Case-insensitive string up to 255 characters, containing a-z, 0-9, _, -, and ..
	ConductorGroup string `json:"conductor_group"`

	// The UTC date and time when the resource was created, ISO 8601 format.
	CreatedAt time.Time `json:"created_at"`

	// The UTC date and time when the resource was updated, ISO 8601 format. May be “null”.
	UpdatedAt time.Time `json:"updated_at"`
}

// ConductorPage abstracts the raw results of making a List() request against
// the API. As OpenStack extensions may freely alter the response bodies of
// structures returned to the client, you may only safely access the data
// provided through the ExtractConductor call.
type ConductorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a page contains no conductor results.
func (r ConductorPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r ConductorPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractConductors interprets the results of a single page from a List() call,
// producing a slice of Conductor entities.
func ExtractConductors(r pagination.Page) ([]Conductor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetResult is the response from a Get operation. Call its Extract
// method to interpret it as a Conductor.
type GetResult struct {
	conductorResult
}
