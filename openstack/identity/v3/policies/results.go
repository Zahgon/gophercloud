package policies

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Policy is an arbitrarily serialized policy engine rule
// set to be consumed by a remote service.
type Policy struct {
	// ID is the unique ID of the policy.
	ID string `json:"id"`

	// Blob is the policy rule as a serialized blob.
	Blob string `json:"blob"`

	// Type is the MIME media type of the serialized policy blob.
	Type string `json:"type"`

	// Links contains referencing links to the policy.
	Links map[string]any `json:"links"`

	// Extra is a collection of miscellaneous key/values.
	Extra map[string]any `json:"-"`
}

func (r *Policy) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Collect other fields and bundle them into Extra
// but only if a field titled "extra" wasn't sent.

type policyResult struct {
	gophercloud.Result
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Policy
type CreateResult struct {
	policyResult
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Policy.
type GetResult struct {
	policyResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a Policy.
type UpdateResult struct {
	policyResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// PolicyPage is a single page of Policy results.
type PolicyPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Policies contains any results.
func (r PolicyPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r PolicyPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractPolicies returns a slice of Policies
// contained in a single page of results.
func ExtractPolicies(r pagination.Page) ([]Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract interprets any policyResults as a Policy.
func (r policyResult) Extract() (*Policy, error) { _ = "STUB: not implemented"; return nil, nil }
