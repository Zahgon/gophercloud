package users

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List lists the existing users.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CommonOpts are the parameters that are shared between CreateOpts and
// UpdateOpts
type CommonOpts struct {
	// Either a name or username is required. When provided, the value must be
	// unique or a 409 conflict error will be returned. If you provide a name but
	// omit a username, the latter will be set to the former; and vice versa.
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`

	// TenantID is the ID of the tenant to which you want to assign this user.
	TenantID string `json:"tenantId,omitempty"`

	// Enabled indicates whether this user is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`

	// Email is the email address of this user.
	Email string `json:"email,omitempty"`
}

// CreateOpts represents the options needed when creating new users.
type CreateOpts CommonOpts

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToUserCreateMap() (map[string]any, error)
}

// ToUserCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToUserCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create is the operation responsible for creating new users.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get requests details on a single user, either by ID or Name.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToUserUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies the base attributes that may be updated on an
// existing server.
type UpdateOpts CommonOpts

// ToUserUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToUserUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update is the operation responsible for updating exist users by their ID.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete is the operation responsible for permanently deleting a User.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListRoles lists the existing roles that can be assigned to users.
func ListRoles(client *gophercloud.ServiceClient, tenantID, userID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
