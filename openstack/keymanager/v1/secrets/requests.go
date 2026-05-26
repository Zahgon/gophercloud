package secrets

import (
	"context"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// DateFilter represents a valid filter to use for filtering
// secrets by their date during a list.
type DateFilter string

const (
	DateFilterGT  DateFilter = "gt"
	DateFilterGTE DateFilter = "gte"
	DateFilterLT  DateFilter = "lt"
	DateFilterLTE DateFilter = "lte"
)

// DateQuery represents a date field to be used for listing secrets.
// If no filter is specified, the query will act as if "equal" is used.
type DateQuery struct {
	Date   time.Time
	Filter DateFilter
}

// SecretType represents a valid secret type.
type SecretType string

const (
	SymmetricSecret   SecretType = "symmetric"
	PublicSecret      SecretType = "public"
	PrivateSecret     SecretType = "private"
	PassphraseSecret  SecretType = "passphrase"
	CertificateSecret SecretType = "certificate"
	OpaqueSecret      SecretType = "opaque"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToSecretListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Offset is the starting index within the total list of the secrets that
	// you would like to retrieve.
	Offset int `q:"offset"`

	// Limit is the maximum number of records to return.
	Limit int `q:"limit"`

	// Name will select all secrets with a matching name.
	Name string `q:"name"`

	// Alg will select all secrets with a matching algorithm.
	Alg string `q:"alg"`

	// Mode will select all secrets with a matching mode.
	Mode string `q:"mode"`

	// Bits will select all secrets with a matching bit length.
	Bits int `q:"bits"`

	// SecretType will select all secrets with a matching secret type.
	SecretType SecretType `q:"secret_type"`

	// ACLOnly will select all secrets with an ACL that contains the user.
	ACLOnly *bool `q:"acl_only"`

	// CreatedQuery will select all secrets with a created date matching
	// the query.
	CreatedQuery *DateQuery

	// UpdatedQuery will select all secrets with an updated date matching
	// the query.
	UpdatedQuery *DateQuery

	// ExpirationQuery will select all secrets with an expiration date
	// matching the query.
	ExpirationQuery *DateQuery

	// Sort will sort the results in the requested order.
	Sort string `q:"sort"`
}

// ToSecretListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSecretListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List retrieves a list of Secrets.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details of a secrets.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// GetPayloadOpts represents options used for obtaining a payload.
type GetPayloadOpts struct {
	PayloadContentType string `h:"Accept"`
}

// GetPayloadOptsBuilder allows extensions to add additional parameters to
// the GetPayload request.
type GetPayloadOptsBuilder interface {
	ToSecretPayloadGetParams() (map[string]string, error)
}

// ToSecretPayloadGetParams formats a GetPayloadOpts into a query string.
func (opts GetPayloadOpts) ToSecretPayloadGetParams() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPayload retrieves the payload of a secret.
func GetPayload(ctx context.Context, client *gophercloud.ServiceClient, id string, opts GetPayloadOptsBuilder) (r PayloadResult) {
	_ = "STUB: not implemented"
	return *new(PayloadResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToSecretCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a secrets.
type CreateOpts struct {
	// Algorithm is the algorithm of the secret.
	Algorithm string `json:"algorithm,omitempty"`

	// BitLength is the bit length of the secret.
	BitLength int `json:"bit_length,omitempty"`

	// Mode is the mode of encryption for the secret.
	Mode string `json:"mode,omitempty"`

	// Name is the name of the secret
	Name string `json:"name,omitempty"`

	// Payload is the secret.
	Payload string `json:"payload,omitempty"`

	// PayloadContentType is the content type of the payload.
	PayloadContentType string `json:"payload_content_type,omitempty"`

	// PayloadContentEncoding is the content encoding of the payload.
	PayloadContentEncoding string `json:"payload_content_encoding,omitempty"`

	// SecretType is the type of secret.
	SecretType SecretType `json:"secret_type,omitempty"`

	// Expiration is the expiration date of the secret.
	Expiration *time.Time `json:"-"`
}

// ToSecretCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToSecretCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new secrets.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a secrets.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToSecretUpdateRequest() (string, map[string]string, error)
}

// UpdateOpts represents parameters to add a payload to an existing
// secret which does not already contain a payload.
type UpdateOpts struct {
	// ContentType represents the content type of the payload.
	ContentType string `h:"Content-Type"`

	// ContentEncoding represents the content encoding of the payload.
	ContentEncoding string `h:"Content-Encoding"`

	// Payload is the payload of the secret.
	Payload string
}

// ToUpdateCreateRequest formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToSecretUpdateRequest() (string, map[string]string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Update modifies the attributes of a secrets.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// GetMetadata will list metadata for a given secret.
func GetMetadata(ctx context.Context, client *gophercloud.ServiceClient, secretID string) (r MetadataResult) {
	_ = "STUB: not implemented"
	return *new(MetadataResult)
}

// MetadataOpts is a map that contains key-value pairs for secret metadata.
type MetadataOpts map[string]string

// CreateMetadataOptsBuilder allows extensions to add additional parameters to
// the CreateMetadata request.
type CreateMetadataOptsBuilder interface {
	ToMetadataCreateMap() (map[string]any, error)
}

// ToMetadataCreateMap converts a MetadataOpts into a request body.
func (opts MetadataOpts) ToMetadataCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMetadata will set metadata for a given secret.
func CreateMetadata(ctx context.Context, client *gophercloud.ServiceClient, secretID string, opts CreateMetadataOptsBuilder) (r MetadataCreateResult) {
	_ = "STUB: not implemented"
	return *new(MetadataCreateResult)
}

// GetMetadatum will get a single key/value metadata from a secret.
func GetMetadatum(ctx context.Context, client *gophercloud.ServiceClient, secretID string, key string) (r MetadatumResult) {
	_ = "STUB: not implemented"
	return *new(MetadatumResult)
}

// MetadatumOpts represents a single metadata.
type MetadatumOpts struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value" required:"true"`
}

// CreateMetadatumOptsBuilder allows extensions to add additional parameters to
// the CreateMetadatum request.
type CreateMetadatumOptsBuilder interface {
	ToMetadatumCreateMap() (map[string]any, error)
}

// ToMetadatumCreateMap converts a MetadatumOpts into a request body.
func (opts MetadatumOpts) ToMetadatumCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMetadatum will add a single key/value metadata to a secret.
func CreateMetadatum(ctx context.Context, client *gophercloud.ServiceClient, secretID string, opts CreateMetadatumOptsBuilder) (r MetadatumCreateResult) {
	_ = "STUB: not implemented"
	return *new(MetadatumCreateResult)
}

// UpdateMetadatumOptsBuilder allows extensions to add additional parameters to
// the UpdateMetadatum request.
type UpdateMetadatumOptsBuilder interface {
	ToMetadatumUpdateMap() (map[string]any, string, error)
}

// ToMetadatumUpdateMap converts a MetadataOpts into a request body.
func (opts MetadatumOpts) ToMetadatumUpdateMap() (map[string]any, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// UpdateMetadatum will update a single key/value metadata to a secret.
func UpdateMetadatum(ctx context.Context, client *gophercloud.ServiceClient, secretID string, opts UpdateMetadatumOptsBuilder) (r MetadatumResult) {
	_ = "STUB: not implemented"
	return *new(MetadatumResult)
}

// DeleteMetadatum will delete an individual metadatum from a secret.
func DeleteMetadatum(ctx context.Context, client *gophercloud.ServiceClient, secretID string, key string) (r MetadatumDeleteResult) {
	_ = "STUB: not implemented"
	return *new(MetadatumDeleteResult)
}
