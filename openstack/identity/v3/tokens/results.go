package tokens

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
)

// Endpoint represents a single API endpoint offered by a service.
// It matches either a public, internal or admin URL.
// If supported, it contains a region specifier, again if provided.
// The significance of the Region field will depend upon your provider.
type Endpoint struct {
	ID        string `json:"id"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
	Interface string `json:"interface"`
	URL       string `json:"url"`
}

// CatalogEntry provides a type-safe interface to an Identity API V3 service
// catalog listing. Each class of service, such as cloud DNS or block storage
// services, could have multiple CatalogEntry representing it (one by interface
// type, e.g public, admin or internal).
//
// Note: when looking for the desired service, try, whenever possible, to key
// off the type field. Otherwise, you'll tie the representation of the service
// to a specific provider.
type CatalogEntry struct {
	// Service ID
	ID string `json:"id"`

	// Name will contain the provider-specified name for the service.
	Name string `json:"name"`

	// Type will contain a type string if OpenStack defines a type for the
	// service. Otherwise, for provider-specific services, the provider may
	// assign their own type strings.
	Type string `json:"type"`

	// Endpoints will let the caller iterate over all the different endpoints that
	// may exist for the service.
	Endpoints []Endpoint `json:"endpoints"`
}

// ServiceCatalog provides a view into the service catalog from a previous,
// successful authentication.
type ServiceCatalog struct {
	Entries []CatalogEntry `json:"catalog"`
}

// Domain provides information about the domain to which this token grants
// access.
type Domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// User represents a user resource that exists in the Identity Service.
type User struct {
	Domain Domain `json:"domain"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

// Role provides information about roles to which User is authorized.
type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Project provides information about project to which User is authorized.
type Project struct {
	Domain Domain `json:"domain"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

type TrustUser struct {
	ID string `json:"id"`
}

// Trust provides information about trust with which User is authorized.
type Trust struct {
	ID            string    `json:"id"`
	Impersonation bool      `json:"impersonation"`
	TrusteeUserID TrustUser `json:"trustee_user"`
	TrustorUserID TrustUser `json:"trustor_user"`
}

// AccessRule represents an access rule for an application credential.
type AccessRule struct {
	// The ID of the access rule.
	ID string `json:"id"`
	// The API path that the application credential is permitted to access.
	Path string `json:"path"`
	// The request method that the application credential is permitted to use
	// for a given API endpoint.
	Method string `json:"method"`
	// The service type identifier for the service that the application
	// credential is permitted to access.
	Service string `json:"service"`
}

// ApplicationCredential represents the application credential information
// included in a token when the token was created using an application credential.
type ApplicationCredential struct {
	// The ID of the application credential.
	ID string `json:"id"`
	// The name of the application credential.
	Name string `json:"name"`
	// A flag indicating whether the application credential is restricted.
	Restricted bool `json:"restricted"`
	// A list of access rules for the application credential.
	AccessRules []AccessRule `json:"access_rules"`
}

// commonResult is the response from a request. A commonResult has various
// methods which can be used to extract different details about the result.
type commonResult struct {
	gophercloud.Result
}

// Extract is a shortcut for ExtractToken.
// This function is deprecated and still present for backward compatibility.
func (r commonResult) Extract() (*Token, error) {
	_ = "STUB: not implemented"
	return nil,

		// ExtractToken interprets a commonResult as a Token.
		nil
}

func (r commonResult) ExtractToken() (*Token, error) { _ = "STUB: not implemented"; return nil, nil }

// Parse the token itself from the stored headers.

// ExtractTokenID implements the gophercloud.AuthResult interface. The returned
// string is the same as the ID field of the Token struct returned from
// ExtractToken().
func (r CreateResult) ExtractTokenID() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ExtractTokenID implements the gophercloud.AuthResult interface. The returned
// string is the same as the ID field of the Token struct returned from
// ExtractToken().
func (r GetResult) ExtractTokenID() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ExtractServiceCatalog returns the ServiceCatalog that was generated along
// with the user's Token.
func (r commonResult) ExtractServiceCatalog() (*ServiceCatalog, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractUser returns the User that is the owner of the Token.
func (r commonResult) ExtractUser() (*User, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractRoles returns Roles to which User is authorized.
func (r commonResult) ExtractRoles() ([]Role, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractProject returns Project to which User is authorized.
func (r commonResult) ExtractProject() (*Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractDomain returns Domain to which User is authorized.
func (r commonResult) ExtractDomain() (*Domain, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractTrust returns Trust to which User is authorized.
func (r commonResult) ExtractTrust() (*Trust, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractApplicationCredential returns the ApplicationCredential that was used
// to create the token. This is only present when the token was created using
// an application credential.
func (r commonResult) ExtractApplicationCredential() (*ApplicationCredential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateResult is the response from a Create request. Use ExtractToken()
// to interpret it as a Token, or ExtractServiceCatalog() to interpret it
// as a service catalog.
type CreateResult struct {
	commonResult
}

// GetResult is the response from a Get request. Use ExtractToken()
// to interpret it as a Token, or ExtractServiceCatalog() to interpret it
// as a service catalog.
type GetResult struct {
	commonResult
}

// RevokeResult is response from a Revoke request.
type RevokeResult struct {
	commonResult
}

// Token is a string that grants a user access to a controlled set of services
// in an OpenStack provider. Each Token is valid for a set length of time.
type Token struct {
	// ID is the issued token.
	ID string `json:"id"`

	// ExpiresAt is the timestamp at which this token will no longer be accepted.
	ExpiresAt time.Time `json:"expires_at"`
}

func (r commonResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }
