package quotas

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get returns load balancer Quotas for a project.
func Get(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToQuotaUpdateMap() (map[string]any, error)
}

// UpdateOpts represents options used to update the load balancer Quotas.
type UpdateOpts struct {
	// Loadbalancer represents the number of load balancers. A "-1" value means no limit.
	Loadbalancer *int `json:"loadbalancer,omitempty"`

	// Listener represents the number of listeners. A "-1" value means no limit.
	Listener *int `json:"listener,omitempty"`

	// Member represents the number of members. A "-1" value means no limit.
	Member *int `json:"member,omitempty"`

	// Poool represents the number of pools. A "-1" value means no limit.
	Pool *int `json:"pool,omitempty"`

	// HealthMonitor represents the number of healthmonitors. A "-1" value means no limit.
	Healthmonitor *int `json:"healthmonitor,omitempty"`

	// L7Policy represents the number of l7policies. A "-1" value means no limit.
	L7Policy *int `json:"l7policy,omitempty"`

	// L7Rule represents the number of l7rules. A "-1" value means no limit.
	L7Rule *int `json:"l7rule,omitempty"`
}

// ToQuotaUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToQuotaUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update accepts a UpdateOpts struct and updates an existing load balancer Quotas using the
// values provided.
func Update(ctx context.Context, c *gophercloud.ServiceClient, projectID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// allow 200 (neutron/lbaasv2) and 202 (octavia)

func Delete(ctx context.Context, c *gophercloud.ServiceClient, projectID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
