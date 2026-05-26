package trusts

import (
	"context"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToTrustCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a new trust.
type CreateOpts struct {
	// Impersonation allows the trustee to impersonate the trustor.
	Impersonation bool `json:"impersonation"`

	// TrusteeUserID is a user who is capable of consuming the trust.
	TrusteeUserID string `json:"trustee_user_id" required:"true"`

	// TrustorUserID is a user who created the trust.
	TrustorUserID string `json:"trustor_user_id" required:"true"`

	// AllowRedelegation enables redelegation of a trust.
	AllowRedelegation bool `json:"allow_redelegation,omitempty"`

	// ExpiresAt sets expiration time on trust.
	ExpiresAt *time.Time `json:"-"`

	// ProjectID identifies the project.
	ProjectID string `json:"project_id,omitempty"`

	// RedelegationCount specifies a depth of the redelegation chain.
	RedelegationCount int `json:"redelegation_count,omitempty"`

	// RemainingUses specifies how many times a trust can be used to get a token.
	RemainingUses int `json:"remaining_uses,omitempty"`

	// Roles specifies roles that need to be granted to trustee.
	Roles []Role `json:"roles,omitempty"`
}

// ToTrustCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToTrustCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ListOptsBuilder interface {
	ToTrustListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// TrustorUserID filters the response by a trustor user Id.
	TrustorUserID string `q:"trustor_user_id"`

	// TrusteeUserID filters the response by a trustee user Id.
	TrusteeUserID string `q:"trustee_user_id"`
}

// ToTrustListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTrustListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Create creates a new Trust.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a Trust.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, trustID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// List enumerates the Trust to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single Trust, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ListRoles lists roles delegated by a Trust.
func ListRoles(client *gophercloud.ServiceClient, id string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetRole retrieves details on a single role delegated by a Trust.
func GetRole(ctx context.Context, client *gophercloud.ServiceClient, id string, roleID string) (r GetRoleResult) {
	_ = "STUB: not implemented"
	return *new(GetRoleResult)
}

// CheckRole checks whether a role ID is delegated by a Trust.
func CheckRole(ctx context.Context, client *gophercloud.ServiceClient, id string, roleID string) (r CheckRoleResult) {
	_ = "STUB: not implemented"
	return *new(CheckRoleResult)
}
