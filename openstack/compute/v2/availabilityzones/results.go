package availabilityzones

import (
	"time"

	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ServiceState represents the state of a service in an AvailabilityZone.
type ServiceState struct {
	Active    bool      `json:"active"`
	Available bool      `json:"available"`
	UpdatedAt time.Time `json:"-"`
}

// UnmarshalJSON to override default
func (r *ServiceState) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Services is a map of services contained in an AvailabilityZone.
type Services map[string]ServiceState

// Hosts is map of hosts/nodes contained in an AvailabilityZone.
// Each host can have multiple services.
type Hosts map[string]Services

// ZoneState represents the current state of the availability zone.
type ZoneState struct {
	// Returns true if the availability zone is available
	Available bool `json:"available"`
}

// AvailabilityZone contains all the information associated with an OpenStack
// AvailabilityZone.
type AvailabilityZone struct {
	Hosts Hosts `json:"hosts"`
	// The availability zone name
	ZoneName  string    `json:"zoneName"`
	ZoneState ZoneState `json:"zoneState"`
}

type AvailabilityZonePage struct {
	pagination.SinglePageBase
}

// ExtractAvailabilityZones returns a slice of AvailabilityZones contained in a
// single page of results.
func ExtractAvailabilityZones(r pagination.Page) ([]AvailabilityZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
