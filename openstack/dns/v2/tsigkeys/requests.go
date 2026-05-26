package tsigkeys

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToTSIGKeyListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
// https://docs.openstack.org/api-ref/dns/
type ListOpts struct {
	// Integer value for the limit of values to return.
	Limit int `q:"limit"`

	// UUID of the TSIG key at which you want to set a marker.
	Marker string `q:"marker"`

	// Name of the TSIG key.
	Name string `q:"name"`

	// Algorithm used by the TSIG key.
	Algorithm string `q:"algorithm"`

	// Scope of the TSIG key (ZONE or POOL).
	Scope string `q:"scope"`
}

// ToTSIGKeyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTSIGKeyListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List implements a TSIG key List request.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get returns information about a TSIG key, given its ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, tsigkeyID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional attributes to the
// Create request.
type CreateOptsBuilder interface {
	ToTSIGKeyCreateMap() (map[string]any, error)
}

// CreateOpts specifies the attributes used to create a TSIG key.
type CreateOpts struct {
	// Name of the TSIG key.
	Name string `json:"name" required:"true"`

	// Algorithm is the TSIG algorithm (e.g., hmac-sha256, hmac-sha512).
	Algorithm string `json:"algorithm" required:"true"`

	// Secret is the base64-encoded secret key.
	Secret string `json:"secret" required:"true"`

	// Scope defines the scope of the TSIG key (ZONE or POOL).
	Scope string `json:"scope" required:"true"`

	// ResourceID is the ID of the resource (zone or pool) this key is associated with.
	ResourceID string `json:"resource_id" required:"true"`
}

// ToTSIGKeyCreateMap formats a CreateOpts structure into a request body.
func (opts CreateOpts) ToTSIGKeyCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create implements a TSIG key create request.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToTSIGKeyUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies the attributes to update a TSIG key.
type UpdateOpts struct {
	// Name of the TSIG key.
	Name string `json:"name,omitempty"`

	// Algorithm is the TSIG algorithm.
	Algorithm string `json:"algorithm,omitempty"`

	// Secret is the base64-encoded secret key.
	Secret string `json:"secret,omitempty"`

	// Scope defines the scope of the TSIG key.
	Scope string `json:"scope,omitempty"`

	// ResourceID is the ID of the resource this key is associated with.
	ResourceID string `json:"resource_id,omitempty"`
}

// ToTSIGKeyUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToTSIGKeyUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update implements a TSIG key update request.
func Update(ctx context.Context, client *gophercloud.ServiceClient, tsigkeyID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete implements a TSIG key delete request.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, tsigkeyID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
