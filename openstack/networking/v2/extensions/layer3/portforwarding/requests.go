package portforwarding

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type ListOptsBuilder interface {
	ToPortForwardingListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the port forwarding attributes you want to see returned. SortKey allows you to
// sort by a particular network attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	ID                string `q:"id"`
	Description       string `q:"description"`
	InternalPortID    string `q:"internal_port_id"`
	ExternalPort      string `q:"external_port"`
	InternalIPAddress string `q:"internal_ip_address"`
	Protocol          string `q:"protocol"`
	InternalPort      string `q:"internal_port"`
	SortKey           string `q:"sort_key"`
	SortDir           string `q:"sort_dir"`
	Fields            string `q:"fields"`
	Limit             int    `q:"limit"`
	Marker            string `q:"marker"`
}

// ToPortForwardingListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPortForwardingListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// Port Forwarding resources. It accepts a ListOpts struct, which allows you to
// filter and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder, id string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves a particular port forwarding resource based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, floatingIpId string, pfId string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOpts contains all the values needed to create a new port forwarding
// resource. All attributes are required.
type CreateOpts struct {
	Description       string `json:"description,omitempty"`
	InternalPortID    string `json:"internal_port_id"`
	InternalIPAddress string `json:"internal_ip_address"`
	InternalPort      int    `json:"internal_port"`
	ExternalPort      int    `json:"external_port"`
	Protocol          string `json:"protocol"`
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPortForwardingCreateMap() (map[string]any, error)
}

// ToPortForwardingCreateMap allows CreateOpts to satisfy the CreateOptsBuilder
// interface
func (opts CreateOpts) ToPortForwardingCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create accepts a CreateOpts struct and uses the values provided to create a
// new port forwarding for an existing floating IP.
func Create(ctx context.Context, c *gophercloud.ServiceClient, floatingIpId string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOpts contains the values used when updating a port forwarding resource.
type UpdateOpts struct {
	Description       *string `json:"description,omitempty"`
	InternalPortID    string  `json:"internal_port_id,omitempty"`
	InternalIPAddress string  `json:"internal_ip_address,omitempty"`
	InternalPort      int     `json:"internal_port,omitempty"`
	ExternalPort      int     `json:"external_port,omitempty"`
	Protocol          string  `json:"protocol,omitempty"`
}

// ToPortForwardingUpdateMap allows UpdateOpts to satisfy the UpdateOptsBuilder
// interface
func (opts UpdateOpts) ToPortForwardingUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPortForwardingUpdateMap() (map[string]any, error)
}

// Update allows port forwarding resources to be updated.
func Update(ctx context.Context, c *gophercloud.ServiceClient, fipID string, pfID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete will permanently delete a particular port forwarding for a given floating ID.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, floatingIpId string, pfId string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
