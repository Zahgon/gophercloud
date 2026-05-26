package instances

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	db "github.com/gophercloud/gophercloud/v2/openstack/db/v1/databases"
	"github.com/gophercloud/gophercloud/v2/openstack/db/v1/users"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder is the top-level interface for create options.
type CreateOptsBuilder interface {
	ToInstanceCreateMap() (map[string]any, error)
}

// DatastoreOpts represents the configuration for how an instance stores data.
type DatastoreOpts struct {
	Version string `json:"version"`
	Type    string `json:"type"`
}

// ToMap converts a DatastoreOpts to a map[string]string (for a request body)
func (opts DatastoreOpts) ToMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetworkOpts is used within CreateOpts to control a new server's network attachments.
type NetworkOpts struct {
	// UUID of a nova-network to attach to the newly provisioned server.
	// Required unless Port is provided.
	UUID string `json:"net-id,omitempty"`

	// Port of a neutron network to attach to the newly provisioned server.
	// Required unless UUID is provided.
	Port string `json:"port-id,omitempty"`

	// V4FixedIP [optional] specifies a fixed IPv4 address to be used on this network.
	V4FixedIP string `json:"v4-fixed-ip,omitempty"`

	// V6FixedIP [optional] specifies a fixed IPv6 address to be used on this network.
	V6FixedIP string `json:"v6-fixed-ip,omitempty"`
}

// ToMap converts a NetworkOpts to a map[string]string (for a request body)
func (opts NetworkOpts) ToMap() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateOpts is the struct responsible for configuring a new database instance.
type CreateOpts struct {
	// The availability zone of the instance.
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// ID of the configuration group that you want to attach to the instance.
	Configuration string `json:"configuration,omitempty"`
	// Either the integer UUID (in string form) of the flavor, or its URI
	// reference as specified in the response from the List() call. Required.
	FlavorRef string
	// Specifies the volume size in gigabytes (GB). The value must be between 1
	// and 300. Required.
	Size int
	// Specifies the volume type.
	VolumeType string
	// Name of the instance to create. The length of the name is limited to
	// 255 characters and any characters are permitted. Optional.
	Name string
	// A slice of database information options.
	Databases db.CreateOptsBuilder
	// A slice of user information options.
	Users users.CreateOptsBuilder
	// Options to configure the type of datastore the instance will use. This is
	// optional, and if excluded will default to MySQL.
	Datastore *DatastoreOpts
	// Networks dictates how this server will be attached to available networks.
	Networks []NetworkOpts
}

// ToInstanceCreateMap will render a JSON map.
func (opts CreateOpts) ToInstanceCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create asynchronously provisions a new database instance. It requires the
// user to specify a flavor and a volume size. The API service then provisions
// the instance with the requested flavor and sets up a volume of the specified
// size, which is the storage for the database instance.
//
// Although this call only allows the creation of 1 instance per request, you
// can create an instance with multiple databases and users. The default
// binding for a MySQL instance is port 3306.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// List retrieves the status and information for all database instances.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves the status and information for a specified database instance.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Delete permanently destroys the database instance.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// EnableRootUser enables the login from any host for the root user and
// provides the user with a generated root password.
func EnableRootUser(ctx context.Context, client *gophercloud.ServiceClient, id string) (r EnableRootUserResult) {
	_ = "STUB: not implemented"
	return *new(EnableRootUserResult)
}

// IsRootEnabled checks an instance to see if root access is enabled. It returns
// True if root user is enabled for the specified database instance or False
// otherwise.
func IsRootEnabled(ctx context.Context, client *gophercloud.ServiceClient, id string) (r IsRootEnabledResult) {
	_ = "STUB: not implemented"
	return *new(IsRootEnabledResult)
}

// Restart will restart only the MySQL Instance. Restarting MySQL will
// erase any dynamic configuration settings that you have made within MySQL.
// The MySQL service will be unavailable until the instance restarts.
func Restart(ctx context.Context, client *gophercloud.ServiceClient, id string) (r ActionResult) {
	_ = "STUB: not implemented"
	return *new(ActionResult)
}

// Resize changes the memory size of the instance, assuming a valid
// flavorRef is provided. It will also restart the MySQL service.
func Resize(ctx context.Context, client *gophercloud.ServiceClient, id, flavorRef string) (r ActionResult) {
	_ = "STUB: not implemented"
	return *new(ActionResult)
}

// ResizeVolume will resize the attached volume for an instance. It supports
// only increasing the volume size and does not support decreasing the size.
// The volume size is in gigabytes (GB) and must be an integer.
func ResizeVolume(ctx context.Context, client *gophercloud.ServiceClient, id string, size int) (r ActionResult) {
	_ = "STUB: not implemented"
	return *new(ActionResult)
}

// AttachConfigurationGroup will attach configuration group to the instance
func AttachConfigurationGroup(ctx context.Context, client *gophercloud.ServiceClient, instanceID string, configID string) (r ConfigurationResult) {
	_ = "STUB: not implemented"
	return *new(ConfigurationResult)
}

// DetachConfigurationGroup will dettach configuration group from the instance
func DetachConfigurationGroup(ctx context.Context, client *gophercloud.ServiceClient, instanceID string) (r ConfigurationResult) {
	_ = "STUB: not implemented"
	return *new(ConfigurationResult)
}
