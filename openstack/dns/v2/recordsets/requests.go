package recordsets

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToRecordSetListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
// https://developer.openstack.org/api-ref/dns/
type ListOpts struct {
	// Integer value for the limit of values to return.
	Limit int `q:"limit"`

	// UUID of the recordset at which you want to set a marker.
	Marker string `q:"marker"`

	Data        string `q:"data"`
	Description string `q:"description"`
	Name        string `q:"name"`
	SortDir     string `q:"sort_dir"`
	SortKey     string `q:"sort_key"`
	Status      string `q:"status"`
	TTL         int    `q:"ttl"`
	Type        string `q:"type"`
	ZoneID      string `q:"zone_id"`
}

// ToRecordSetListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRecordSetListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ListByZone implements the recordset list request.
func ListByZone(client *gophercloud.ServiceClient, zoneID string, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get implements the recordset Get request.
func Get(ctx context.Context, client *gophercloud.ServiceClient, zoneID string, rrsetID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional attributes to the
// Create request.
type CreateOptsBuilder interface {
	ToRecordSetCreateMap() (map[string]any, error)
}

// CreateOpts specifies the base attributes that may be used to create a
// RecordSet.
type CreateOpts struct {
	// Name is the name of the RecordSet.
	Name string `json:"name" required:"true"`

	// Description is a description of the RecordSet.
	Description string `json:"description,omitempty"`

	// Records are the DNS records of the RecordSet.
	Records []string `json:"records,omitempty"`

	// TTL is the time to live of the RecordSet.
	TTL int `json:"ttl,omitempty"`

	// Type is the RRTYPE of the RecordSet.
	Type string `json:"type,omitempty"`
}

// ToRecordSetCreateMap formats an CreateOpts structure into a request body.
func (opts CreateOpts) ToRecordSetCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a recordset in a given zone.
func Create(ctx context.Context, client *gophercloud.ServiceClient, zoneID string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToRecordSetUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies the base attributes that may be updated on an existing
// RecordSet.
type UpdateOpts struct {
	// Description is a description of the RecordSet.
	Description *string `json:"description,omitempty"`

	// TTL is the time to live of the RecordSet.
	TTL *int `json:"ttl,omitempty"`

	// Records are the DNS records of the RecordSet.
	Records []string `json:"records,omitempty"`
}

// ToRecordSetUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToRecordSetUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If opts.TTL was actually set, use 0 as a special value to send "null",
// even though the result from the API is 0.
//
// Otherwise, don't send the TTL field.

// Update updates a recordset in a given zone
func Update(ctx context.Context, client *gophercloud.ServiceClient, zoneID string, rrsetID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete removes an existing RecordSet.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, zoneID string, rrsetID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
