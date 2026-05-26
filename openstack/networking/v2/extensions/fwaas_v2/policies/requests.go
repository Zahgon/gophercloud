package policies

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the firewall policy attributes you want to see returned. SortKey allows you
// to sort by a particular firewall policy attribute. SortDir sets the direction,
// and is either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	TenantID    string `q:"tenant_id"`
	ProjectID   string `q:"project_id"`
	Name        string `q:"name"`
	Description string `q:"description"`
	Shared      *bool  `q:"shared"`
	Audited     *bool  `q:"audited"`
	ID          string `q:"id"`
	Limit       int    `q:"limit"`
	Marker      string `q:"marker"`
	SortKey     string `q:"sort_key"`
	SortDir     string `q:"sort_dir"`
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List returns a Pager which allows you to iterate over a collection of
// firewall policies. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
//
// Default policy settings return only those firewall policies that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Create operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type CreateOptsBuilder interface {
	ToFirewallPolicyCreateMap() (map[string]any, error)
}

// CreateOpts contains all the values needed to create a new firewall policy.
type CreateOpts struct {
	// Only required if the caller has an admin role and wants to create a firewall policy
	// for another tenant.
	TenantID      string   `json:"tenant_id,omitempty"`
	ProjectID     string   `json:"project_id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	Shared        *bool    `json:"shared,omitempty"`
	Audited       *bool    `json:"audited,omitempty"`
	FirewallRules []string `json:"firewall_rules,omitempty"`
}

// ToFirewallPolicyCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToFirewallPolicyCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create accepts a CreateOpts struct and uses the values to create a new firewall policy
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get retrieves a particular firewall policy based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Update operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type UpdateOptsBuilder interface {
	ToFirewallPolicyUpdateMap() (map[string]any, error)
}

// UpdateOpts contains the values used when updating a firewall policy.
type UpdateOpts struct {
	Name          *string   `json:"name,omitempty"`
	Description   *string   `json:"description,omitempty"`
	Shared        *bool     `json:"shared,omitempty"`
	Audited       *bool     `json:"audited,omitempty"`
	FirewallRules *[]string `json:"firewall_rules,omitempty"`
}

// ToFirewallPolicyUpdateMap casts a CreateOpts struct to a map.
func (opts UpdateOpts) ToFirewallPolicyUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update allows firewall policies to be updated.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete will permanently delete a particular firewall policy based on its unique ID.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

type InsertRuleOptsBuilder interface {
	ToFirewallPolicyInsertRuleMap() (map[string]any, error)
}

type InsertRuleOpts struct {
	ID           string `json:"firewall_rule_id" required:"true"`
	InsertBefore string `json:"insert_before,omitempty" xor:"InsertAfter"`
	InsertAfter  string `json:"insert_after,omitempty" xor:"InsertBefore"`
}

func (opts InsertRuleOpts) ToFirewallPolicyInsertRuleMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InsertRule(ctx context.Context, c *gophercloud.ServiceClient, id string, opts InsertRuleOptsBuilder) (r InsertRuleResult) {
	_ = "STUB: not implemented"
	return *new(InsertRuleResult)
}

func RemoveRule(ctx context.Context, c *gophercloud.ServiceClient, id, ruleID string) (r RemoveRuleResult) {
	_ = "STUB: not implemented"
	return *new(RemoveRuleResult)
}
