package rules

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Rule represents a firewall rule
type Rule struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	Protocol             string   `json:"protocol"`
	Action               string   `json:"action"`
	IPVersion            int      `json:"ip_version,omitempty"`
	SourceIPAddress      string   `json:"source_ip_address,omitempty"`
	DestinationIPAddress string   `json:"destination_ip_address,omitempty"`
	SourcePort           string   `json:"source_port,omitempty"`
	DestinationPort      string   `json:"destination_port,omitempty"`
	Shared               bool     `json:"shared,omitempty"`
	Enabled              bool     `json:"enabled,omitempty"`
	FirewallPolicyID     []string `json:"firewall_policy_id"`
	TenantID             string   `json:"tenant_id"`
	ProjectID            string   `json:"project_id"`
}

// RulePage is the page returned by a pager when traversing over a
// collection of firewall rules.
type RulePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of firewall rules has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r RulePage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsEmpty checks whether a RulePage struct is empty.
func (r RulePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractRules accepts a Page struct, specifically a RouterPage struct,
// and extracts the elements into a slice of Router structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractRules(r pagination.Page) ([]Rule, error) { _ = "STUB: not implemented"; return nil, nil }

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a firewall rule.
func (r commonResult) Extract() (*Rule, error) { _ = "STUB: not implemented"; return nil, nil }

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// CreateResult represents the result of a create operation.
type CreateResult struct {
	commonResult
}
