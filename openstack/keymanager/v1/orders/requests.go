package orders

import (
	"context"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// OrderType represents the valid types of orders.
type OrderType string

const (
	KeyOrder        OrderType = "key"
	AsymmetricOrder OrderType = "asymmetric"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToOrderListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Limit is the amount of containers to retrieve.
	Limit int `q:"limit"`

	// Offset is the index within the list to retrieve.
	Offset int `q:"offset"`
}

// ToOrderListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToOrderListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List retrieves a list of orders.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details of a orders.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToOrderCreateMap() (map[string]any, error)
}

// MetaOpts represents options used for creating an order.
type MetaOpts struct {
	// Algorithm is the algorithm of the secret.
	Algorithm string `json:"algorithm"`

	// BitLength is the bit length of the secret.
	BitLength int `json:"bit_length"`

	// Expiration is the expiration date of the order.
	Expiration *time.Time `json:"-"`

	// Mode is the mode of the secret.
	Mode string `json:"mode"`

	// Name is the name of the secret.
	Name string `json:"name,omitempty"`

	// PayloadContentType is the content type of the secret payload.
	PayloadContentType string `json:"payload_content_type,omitempty"`
}

// CreateOpts provides options used to create a orders.
type CreateOpts struct {
	// Type is the type of order to create.
	Type OrderType `json:"type"`

	// Meta contains secrets data to create a secret.
	Meta MetaOpts `json:"meta"`
}

// ToOrderCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToOrderCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new orders.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a orders.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
