package external

import (
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
)

// ListOptsExt adds the external network options to the base ListOpts.
type ListOptsExt struct {
	networks.ListOptsBuilder
	External *bool `q:"router:external"`
}

// ToNetworkListQuery adds the router:external option to the base network
// list options.
func (opts ListOptsExt) ToNetworkListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CreateOptsExt is the structure used when creating new external network
// resources. It embeds networks.CreateOpts and so inherits all of its required
// and optional fields, with the addition of the External field.
type CreateOptsExt struct {
	networks.CreateOptsBuilder
	External *bool `json:"router:external,omitempty"`
}

// ToNetworkCreateMap adds the router:external options to the base network
// creation options.
func (opts CreateOptsExt) ToNetworkCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateOptsExt is the structure used when updating existing external network
// resources. It embeds networks.UpdateOpts and so inherits all of its required
// and optional fields, with the addition of the External field.
type UpdateOptsExt struct {
	networks.UpdateOptsBuilder
	External *bool `json:"router:external,omitempty"`
}

// ToNetworkUpdateMap casts an UpdateOpts struct to a map.
func (opts UpdateOptsExt) ToNetworkUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
