package extraroutes

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/routers"
)

// OptsBuilder allows extensions to add additional parameters to the Add or
// Remove requests.
type OptsBuilder interface {
	ToExtraRoutesUpdateMap() (map[string]any, error)
}

// Opts contains the values needed to add or remove a list og routes on a
// router.
type Opts struct {
	Routes *[]routers.Route `json:"routes,omitempty"`
}

// ToExtraRoutesUpdateMap builds a body based on Opts.
func (opts Opts) ToExtraRoutesUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add allows routers to be updated with a list of routes to be added.
func Add(ctx context.Context, c *gophercloud.ServiceClient, id string, opts OptsBuilder) (r AddResult) {
	_ = "STUB: not implemented"
	return *new(AddResult)
}

// Remove allows routers to be updated with a list of routes to be removed.
func Remove(ctx context.Context, c *gophercloud.ServiceClient, id string, opts OptsBuilder) (r RemoveResult) {
	_ = "STUB: not implemented"
	return *new(RemoveResult)
}
