package services

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Service represents a Compute service in the OpenStack cloud.
type Service struct {
	// The binary name of the service.
	Binary string `json:"binary"`

	// The reason for disabling a service.
	DisabledReason string `json:"disabled_reason"`

	// Whether or not service was forced down manually.
	ForcedDown bool `json:"forced_down"`

	// The name of the host.
	Host string `json:"host"`

	// The id of the service.
	ID string `json:"-"`

	// The state of the service. One of up or down.
	State string `json:"state"`

	// The status of the service. One of enabled or disabled.
	Status string `json:"status"`

	// The date and time when the resource was updated.
	UpdatedAt time.Time `json:"-"`

	// The availability zone name.
	Zone string `json:"zone"`
}

// UnmarshalJSON to override default
func (r *Service) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// OpenStack Compute service returns ID in string representation since
// 2.53 microversion API (Pike release).

type serviceResult struct {
	gophercloud.Result
}

// Extract interprets any UpdateResult as a service, if possible.
func (r serviceResult) Extract() (*Service, error) { _ = "STUB: not implemented"; return nil, nil }

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a Server.
type UpdateResult struct {
	serviceResult
}

// ServicePage represents a single page of all Services from a List request.
type ServicePage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a page of Services contains any results.
func (page ServicePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func ExtractServices(r pagination.Page) ([]Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
