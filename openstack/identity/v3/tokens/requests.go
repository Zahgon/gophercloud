package tokens

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

const xSubjectTokenHeader = "X-Subject-Token"

// Scope allows a created token to be limited to a specific domain or project.
type Scope struct {
	ProjectID   string
	ProjectName string
	DomainID    string
	DomainName  string
	System      bool
	TrustID     string
}

// AuthOptionsBuilder provides the ability for extensions to add additional
// parameters to AuthOptions. Extensions must satisfy all required methods.
type AuthOptionsBuilder interface {
	// ToTokenV3CreateMap assembles the Create request body, returning an error
	// if parameters are missing or inconsistent.
	ToTokenV3CreateMap(map[string]any) (map[string]any, error)
	ToTokenV3HeadersMap(map[string]any) (map[string]string, error)
	ToTokenV3ScopeMap() (map[string]any, error)
	CanReauth() bool
}

// AuthOptions represents options for authenticating a user.
type AuthOptions struct {
	// IdentityEndpoint specifies the HTTP endpoint that is required to work with
	// the Identity API of the appropriate version. While it's ultimately needed
	// by all of the identity services, it will often be populated by a
	// provider-level function.
	IdentityEndpoint string `json:"-"`

	// Username is required if using Identity V2 API. Consult with your provider's
	// control panel to discover your account's username. In Identity V3, either
	// UserID or a combination of Username and DomainID or DomainName are needed.
	Username string `json:"username,omitempty"`
	UserID   string `json:"id,omitempty"`

	Password string `json:"password,omitempty"`

	// Passcode is used in TOTP authentication method
	Passcode string `json:"passcode,omitempty"`

	// At most one of DomainID and DomainName must be provided if using Username
	// with Identity V3. Otherwise, either are optional.
	DomainID   string `json:"-"`
	DomainName string `json:"name,omitempty"`

	// AllowReauth should be set to true if you grant permission for Gophercloud
	// to cache your credentials in memory, and to allow Gophercloud to attempt
	// to re-authenticate automatically if/when your token expires.  If you set
	// it to false, it will not cache these settings, but re-authentication will
	// not be possible.  This setting defaults to false.
	AllowReauth bool `json:"-"`

	// TokenID allows users to authenticate (possibly as another user) with an
	// authentication token ID.
	TokenID string `json:"-"`

	// Authentication through Application Credentials requires supplying name, project and secret
	// For project we can use TenantID
	ApplicationCredentialID     string `json:"-"`
	ApplicationCredentialName   string `json:"-"`
	ApplicationCredentialSecret string `json:"-"`

	Scope Scope `json:"-"`
}

// ToTokenV3CreateMap builds a request body from AuthOptions.
func (opts *AuthOptions) ToTokenV3CreateMap(scope map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToTokenV3ScopeMap builds a scope request body from AuthOptions.
func (opts *AuthOptions) ToTokenV3ScopeMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (opts *AuthOptions) CanReauth() bool { _ = "STUB: not implemented"; return false }

// cannot reauth using TOTP passcode

// ToTokenV3HeadersMap allows AuthOptions to satisfy the AuthOptionsBuilder
// interface in the v3 tokens package.
func (opts *AuthOptions) ToTokenV3HeadersMap(map[string]any) (map[string]string, error) {
	_ = "STUB: not implemented"

	// Create authenticates and either generates a new token, or changes the Scope
	// of an existing token.
	return nil, nil
}

func Create(ctx context.Context, c *gophercloud.ServiceClient, opts AuthOptionsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// GetOptsBuilder allows extensions to add additional parameters to
// the Get request.
type GetOptsBuilder interface {
	ToTokenGetParams() (map[string]string, error)
}

// GetOpts provides options for the Get request.
type GetOpts struct {
	// AccessRulesVersion specifies the OpenStack-Identity-Access-Rules header
	// version. This is required when getting tokens that were created using
	// application credentials with access rules. Versions less than 1 will cause
	// Keystone to ignore access rules validation.
	AccessRulesVersion string `h:"OpenStack-Identity-Access-Rules"`
}

// ToTokenGetParams formats GetOpts into request headers.
func (opts GetOpts) ToTokenGetParams() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get validates and retrieves information about another token.
func Get(ctx context.Context, c *gophercloud.ServiceClient, token string, opts GetOptsBuilder) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ValidateOptsBuilder allows extensions to add additional parameters to
// the Validate request.
type ValidateOptsBuilder interface {
	ToTokenValidateParams() (map[string]string, error)
}

// ValidateOpts provides options for the Validate request.
type ValidateOpts struct {
	// AccessRulesVersion specifies the OpenStack-Identity-Access-Rules header
	// version. This is required when validating tokens that were created using
	// application credentials with access rules. Versions less than 1 will cause
	// Keystone to ignore access rules validation.
	AccessRulesVersion string `h:"OpenStack-Identity-Access-Rules"`
}

// ToTokenValidateParams formats ValidateOpts into request headers.
func (opts ValidateOpts) ToTokenValidateParams() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate determines if a specified token is valid or not.
func Validate(ctx context.Context, c *gophercloud.ServiceClient, token string, opts ValidateOptsBuilder) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Revoke immediately makes specified token invalid.
func Revoke(ctx context.Context, c *gophercloud.ServiceClient, token string) (r RevokeResult) {
	_ = "STUB: not implemented"
	return *new(RevokeResult)
}
