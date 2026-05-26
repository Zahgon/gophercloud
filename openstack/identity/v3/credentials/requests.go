package credentials

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToCredentialListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// UserID filters the response by a credential user_id
	UserID string `q:"user_id"`
	// Type filters the response by a credential type
	Type string `q:"type"`
}

// ToCredentialListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToCredentialListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List enumerates the Credentials to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single user, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToCredentialCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a credential.
type CreateOpts struct {
	// Serialized blob containing the credentials
	Blob string `json:"blob" required:"true"`
	// ID of the project.
	ProjectID string `json:"project_id,omitempty"`
	// The type of the credential.
	Type string `json:"type" required:"true"`
	// ID of the user who owns the credential.
	UserID string `json:"user_id" required:"true"`
}

// ToCredentialCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToCredentialCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new Credential.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a credential.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToCredentialsUpdateMap() (map[string]any, error)
}

// UpdateOpts represents parameters to update a credential.
type UpdateOpts struct {
	// Serialized blob containing the credentials.
	Blob string `json:"blob,omitempty"`
	// ID of the project.
	ProjectID string `json:"project_id,omitempty"`
	// The type of the credential.
	Type string `json:"type,omitempty"`
	// ID of the user who owns the credential.
	UserID string `json:"user_id,omitempty"`
}

// ToUpdateCreateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToCredentialsUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update modifies the attributes of a Credential.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
