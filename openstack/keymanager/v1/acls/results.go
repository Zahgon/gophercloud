package acls

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
)

// ACL represents an ACL on a resource.
type ACL map[string]ACLDetails

// ACLDetails represents the details of an ACL.
type ACLDetails struct {
	// Created is when the ACL was created.
	Created time.Time `json:"-"`

	// ProjectAccess denotes project-level access of the resource.
	ProjectAccess bool `json:"project-access"`

	// Updated is when the ACL was updated
	Updated time.Time `json:"-"`

	// Users are the UserIDs who have access to the resource.
	Users []string `json:"users"`
}

func (r *ACLDetails) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// ACLRef represents an ACL reference.
type ACLRef string

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as an ACL.
func (r commonResult) Extract() (*ACL, error) { _ = "STUB: not implemented"; return nil, nil }

// ACLResult is the response from a Get operation. Call its Extract method
// to interpret it as an ACL.
type ACLResult struct {
	commonResult
}

// ACLRefResult is the response from a Set or Update operation. Call its
// Extract method to interpret it as an ACLRef.
type ACLRefResult struct {
	gophercloud.Result
}

func (r ACLRefResult) Extract() (*ACLRef, error) { _ = "STUB: not implemented"; return nil, nil }

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
