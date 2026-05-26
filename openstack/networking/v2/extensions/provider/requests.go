package provider

import (
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
)

// CreateOptsExt adds a Segments option to the base Network CreateOpts.
type CreateOptsExt struct {
	networks.CreateOptsBuilder
	Segments []Segment `json:"segments,omitempty"`
}

// ToNetworkCreateMap adds segments to the base network creation options.
func (opts CreateOptsExt) ToNetworkCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateOptsExt adds a Segments option to the base Network UpdateOpts.
type UpdateOptsExt struct {
	networks.UpdateOptsBuilder
	Segments *[]Segment `json:"segments,omitempty"`
}

// ToNetworkUpdateMap adds segments to the base network update options.
func (opts UpdateOptsExt) ToNetworkUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
