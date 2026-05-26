package roles

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Option is a specific option defined at the API to enable features
// on a role.
type Option string

const (
	Immutable Option = "immutable"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToRoleListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// DomainID filters the response by a domain ID.
	DomainID string `q:"domain_id"`

	// Name filters the response by role name.
	Name string `q:"name"`

	// Filters filters the response by custom filters such as
	// 'name__contains=foo'
	Filters map[string]string `q:"-"`
}

// ToRoleListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRoleListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List enumerates the roles to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single role, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToRoleCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a role.
type CreateOpts struct {
	// Name is the name of the new role.
	Name string `json:"name" required:"true"`

	// DomainID is the ID of the domain the role belongs to.
	DomainID string `json:"domain_id,omitempty"`

	// Description is the description of the new role.
	Description string `json:"description,omitempty"`

	// Extra is free-form extra key/value pairs to describe the role.
	Extra map[string]any `json:"-"`

	// Options are defined options in the API to enable certain features.
	Options map[Option]any `json:"options,omitempty"`
}

// ToRoleCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToRoleCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new Role.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToRoleUpdateMap() (map[string]any, error)
}

// UpdateOpts provides options for updating a role.
type UpdateOpts struct {
	// Name is an updated name for the role.
	Name string `json:"name,omitempty"`

	// Description is an updated description for the role.
	Description *string `json:"description,omitempty"`

	// Extra is free-form extra key/value pairs to describe the role.
	Extra map[string]any `json:"-"`

	// Options are defined options in the API to enable certain features.
	Options map[Option]any `json:"options,omitempty"`
}

// ToRoleUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToRoleUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update updates an existing Role.
func Update(ctx context.Context, client *gophercloud.ServiceClient, roleID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete deletes a role.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, roleID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListAssignmentsOptsBuilder allows extensions to add additional parameters to
// the ListAssignments request.
type ListAssignmentsOptsBuilder interface {
	ToRolesListAssignmentsQuery() (string, error)
}

// ListAssignmentsOpts allows you to query the ListAssignments method.
// Specify one of or a combination of GroupId, RoleId, ScopeDomainId,
// ScopeProjectId, and/or UserId to search for roles assigned to corresponding
// entities.
type ListAssignmentsOpts struct {
	// GroupID is the group ID to query.
	GroupID string `q:"group.id"`

	// RoleID is the specific role to query assignments to.
	RoleID string `q:"role.id"`

	// ScopeDomainID filters the results by the given domain ID.
	ScopeDomainID string `q:"scope.domain.id"`

	// ScopeProjectID filters the results by the given Project ID.
	ScopeProjectID string `q:"scope.project.id"`

	// UserID filterst he results by the given User ID.
	UserID string `q:"user.id"`

	// Effective lists effective assignments at the user, project, and domain
	// level, allowing for the effects of group membership.
	Effective *bool `q:"effective"`

	// IncludeNames indicates whether to include names of any returned entities.
	// Requires microversion 3.6 or later.
	IncludeNames *bool `q:"include_names"`

	// IncludeSubtree indicates whether to include relevant assignments in the project hierarchy below the project
	// specified in the ScopeProjectID. Specify DomainID in ScopeProjectID to get a list for all projects in the domain.
	// Requires microversion 3.6 or later.
	IncludeSubtree *bool `q:"include_subtree"`
}

// ToRolesListAssignmentsQuery formats a ListAssignmentsOpts into a query string.
func (opts ListAssignmentsOpts) ToRolesListAssignmentsQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ListAssignments enumerates the roles assigned to a specified resource.
func ListAssignments(client *gophercloud.ServiceClient, opts ListAssignmentsOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListAssignmentsOnResourceOpts provides options to list role assignments
// for a user/group on a project/domain
type ListAssignmentsOnResourceOpts struct {
	// UserID is the ID of a user to assign a role
	// Note: exactly one of UserID or GroupID must be provided
	UserID string `xor:"GroupID"`

	// GroupID is the ID of a group to assign a role
	// Note: exactly one of UserID or GroupID must be provided
	GroupID string `xor:"UserID"`

	// ProjectID is the ID of a project to assign a role on
	// Note: exactly one of ProjectID or DomainID must be provided
	ProjectID string `xor:"DomainID"`

	// DomainID is the ID of a domain to assign a role on
	// Note: exactly one of ProjectID or DomainID must be provided
	DomainID string `xor:"ProjectID"`
}

// AssignOpts provides options to assign a role
type AssignOpts struct {
	// UserID is the ID of a user to assign a role
	// Note: exactly one of UserID or GroupID must be provided
	UserID string `xor:"GroupID"`

	// GroupID is the ID of a group to assign a role
	// Note: exactly one of UserID or GroupID must be provided
	GroupID string `xor:"UserID"`

	// ProjectID is the ID of a project to assign a role on
	// Note: exactly one of ProjectID or DomainID must be provided
	ProjectID string `xor:"DomainID"`

	// DomainID is the ID of a domain to assign a role on
	// Note: exactly one of ProjectID or DomainID must be provided
	DomainID string `xor:"ProjectID"`
}

// UnassignOpts provides options to unassign a role
type UnassignOpts struct {
	// UserID is the ID of a user to unassign a role
	// Note: exactly one of UserID or GroupID must be provided
	UserID string `xor:"GroupID"`

	// GroupID is the ID of a group to unassign a role
	// Note: exactly one of UserID or GroupID must be provided
	GroupID string `xor:"UserID"`

	// ProjectID is the ID of a project to unassign a role on
	// Note: exactly one of ProjectID or DomainID must be provided
	ProjectID string `xor:"DomainID"`

	// DomainID is the ID of a domain to unassign a role on
	// Note: exactly one of ProjectID or DomainID must be provided
	DomainID string `xor:"ProjectID"`
}

// ListAssignmentsOnResource is the operation responsible for listing role
// assignments for a user/group on a project/domain.
func ListAssignmentsOnResource(client *gophercloud.ServiceClient, opts ListAssignmentsOnResourceOpts) pagination.Pager {
	_ = "STUB: not implemented"
	// Check xor conditions
	return *new(pagination.Pager)
}

// Get corresponding URL

// Assign is the operation responsible for assigning a role
// to a user/group on a project/domain.
func Assign(ctx context.Context, client *gophercloud.ServiceClient, roleID string, opts AssignOpts) (r AssignmentResult) {
	_ = "STUB: not implemented"
	// Check xor conditions
	return *new(AssignmentResult)
}

// Get corresponding URL

// Unassign is the operation responsible for unassigning a role
// from a user/group on a project/domain.
func Unassign(ctx context.Context, client *gophercloud.ServiceClient, roleID string, opts UnassignOpts) (r UnassignmentResult) {
	_ = "STUB: not implemented"
	// Check xor conditions
	return *new(UnassignmentResult)
}

// Get corresponding URL

func CreateRoleInferenceRule(ctx context.Context, client *gophercloud.ServiceClient, priorRoleID, impliedRoleID string) (r CreateImpliedRoleResult) {
	_ = "STUB: not implemented"
	return *new(CreateImpliedRoleResult)
}

func GetRoleInferenceRule(ctx context.Context, client *gophercloud.ServiceClient, priorRoleID, impliedRoleID string) (r CreateImpliedRoleResult) {
	_ = "STUB: not implemented"
	return *new(CreateImpliedRoleResult)
}

func DeleteRoleInferenceRule(ctx context.Context, client *gophercloud.ServiceClient, priorRoleID, impliedRoleID string) (r DeleteImpliedRoleResult) {
	_ = "STUB: not implemented"
	return *new(DeleteImpliedRoleResult)
}

func ListRoleInferenceRules(ctx context.Context, client *gophercloud.ServiceClient) (r ListImpliedRolesResult) {
	_ = "STUB: not implemented"
	return *new(ListImpliedRolesResult)
}
