package tokens

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// PasswordCredentialsV2 represents the required options to authenticate
// with a username and password.
type PasswordCredentialsV2 struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
}

// TokenCredentialsV2 represents the required options to authenticate
// with a token.
type TokenCredentialsV2 struct {
	ID string `json:"id,omitempty" required:"true"`
}

// AuthOptionsV2 wraps a gophercloud AuthOptions in order to adhere to the
// AuthOptionsBuilder interface.
type AuthOptionsV2 struct {
	PasswordCredentials *PasswordCredentialsV2 `json:"passwordCredentials,omitempty" xor:"TokenCredentials"`

	// The TenantID and TenantName fields are optional for the Identity V2 API.
	// Some providers allow you to specify a TenantName instead of the TenantId.
	// Some require both. Your provider's authentication policies will determine
	// how these fields influence authentication.
	TenantID   string `json:"tenantId,omitempty"`
	TenantName string `json:"tenantName,omitempty"`

	// TokenCredentials allows users to authenticate (possibly as another user)
	// with an authentication token ID.
	TokenCredentials *TokenCredentialsV2 `json:"token,omitempty" xor:"PasswordCredentials"`
}

// AuthOptionsBuilder allows extensions to add additional parameters to the
// token create request.
type AuthOptionsBuilder interface {
	// ToTokenCreateMap assembles the Create request body, returning an error
	// if parameters are missing or inconsistent.
	ToTokenV2CreateMap() (map[string]any, error)
	CanReauth() bool
}

// AuthOptions are the valid options for Openstack Identity v2 authentication.
// For field descriptions, see gophercloud.AuthOptions.
type AuthOptions struct {
	IdentityEndpoint string `json:"-"`
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
	TenantID         string `json:"tenantId,omitempty"`
	TenantName       string `json:"tenantName,omitempty"`
	AllowReauth      bool   `json:"-"`
	TokenID          string
}

// ToTokenV2CreateMap builds a token request body from the given AuthOptions.
func (opts AuthOptions) ToTokenV2CreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (opts AuthOptions) CanReauth() bool { _ = "STUB: not implemented"; return false }

// Create authenticates to the identity service and attempts to acquire a Token.
// Generally, rather than interact with this call directly, end users should
// call openstack.AuthenticatedClient(), which abstracts all of the gory details
// about navigating service catalogs and such.
func Create(ctx context.Context, client *gophercloud.ServiceClient, auth AuthOptionsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get validates and retrieves information for user's token.
func Get(ctx context.Context, client *gophercloud.ServiceClient, token string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
