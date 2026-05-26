package ruletypes

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListRuleTypes returns the list of rule types from the server
func ListRuleTypes(c *gophercloud.ServiceClient) (result pagination.Pager) {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetRuleType retrieves a specific QoS RuleType based on its name.
func GetRuleType(ctx context.Context, c *gophercloud.ServiceClient, name string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
