package instanceactions

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// InstanceAction represents an instance action.
type InstanceAction struct {
	// Action is the name of the action.
	Action string `json:"action"`

	// InstanceUUID is the UUID of the instance.
	InstanceUUID string `json:"instance_uuid"`

	// Message is the related error message for when an action fails.
	Message string `json:"message"`

	// Project ID is the ID of the project which initiated the action.
	ProjectID string `json:"project_id"`

	// RequestID is the ID generated when performing the action.
	RequestID string `json:"request_id"`

	// StartTime is the time the action started.
	StartTime time.Time `json:"-"`

	// UserID is the ID of the user which initiated the action.
	UserID string `json:"user_id"`
}

// UnmarshalJSON converts our JSON API response into our instance action struct
func (i *InstanceAction) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// InstanceActionPage abstracts the raw results of making a List() request
// against the API. As OpenStack extensions may freely alter the response bodies
// of structures returned to the client, you may only safely access the data
// provided through the ExtractInstanceActions call.
type InstanceActionPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if an InstanceActionPage contains no instance actions.
func (r InstanceActionPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractInstanceActions interprets a page of results as a slice
// of InstanceAction.
func ExtractInstanceActions(r pagination.Page) ([]InstanceAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Event represents an event of instance action.
type Event struct {
	// Event is the name of the event.
	Event string `json:"event"`

	// Host is the host of the event.
	// This requires microversion 2.62 or later.
	Host *string `json:"host"`

	// HostID is the host id of the event.
	// This requires microversion 2.62 or later.
	HostID *string `json:"hostId"`

	// Result is the result of the event.
	Result string `json:"result"`

	// Traceback is the traceback stack if an error occurred.
	Traceback string `json:"traceback"`

	// StartTime is the time the action started.
	StartTime time.Time `json:"-"`

	// FinishTime is the time the event finished.
	FinishTime time.Time `json:"-"`
}

// UnmarshalJSON converts our JSON API response into our instance action struct.
func (e *Event) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// InstanceActionDetail represents the details of an Action.
type InstanceActionDetail struct {
	// Action is the name of the Action.
	Action string `json:"action"`

	// InstanceUUID is the UUID of the instance.
	InstanceUUID string `json:"instance_uuid"`

	// Message is the related error message for when an action fails.
	Message string `json:"message"`

	// Project ID is the ID of the project which initiated the action.
	ProjectID string `json:"project_id"`

	// RequestID is the ID generated when performing the action.
	RequestID string `json:"request_id"`

	// UserID is the ID of the user which initiated the action.
	UserID string `json:"user_id"`

	// Events is the list of events of the action.
	// This requires microversion 2.50 or later.
	Events *[]Event `json:"events"`

	// UpdatedAt last update date of the action.
	// This requires microversion 2.58 or later.
	UpdatedAt *time.Time `json:"-"`

	// StartTime is the time the action started.
	StartTime time.Time `json:"-"`
}

// UnmarshalJSON converts our JSON API response into our instance action struct
func (i *InstanceActionDetail) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// InstanceActionResult is the result handler of Get.
type InstanceActionResult struct {
	gophercloud.Result
}

// Extract interprets a result as an InstanceActionDetail.
func (r InstanceActionResult) Extract() (InstanceActionDetail, error) {
	_ = "STUB: not implemented"
	return *new(InstanceActionDetail), nil
}

func (r InstanceActionResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

func ExtractInstanceActionsInto(r pagination.Page, v any) error {
	_ = "STUB: not implemented"
	return nil
}
