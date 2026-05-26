package portsbinding

import (
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
)

// CreateOptsExt adds port binding options to the base ports.CreateOpts.
type CreateOptsExt struct {
	// CreateOptsBuilder is the interface options structs have to satisfy in order
	// to be used in the main Create operation in this package.
	ports.CreateOptsBuilder

	// The ID of the host where the port is allocated
	HostID string `json:"binding:host_id,omitempty"`

	// The virtual network interface card (vNIC) type that is bound to the
	// neutron port.
	VNICType string `json:"binding:vnic_type,omitempty"`

	// A dictionary that enables the application running on the specified
	// host to pass and receive virtual network interface (VIF) port-specific
	// information to the plug-in.
	Profile map[string]any `json:"binding:profile,omitempty"`
}

// ToPortCreateMap casts a CreateOpts struct to a map.
func (opts CreateOptsExt) ToPortCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateOptsExt adds port binding options to the base ports.UpdateOpts
type UpdateOptsExt struct {
	// UpdateOptsBuilder is the interface options structs have to satisfy in order
	// to be used in the main Update operation in this package.
	ports.UpdateOptsBuilder

	// The ID of the host where the port is allocated.
	HostID *string `json:"binding:host_id,omitempty"`

	// The virtual network interface card (vNIC) type that is bound to the
	// neutron port.
	VNICType string `json:"binding:vnic_type,omitempty"`

	// A dictionary that enables the application running on the specified
	// host to pass and receive virtual network interface (VIF) port-specific
	// information to the plug-in.
	Profile map[string]any `json:"binding:profile,omitempty"`
}

// ToPortUpdateMap casts an UpdateOpts struct to a map.
func (opts UpdateOptsExt) ToPortUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// send null instead of the empty json object ("{}")
