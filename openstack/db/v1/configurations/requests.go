package configurations

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List will list all of the available configurations.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder is a top-level interface which renders a JSON map.
type CreateOptsBuilder interface {
	ToConfigCreateMap() (map[string]any, error)
}

// DatastoreOpts is the primary options struct for creating and modifying
// how configuration resources are associated with datastores.
type DatastoreOpts struct {
	// The type of datastore. Defaults to "MySQL".
	Type string `json:"type,omitempty"`
	// The specific version of a datastore. Defaults to "5.6".
	Version string `json:"version,omitempty"`
}

// CreateOpts is the struct responsible for configuring new configurations.
type CreateOpts struct {
	// The configuration group name
	Name string `json:"name" required:"true"`
	// A map of user-defined configuration settings that will define
	// how each associated datastore works. Each key/value pair is specific to a
	// datastore type.
	Values map[string]any `json:"values" required:"true"`
	// Associates the configuration group with a particular datastore.
	Datastore *DatastoreOpts `json:"datastore,omitempty"`
	// A human-readable explanation for the group.
	Description string `json:"description,omitempty"`
}

// ToConfigCreateMap casts a CreateOpts struct into a JSON map.
func (opts CreateOpts) ToConfigCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create will create a new configuration group.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get will retrieve the details for a specified configuration group.
func Get(ctx context.Context, client *gophercloud.ServiceClient, configID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder is the top-level interface for casting update options into
// JSON maps.
type UpdateOptsBuilder interface {
	ToConfigUpdateMap() (map[string]any, error)
}

// UpdateOpts is the struct responsible for modifying existing configurations.
type UpdateOpts struct {
	// The configuration group name
	Name string `json:"name,omitempty"`
	// A map of user-defined configuration settings that will define
	// how each associated datastore works. Each key/value pair is specific to a
	// datastore type.
	Values map[string]any `json:"values,omitempty"`
	// Associates the configuration group with a particular datastore.
	Datastore *DatastoreOpts `json:"datastore,omitempty"`
	// A human-readable explanation for the group.
	Description *string `json:"description,omitempty"`
}

// ToConfigUpdateMap will cast an UpdateOpts struct into a JSON map.
func (opts UpdateOpts) ToConfigUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update will modify an existing configuration group by performing a merge
// between new and existing values. If the key already exists, the new value
// will overwrite. All other keys will remain unaffected.
func Update(ctx context.Context, client *gophercloud.ServiceClient, configID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Replace will modify an existing configuration group by overwriting the
// entire parameter group with the new values provided. Any existing keys not
// included in UpdateOptsBuilder will be deleted.
func Replace(ctx context.Context, client *gophercloud.ServiceClient, configID string, opts UpdateOptsBuilder) (r ReplaceResult) {
	_ = "STUB: not implemented"
	return *new(ReplaceResult)
}

// Delete will permanently delete a configuration group. Please note that
// config groups cannot be deleted whilst still attached to running instances -
// you must detach and then delete them.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, configID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListInstances will list all the instances associated with a particular
// configuration group.
func ListInstances(client *gophercloud.ServiceClient, configID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListDatastoreParams will list all the available and supported parameters
// that can be used for a particular datastore ID and a particular version.
// For example, if you are wondering how you can configure a MySQL 5.6 instance,
// you can use this operation (you will need to retrieve the MySQL datastore ID
// by using the datastores API).
func ListDatastoreParams(client *gophercloud.ServiceClient, datastoreID, versionID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetDatastoreParam will retrieve information about a specific configuration
// parameter. For example, you can use this operation to understand more about
// "innodb_file_per_table" configuration param for MySQL datastores. You will
// need the param's ID first, which can be attained by using the ListDatastoreParams
// operation.
func GetDatastoreParam(ctx context.Context, client *gophercloud.ServiceClient, datastoreID, versionID, paramID string) (r ParamResult) {
	_ = "STUB: not implemented"
	return *new(ParamResult)
}

// ListGlobalParams is similar to ListDatastoreParams but does not require a
// DatastoreID.
func ListGlobalParams(client *gophercloud.ServiceClient, versionID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetGlobalParam is similar to GetDatastoreParam but does not require a
// DatastoreID.
func GetGlobalParam(ctx context.Context, client *gophercloud.ServiceClient, versionID, paramID string) (r ParamResult) {
	_ = "STUB: not implemented"
	return *new(ParamResult)
}
