package stackevents

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Event represents a stack event.
type Event struct {
	// The name of the resource for which the event occurred.
	ResourceName string `json:"resource_name"`
	// The time the event occurred.
	Time time.Time `json:"-"`
	// The URLs to the event.
	Links []gophercloud.Link `json:"links"`
	// The logical ID of the stack resource.
	LogicalResourceID string `json:"logical_resource_id"`
	// The reason of the status of the event.
	ResourceStatusReason string `json:"resource_status_reason"`
	// The status of the event.
	ResourceStatus string `json:"resource_status"`
	// The physical ID of the stack resource.
	PhysicalResourceID string `json:"physical_resource_id"`
	// The event ID.
	ID string `json:"id"`
	// Properties of the stack resource.
	ResourceProperties map[string]any `json:"resource_properties"`
}

func (r *Event) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// FindResult represents the result of a Find operation.
type FindResult struct {
	gophercloud.Result
}

// Extract returns a slice of Event objects and is called after a
// Find operation.
func (r FindResult) Extract() ([]Event, error) { _ = "STUB: not implemented"; return nil, nil }

// EventPage abstracts the raw results of making a List() request against the API.
// As OpenStack extensions may freely alter the response bodies of structures returned to the client, you may only safely access the
// data provided through the ExtractResources call.
type EventPage struct {
	pagination.MarkerPageBase
}

// IsEmpty returns true if a page contains no Server results.
func (r EventPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// LastMarker returns the last stack ID in a ListResult.
func (r EventPage) LastMarker() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ExtractEvents interprets the results of a single page from a List() call, producing a slice of Event entities.
func ExtractEvents(r pagination.Page) ([]Event, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractResourceEvents interprets the results of a single page from a
// ListResourceEvents() call, producing a slice of Event entities.
func ExtractResourceEvents(page pagination.Page) ([]Event, error) {
	_ = "STUB: not implemented"
	return nil,

		// GetResult represents the result of a Get operation.
		nil
}

type GetResult struct {
	gophercloud.Result
}

// Extract returns a pointer to an Event object and is called after a
// Get operation.
func (r GetResult) Extract() (*Event, error) { _ = "STUB: not implemented"; return nil, nil }
