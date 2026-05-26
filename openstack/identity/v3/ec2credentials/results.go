package ec2credentials

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Credential represents the application credential object
type Credential struct {
	// UserID contains a User ID of the EC2 credential owner.
	UserID string `json:"user_id"`
	// TenantID contains an EC2 credential project scope.
	TenantID string `json:"tenant_id"`
	// Access contains an EC2 credential access UUID.
	Access string `json:"access"`
	// Secret contains an EC2 credential secret UUID.
	Secret string `json:"secret"`
	// TrustID contains an EC2 credential trust ID scope.
	TrustID string `json:"trust_id"`
	// Links contains referencing links to the application credential.
	Links map[string]any `json:"links"`
}

type credentialResult struct {
	gophercloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as an Credential.
type GetResult struct {
	credentialResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as an Credential.
type CreateResult struct {
	credentialResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// an CredentialPage is a single page of an Credential results.
type CredentialPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a an CredentialPage contains any results.
func (r CredentialPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r CredentialPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Extractan Credentials returns a slice of Credentials contained in a single page of results.
func ExtractCredentials(r pagination.Page) ([]Credential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract interprets any Credential results as a Credential.
func (r credentialResult) Extract() (*Credential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
