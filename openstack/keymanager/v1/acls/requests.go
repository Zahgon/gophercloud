package acls

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// GetContainerACL retrieves the ACL of a container.
func GetContainerACL(ctx context.Context, client *gophercloud.ServiceClient, containerID string) (r ACLResult) {
	_ = "STUB: not implemented"
	return *new(ACLResult)
}

// GetSecretACL retrieves the ACL of a secret.
func GetSecretACL(ctx context.Context, client *gophercloud.ServiceClient, secretID string) (r ACLResult) {
	_ = "STUB: not implemented"
	return *new(ACLResult)
}

// SetOptsBuilder allows extensions to add additional parameters to the
// Set request.
type SetOptsBuilder interface {
	ToACLSetMap() (map[string]any, error)
}

// SetOpt represents options to set a particular ACL type on a resource.
type SetOpt struct {
	// Type is the type of ACL to set. ie: read.
	Type string `json:"-" required:"true"`

	// Users are the list of Keystone user UUIDs.
	Users *[]string `json:"users,omitempty"`

	// ProjectAccess toggles if all users in a project can access the resource.
	ProjectAccess *bool `json:"project-access,omitempty"`
}

// SetOpts represents options to set an ACL on a resource.
type SetOpts []SetOpt

// ToACLSetMap formats a SetOpts into a set request.
func (opts SetOpts) ToACLSetMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetContainerACL will set an ACL on a container.
func SetContainerACL(ctx context.Context, client *gophercloud.ServiceClient, containerID string, opts SetOptsBuilder) (r ACLRefResult) {
	_ = "STUB: not implemented"
	return *new(ACLRefResult)
}

// SetSecretACL will set an ACL on a secret.
func SetSecretACL(ctx context.Context, client *gophercloud.ServiceClient, secretID string, opts SetOptsBuilder) (r ACLRefResult) {
	_ = "STUB: not implemented"
	return *new(ACLRefResult)
}

// UpdateContainerACL will update an ACL on a container.
func UpdateContainerACL(ctx context.Context, client *gophercloud.ServiceClient, containerID string, opts SetOptsBuilder) (r ACLRefResult) {
	_ = "STUB: not implemented"
	return *new(ACLRefResult)
}

// UpdateSecretACL will update an ACL on a secret.
func UpdateSecretACL(ctx context.Context, client *gophercloud.ServiceClient, secretID string, opts SetOptsBuilder) (r ACLRefResult) {
	_ = "STUB: not implemented"
	return *new(ACLRefResult)
}

// DeleteContainerACL will delete an ACL from a conatiner.
func DeleteContainerACL(ctx context.Context, client *gophercloud.ServiceClient, containerID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// DeleteSecretACL will delete an ACL from a secret.
func DeleteSecretACL(ctx context.Context, client *gophercloud.ServiceClient, secretID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
