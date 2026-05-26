package availabilityzones

import (
	"time"

	"github.com/gophercloud/gophercloud/v2/pagination"
)

// AvailabilityZone contains all the information associated with an OpenStack
// AvailabilityZone.
type AvailabilityZone struct {
	// The availability zone ID.
	ID string `json:"id"`
	// The name of the availability zone.
	Name string `json:"name"`
	// The date and time stamp when the availability zone was created.
	CreatedAt time.Time `json:"-"`
	// The date and time stamp when the availability zone was updated.
	UpdatedAt time.Time `json:"-"`
}

// ListResult contains the response body and error from a List request.
type AvailabilityZonePage struct {
	pagination.SinglePageBase
}

// ExtractAvailabilityZones will get the AvailabilityZone objects out of the shareTypeAccessResult object.
func ExtractAvailabilityZones(r pagination.Page) ([]AvailabilityZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *AvailabilityZone) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
