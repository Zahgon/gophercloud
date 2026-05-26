package users

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Option is a specific option defined at the API to enable features
// on a user account.
type Option string

const (
	IgnoreChangePasswordUponFirstUse Option = "ignore_change_password_upon_first_use"
	IgnorePasswordExpiry             Option = "ignore_password_expiry"
	IgnoreLockoutFailureAttempts     Option = "ignore_lockout_failure_attempts"
	MultiFactorAuthRules             Option = "multi_factor_auth_rules"
	MultiFactorAuthEnabled           Option = "multi_factor_auth_enabled"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToUserListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// DomainID filters the response by a domain ID.
	DomainID string `q:"domain_id"`

	// Enabled filters the response by enabled users.
	Enabled *bool `q:"enabled"`

	// IdpID filters the response by an Identity Provider ID.
	IdPID string `q:"idp_id"`

	// Name filters the response by username.
	Name string `q:"name"`

	// PasswordExpiresAt filters the response based on expiring passwords.
	PasswordExpiresAt string `q:"password_expires_at"`

	// ProtocolID filters the response by protocol ID.
	ProtocolID string `q:"protocol_id"`

	// UniqueID filters the response by unique ID.
	UniqueID string `q:"unique_id"`

	// Filters filters the response by custom filters such as
	// 'name__contains=foo'
	Filters map[string]string `q:"-"`
}

// ToUserListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToUserListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List enumerates the Users to which the current token has access.
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
	ToUserCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a user.
type CreateOpts struct {
	// Name is the name of the new user.
	Name string `json:"name" required:"true"`

	// DefaultProjectID is the ID of the default project of the user.
	DefaultProjectID string `json:"default_project_id,omitempty"`

	// Description is a description of the user.
	Description string `json:"description,omitempty"`

	// DomainID is the ID of the domain the user belongs to.
	DomainID string `json:"domain_id,omitempty"`

	// Enabled sets the user status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Extra is free-form extra key/value pairs to describe the user.
	Extra map[string]any `json:"-"`

	// Options are defined options in the API to enable certain features.
	Options map[Option]any `json:"options,omitempty"`

	// Password is the password of the new user.
	Password string `json:"password,omitempty"`
}

// ToUserCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToUserCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new User.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToUserUpdateMap() (map[string]any, error)
}

// UpdateOpts provides options for updating a user account.
type UpdateOpts struct {
	// Name is the name of the new user.
	Name string `json:"name,omitempty"`

	// DefaultProjectID is the ID of the default project of the user.
	DefaultProjectID string `json:"default_project_id,omitempty"`

	// Description is a description of the user.
	Description *string `json:"description,omitempty"`

	// DomainID is the ID of the domain the user belongs to.
	DomainID string `json:"domain_id,omitempty"`

	// Enabled sets the user status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Extra is free-form extra key/value pairs to describe the user.
	Extra map[string]any `json:"-"`

	// Options are defined options in the API to enable certain features.
	Options map[Option]any `json:"options,omitempty"`

	// Password is the password of the new user.
	Password string `json:"password,omitempty"`
}

// ToUserUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToUserUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update updates an existing User.
func Update(ctx context.Context, client *gophercloud.ServiceClient, userID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// ChangePasswordOptsBuilder allows extensions to add additional parameters to
// the ChangePassword request.
type ChangePasswordOptsBuilder interface {
	ToUserChangePasswordMap() (map[string]any, error)
}

// ChangePasswordOpts provides options for changing password for a user.
type ChangePasswordOpts struct {
	// OriginalPassword is the original password of the user.
	OriginalPassword string `json:"original_password"`

	// Password is the new password of the user.
	Password string `json:"password"`
}

// ToUserChangePasswordMap formats a ChangePasswordOpts into a ChangePassword request.
func (opts ChangePasswordOpts) ToUserChangePasswordMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChangePassword changes password for a user.
func ChangePassword(ctx context.Context, client *gophercloud.ServiceClient, userID string, opts ChangePasswordOptsBuilder) (r ChangePasswordResult) {
	_ = "STUB: not implemented"
	return *new(ChangePasswordResult)
}

// Delete deletes a user.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, userID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListGroups enumerates groups user belongs to.
func ListGroups(client *gophercloud.ServiceClient, userID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// AddToGroup adds a user to a group.
func AddToGroup(ctx context.Context, client *gophercloud.ServiceClient, groupID, userID string) (r AddToGroupResult) {
	_ = "STUB: not implemented"
	return *new(AddToGroupResult)
}

// IsMemberOfGroup checks whether a user belongs to a group.
func IsMemberOfGroup(ctx context.Context, client *gophercloud.ServiceClient, groupID, userID string) (r IsMemberOfGroupResult) {
	_ = "STUB: not implemented"
	return *new(IsMemberOfGroupResult)
}

// RemoveFromGroup removes a user from a group.
func RemoveFromGroup(ctx context.Context, client *gophercloud.ServiceClient, groupID, userID string) (r RemoveFromGroupResult) {
	_ = "STUB: not implemented"
	return *new(RemoveFromGroupResult)
}

// ListProjects enumerates groups user belongs to.
func ListProjects(client *gophercloud.ServiceClient, userID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListInGroup enumerates users that belong to a group.
func ListInGroup(client *gophercloud.ServiceClient, groupID string, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
