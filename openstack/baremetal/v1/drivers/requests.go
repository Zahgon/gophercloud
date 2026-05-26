package drivers

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListDriversOptsBuilder allows extensions to add additional parameters to the
// ListDrivers request.
type ListDriversOptsBuilder interface {
	ToListDriversOptsQuery() (string, error)
}

// ListDriversOpts defines query options that can be passed to ListDrivers
type ListDriversOpts struct {
	// Provide detailed information about the drivers
	Detail bool `q:"detail"`

	// Filter the list by the type of the driver
	Type string `q:"type"`
}

// ToListDriversOptsQuery formats a ListOpts into a query string
func (opts ListDriversOpts) ToListDriversOptsQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ListDrivers makes a request against the API to list all drivers
func ListDrivers(client *gophercloud.ServiceClient, opts ListDriversOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetDriverDetails Shows details for a driver
func GetDriverDetails(ctx context.Context, client *gophercloud.ServiceClient, driverName string) (r GetDriverResult) {
	_ = "STUB: not implemented"
	return *new(GetDriverResult)
}

// GetDriverProperties Shows the required and optional parameters that
// driverName expects to be supplied in the driver_info field for every
// Node it manages
func GetDriverProperties(ctx context.Context, client *gophercloud.ServiceClient, driverName string) (r GetPropertiesResult) {
	_ = "STUB: not implemented"
	return *new(GetPropertiesResult)
}

// GetDriverDiskProperties Show the required and optional parameters that
// driverName expects to be supplied in the node’s raid_config field, if a
// RAID configuration change is requested.
func GetDriverDiskProperties(ctx context.Context, client *gophercloud.ServiceClient, driverName string) (r GetDiskPropertiesResult) {
	_ = "STUB: not implemented"
	return *new(GetDiskPropertiesResult)
}
