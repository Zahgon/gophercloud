package secrets

import (
	"io"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Secret represents a secret stored in the key manager service.
type Secret struct {
	// BitLength is the bit length of the secret.
	BitLength int `json:"bit_length"`

	// Algorithm is the algorithm type of the secret.
	Algorithm string `json:"algorithm"`

	// Expiration is the expiration date of the secret.
	Expiration time.Time `json:"-"`

	// ContentTypes are the content types of the secret.
	ContentTypes map[string]string `json:"content_types"`

	// Created is the created date of the secret.
	Created time.Time `json:"-"`

	// CreatorID is the creator of the secret.
	CreatorID string `json:"creator_id"`

	// Mode is the mode of the secret.
	Mode string `json:"mode"`

	// Name is the name of the secret.
	Name string `json:"name"`

	// SecretRef is the URL to the secret.
	SecretRef string `json:"secret_ref"`

	// SecretType represents the type of secret.
	SecretType string `json:"secret_type"`

	// Status represents the status of the secret.
	Status string `json:"status"`

	// Updated is the updated date of the secret.
	Updated time.Time `json:"-"`
}

func (r *Secret) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as a Secret.
func (r commonResult) Extract() (*Secret, error) { _ = "STUB: not implemented"; return nil, nil }

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a secrets.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a secrets.
type CreateResult struct {
	commonResult
}

// UpdateResult is the response from an Update operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type UpdateResult struct {
	gophercloud.ErrResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// PayloadResult is the response from a GetPayload operation. Call its Extract
// method to extract the payload as a string.
type PayloadResult struct {
	gophercloud.Result
	Body io.ReadCloser
}

// Extract is a method that takes a PayloadResult's io.Reader body and reads
// all available data into a slice of bytes. Please be aware that its io.Reader
// is forward-only - meaning that it can only be read once and not rewound. You
// can recreate a reader from the output of this function by using
// bytes.NewReader(downloadBytes)
func (r PayloadResult) Extract() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SecretPage is a single page of secrets results.
type SecretPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of secrets contains any results.
func (r SecretPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r SecretPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractSecrets returns a slice of Secrets contained in a single page of
// results.
func ExtractSecrets(r pagination.Page) ([]Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MetadataResult is the result of a metadata request. Call its Extract method
// to interpret it as a map[string]string.
type MetadataResult struct {
	gophercloud.Result
}

// Extract interprets any MetadataResult as map[string]string.
func (r MetadataResult) Extract() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MetadataCreateResult is the result of a metadata create request. Call its
// Extract method to interpret it as a map[string]string.
type MetadataCreateResult struct {
	gophercloud.Result
}

// Extract interprets any MetadataCreateResult as a map[string]string.
func (r MetadataCreateResult) Extract() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Metadatum represents an individual metadata.
type Metadatum struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// MetadatumResult is the result of a metadatum request. Call its
// Extract method to interpret it as a map[string]string.
type MetadatumResult struct {
	gophercloud.Result
}

// Extract interprets any MetadatumResult as a map[string]string.
func (r MetadatumResult) Extract() (*Metadatum, error) { _ = "STUB: not implemented"; return nil, nil }

// MetadatumCreateResult is the response from a metadata Create operation. Call
// its ExtractErr method to determine if the request succeeded or failed.
//
// NOTE: This could be a MetadatumResponse but, at the time of testing, it looks
// like Barbican was returning errneous JSON in the response.
type MetadatumCreateResult struct {
	gophercloud.ErrResult
}

// MetadatumDeleteResult is the response from a metadatum Delete operation. Call
// its ExtractErr method to determine if the request succeeded or failed.
type MetadatumDeleteResult struct {
	gophercloud.ErrResult
}
