package roles

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List is the operation responsible for listing all available global roles
// that a user can adopt.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// AddUser is the operation responsible for assigning a particular role to
// a user. This is confined to the scope of the user's tenant - so the tenant
// ID is a required argument.
func AddUser(ctx context.Context, client *gophercloud.ServiceClient, tenantID, userID, roleID string) (r UserRoleResult) {
	_ = "STUB: not implemented"
	return *new(UserRoleResult)
}

// DeleteUser is the operation responsible for deleting a particular role
// from a user. This is confined to the scope of the user's tenant - so the
// tenant ID is a required argument.
func DeleteUser(ctx context.Context, client *gophercloud.ServiceClient, tenantID, userID, roleID string) (r UserRoleResult) {
	_ = "STUB: not implemented"
	return *new(UserRoleResult)
}
