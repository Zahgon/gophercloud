package servergroups

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// A ServerGroup creates a policy for instance placement in the cloud.
// You should use extract methods from microversions.go to retrieve additional
// fields.
type ServerGroup struct {
	// ID is the unique ID of the Server Group.
	ID string `json:"id"`

	// Name is the common name of the server group.
	Name string `json:"name"`

	// Polices are the group policies.
	//
	// Normally a single policy is applied:
	//
	// "affinity" will place all servers within the server group on the
	// same compute node.
	//
	// "anti-affinity" will place servers within the server group on different
	// compute nodes.
	Policies []string `json:"policies"`

	// Members are the members of the server group.
	Members []string `json:"members"`

	// UserID of the server group.
	UserID string `json:"user_id"`

	// ProjectID of the server group.
	ProjectID string `json:"project_id"`

	// Metadata includes a list of all user-specified key-value pairs attached
	// to the Server Group.
	Metadata map[string]any

	// Policy is the policy of a server group.
	// This requires microversion 2.64 or later.
	Policy *string `json:"policy"`

	// Rules are the rules of the server group.
	// This requires microversion 2.64 or later.
	Rules *Rules `json:"rules"`
}

// Rules represents set of rules for a policy.
// This requires microversion 2.64 or later.
type Rules struct {
	// MaxServerPerHost specifies how many servers can reside on a single compute host.
	// It can be used only with the "anti-affinity" policy.
	MaxServerPerHost int `json:"max_server_per_host"`
}

// ServerGroupPage stores a single page of all ServerGroups results from a
// List call.
type ServerGroupPage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a ServerGroupsPage is empty.
func (page ServerGroupPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractServerGroups interprets a page of results as a slice of
// ServerGroups.
func ExtractServerGroups(r pagination.Page) ([]ServerGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ServerGroupResult struct {
	gophercloud.Result
}

// Extract is a method that attempts to interpret any Server Group resource
// response as a ServerGroup struct.
func (r ServerGroupResult) Extract() (*ServerGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a ServerGroup.
type CreateResult struct {
	ServerGroupResult
}

// GetResult is the response from a Get operation. Call its Extract method to
// interpret it as a ServerGroup.
type GetResult struct {
	ServerGroupResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
