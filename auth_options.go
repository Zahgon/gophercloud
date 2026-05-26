package gophercloud

/*
AuthOptions stores information needed to authenticate to an OpenStack Cloud.
You can populate one manually, or use a provider's AuthOptionsFromEnv() function
to read relevant information from the standard environment variables. Pass one
to a provider's AuthenticatedClient function to authenticate and obtain a
ProviderClient representing an active session on that provider.

Its fields are the union of those recognized by each identity implementation and
provider.

An example of manually providing authentication information:

	opts := gophercloud.AuthOptions{
	  IdentityEndpoint: "https://openstack.example.com:5000/v2.0",
	  Username: "{username}",
	  Password: "{password}",
	  TenantID: "{tenant_id}",
	}

	provider, err := openstack.AuthenticatedClient(context.TODO(), opts)

An example of using AuthOptionsFromEnv(), where the environment variables can
be read from a file, such as a standard openrc file:

	opts, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(context.TODO(), opts)
*/
type AuthOptions struct {
	// IdentityEndpoint specifies the HTTP endpoint that is required to work with
	// the Identity API of the appropriate version. While it's ultimately needed by
	// all of the identity services, it will often be populated by a provider-level
	// function.
	//
	// The IdentityEndpoint is typically referred to as the "auth_url" or
	// "OS_AUTH_URL" in the information provided by the cloud operator.
	IdentityEndpoint string `json:"-"`

	// Username is required if using Identity V2 API. Consult with your provider's
	// control panel to discover your account's username. In Identity V3, either
	// UserID or a combination of Username and DomainID or DomainName are needed.
	Username string `json:"username,omitempty"`
	UserID   string `json:"-"`

	Password string `json:"password,omitempty"`

	// Passcode is used in TOTP authentication method
	Passcode string `json:"passcode,omitempty"`

	// At most one of DomainID and DomainName must be provided if using Username
	// with Identity V3. Otherwise, either are optional.
	DomainID   string `json:"-"`
	DomainName string `json:"name,omitempty"`

	// The TenantID and TenantName fields are optional for the Identity V2 API.
	// The same fields are known as project_id and project_name in the Identity
	// V3 API, but are collected as TenantID and TenantName here in both cases.
	// Some providers allow you to specify a TenantName instead of the TenantId.
	// Some require both. Your provider's authentication policies will determine
	// how these fields influence authentication.
	// If DomainID or DomainName are provided, they will also apply to TenantName.
	// It is not currently possible to authenticate with Username and a Domain
	// and scope to a Project in a different Domain by using TenantName. To
	// accomplish that, the ProjectID will need to be provided as the TenantID
	// option.
	TenantID   string `json:"tenantId,omitempty"`
	TenantName string `json:"tenantName,omitempty"`

	// AllowReauth should be set to true if you grant permission for Gophercloud to
	// cache your credentials in memory, and to allow Gophercloud to attempt to
	// re-authenticate automatically if/when your token expires.  If you set it to
	// false, it will not cache these settings, but re-authentication will not be
	// possible.  This setting defaults to false.
	//
	// NOTE: The reauth function will try to re-authenticate endlessly if left
	// unchecked. The way to limit the number of attempts is to provide a custom
	// HTTP client to the provider client and provide a transport that implements
	// the RoundTripper interface and stores the number of failed retries. For an
	// example of this, see here:
	// https://github.com/rackspace/rack/blob/1.0.0/auth/clients.go#L311
	AllowReauth bool `json:"-"`

	// TokenID allows users to authenticate (possibly as another user) with an
	// authentication token ID.
	TokenID string `json:"-"`

	// Scope determines the scoping of the authentication request.
	Scope *AuthScope `json:"-"`

	// Authentication through Application Credentials requires supplying name, project and secret
	// For project we can use TenantID
	ApplicationCredentialID     string `json:"-"`
	ApplicationCredentialName   string `json:"-"`
	ApplicationCredentialSecret string `json:"-"`
}

// AuthScope allows a created token to be limited to a specific domain or project.
type AuthScope struct {
	ProjectID   string
	ProjectName string
	DomainID    string
	DomainName  string
	System      bool
	TrustID     string
}

// ToTokenV2CreateMap allows AuthOptions to satisfy the AuthOptionsBuilder
// interface in the v2 tokens package
func (opts AuthOptions) ToTokenV2CreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	// Populate the request map.
	return nil, nil
}

// ToTokenV3CreateMap allows AuthOptions to satisfy the AuthOptionsBuilder
// interface in the v3 tokens package
func (opts *AuthOptions) ToTokenV3CreateMap(scope map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Populate the request structure based on the provided arguments. Create and return an error
// if insufficient or incompatible information is present.

// Because we aren't using password authentication, it's an error to also provide any of the user-based authentication
// parameters.

// Configure the request for Token authentication.

// Configure the request for ApplicationCredentialID authentication.
// https://github.com/openstack/keystoneauth/blob/stable/rocky/keystoneauth1/identity/v3/application_credential.py#L48-L67
// There are three kinds of possible application_credential requests
// 1. application_credential id + secret
// 2. application_credential name + secret + user_id
// 3. application_credential name + secret + username + domain_id / domain_name

// UserID could be used without the domain information

// Make sure that Username or UserID are provided

// Make sure that DomainID or DomainName are provided among Username

// If no password or token ID or ApplicationCredential are available, authentication can't continue.

// Password authentication.

// TOTP authentication.

// At least one of Username and UserID must be specified.

// If Username is provided, UserID may not be provided.

// Either DomainID or DomainName must also be specified.

// Configure the request for Username and Password authentication with a DomainID.

// Configure the request for Username and Password authentication with a DomainName.

// If UserID is specified, neither DomainID nor DomainName may be.

// Configure the request for UserID and Password authentication.

// ToTokenV3ScopeMap builds a scope from AuthOptions and satisfies interface in
// the v3 tokens package.
func (opts *AuthOptions) ToTokenV3ScopeMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	// For backwards compatibility.
	// If AuthOptions.Scope was not set, try to determine it.
	// This works well for common scenarios.
	return nil, nil
}

// ProjectName provided: either DomainID or DomainName must also be supplied.
// ProjectID may not be supplied.

// ProjectName + DomainID

// ProjectName + DomainName

// ProjectID provided. ProjectName, DomainID, and DomainName may not be provided.

// ProjectID

// DomainID provided. ProjectID, ProjectName, and DomainName may not be provided.

// DomainID

// DomainName

func (opts AuthOptions) CanReauth() bool { _ = "STUB: not implemented"; return false }

// cannot reauth using TOTP passcode

// ToTokenV3HeadersMap allows AuthOptions to satisfy the AuthOptionsBuilder
// interface in the v3 tokens package.
func (opts *AuthOptions) ToTokenV3HeadersMap(map[string]any) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
