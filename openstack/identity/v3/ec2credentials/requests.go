package ec2credentials

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List enumerates the Credentials to which the current token has access.
func List(client *gophercloud.ServiceClient, userID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single EC2 credential by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, userID string, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToCredentialCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create an EC2 credential.
type CreateOpts struct {
	// TenantID is the project ID scope of the EC2 credential.
	TenantID string `json:"tenant_id" required:"true"`
}

// ToCredentialCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToCredentialCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new EC2 Credential.
func Create(ctx context.Context, client *gophercloud.ServiceClient, userID string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes an EC2 credential.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, userID string, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
