package vlantransparent

import (
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
)

// ListOptsExt adds the vlan-transparent network options to the base ListOpts.
type ListOptsExt struct {
	networks.ListOptsBuilder
	VLANTransparent *bool `q:"vlan_transparent"`
}

// ToNetworkListQuery adds the vlan_transparent option to the base network
// list options.
func (opts ListOptsExt) ToNetworkListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CreateOptsExt is the structure used when creating new vlan-transparent
// network resources. It embeds networks.CreateOpts and so inherits all of its
// required and optional fields, with the addition of the VLANTransparent field.
type CreateOptsExt struct {
	networks.CreateOptsBuilder
	VLANTransparent *bool `json:"vlan_transparent,omitempty"`
}

// ToNetworkCreateMap adds the vlan_transparent option to the base network
// creation options.
func (opts CreateOptsExt) ToNetworkCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateOptsExt is the structure used when updating existing vlan-transparent
// network resources. It embeds networks.UpdateOpts and so inherits all of its
// required and optional fields, with the addition of the VLANTransparent field.
type UpdateOptsExt struct {
	networks.UpdateOptsBuilder
	VLANTransparent *bool `json:"vlan_transparent,omitempty"`
}

// ToNetworkUpdateMap casts an UpdateOpts struct to a map.
func (opts UpdateOptsExt) ToNetworkUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
