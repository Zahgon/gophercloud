package mtu

import (
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
)

// ListOptsExt adds an MTU option to the base ListOpts.
type ListOptsExt struct {
	networks.ListOptsBuilder

	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	MTU int `q:"mtu"`
}

// ToNetworkListQuery adds the router:external option to the base network
// list options.
func (opts ListOptsExt) ToNetworkListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CreateOptsExt adds an MTU option to the base Network CreateOpts.
type CreateOptsExt struct {
	networks.CreateOptsBuilder

	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	MTU int `json:"mtu,omitempty"`
}

// ToNetworkCreateMap adds an MTU to the base network creation options.
func (opts CreateOptsExt) ToNetworkCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateOptsExt adds an MTU option to the base Network UpdateOpts.
type UpdateOptsExt struct {
	networks.UpdateOptsBuilder

	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	MTU int `json:"mtu,omitempty"`
}

// ToNetworkUpdateMap adds an MTU to the base network uptade options.
func (opts UpdateOptsExt) ToNetworkUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
