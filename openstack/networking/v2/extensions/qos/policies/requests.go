package policies

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// PortCreateOptsExt adds QoS options to the base ports.CreateOpts.
type PortCreateOptsExt struct {
	ports.CreateOptsBuilder

	// QoSPolicyID represents an associated QoS policy.
	QoSPolicyID string `json:"qos_policy_id,omitempty"`
}

// ToPortCreateMap casts a CreateOpts struct to a map.
func (opts PortCreateOptsExt) ToPortCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PortUpdateOptsExt adds QoS options to the base ports.UpdateOpts.
type PortUpdateOptsExt struct {
	ports.UpdateOptsBuilder

	// QoSPolicyID represents an associated QoS policy.
	// Setting it to a pointer of an empty string will remove associated QoS policy from port.
	QoSPolicyID *string `json:"qos_policy_id,omitempty"`
}

// ToPortUpdateMap casts a UpdateOpts struct to a map.
func (opts PortUpdateOptsExt) ToPortUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetworkCreateOptsExt adds QoS options to the base networks.CreateOpts.
type NetworkCreateOptsExt struct {
	networks.CreateOptsBuilder

	// QoSPolicyID represents an associated QoS policy.
	QoSPolicyID string `json:"qos_policy_id,omitempty"`
}

// ToNetworkCreateMap casts a CreateOpts struct to a map.
func (opts NetworkCreateOptsExt) ToNetworkCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetworkUpdateOptsExt adds QoS options to the base networks.UpdateOpts.
type NetworkUpdateOptsExt struct {
	networks.UpdateOptsBuilder

	// QoSPolicyID represents an associated QoS policy.
	// Setting it to a pointer of an empty string will remove associated QoS policy from network.
	QoSPolicyID *string `json:"qos_policy_id,omitempty"`
}

// ToNetworkUpdateMap casts a UpdateOpts struct to a map.
func (opts NetworkUpdateOptsExt) ToNetworkUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PolicyListOptsBuilder allows extensions to add additional parameters to the List request.
type PolicyListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the Neutron API. Filtering is achieved by passing in struct field values
// that map to the Policy attributes you want to see returned.
// SortKey allows you to sort by a particular Policy attribute.
// SortDir sets the direction, and is either `asc' or `desc'.
// Marker and Limit are used for the pagination.
type ListOpts struct {
	ID             string `q:"id"`
	TenantID       string `q:"tenant_id"`
	ProjectID      string `q:"project_id"`
	Name           string `q:"name"`
	Description    string `q:"description"`
	IsDefault      *bool  `q:"is_default"`
	Shared         *bool  `q:"shared"`
	Limit          int    `q:"limit"`
	Marker         string `q:"marker"`
	SortKey        string `q:"sort_key"`
	SortDir        string `q:"sort_dir"`
	Tags           string `q:"tags"`
	TagsAny        string `q:"tags-any"`
	NotTags        string `q:"not-tags"`
	NotTagsAny     string `q:"not-tags-any"`
	RevisionNumber *int   `q:"revision_number"`
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List returns a Pager which allows you to iterate over a collection of
// Policy. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts PolicyListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves a specific QoS policy based on its ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]any, error)
}

// CreateOpts specifies parameters of a new QoS policy.
type CreateOpts struct {
	// Name is the human-readable name of the QoS policy.
	Name string `json:"name"`

	// TenantID is the id of the Identity project.
	TenantID string `json:"tenant_id,omitempty"`

	// ProjectID is the id of the Identity project.
	ProjectID string `json:"project_id,omitempty"`

	// Shared indicates whether this QoS policy is shared across all projects.
	Shared bool `json:"shared,omitempty"`

	// Description is the human-readable description for the QoS policy.
	Description string `json:"description,omitempty"`

	// IsDefault indicates if this QoS policy is default policy or not.
	IsDefault bool `json:"is_default,omitempty"`
}

// ToPolicyCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToPolicyCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new QoS policy on the server.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]any, error)
}

// UpdateOpts represents options used to update a QoS policy.
type UpdateOpts struct {
	// Name is the human-readable name of the QoS policy.
	Name string `json:"name,omitempty"`

	// Shared indicates whether this QoS policy is shared across all projects.
	Shared *bool `json:"shared,omitempty"`

	// Description is the human-readable description for the QoS policy.
	Description *string `json:"description,omitempty"`

	// IsDefault indicates if this QoS policy is default policy or not.
	IsDefault *bool `json:"is_default,omitempty"`

	// RevisionNumber implements extension:standard-attr-revisions. If != "" it
	// will set revision_number=%s. If the revision number does not match, the
	// update will fail.
	RevisionNumber *int `json:"-" h:"If-Match"`
}

// ToPolicyUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update accepts a UpdateOpts struct and updates an existing policy using the
// values provided.
func Update(ctx context.Context, c *gophercloud.ServiceClient, policyID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete accepts a unique ID and deletes the QoS policy associated with it.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
