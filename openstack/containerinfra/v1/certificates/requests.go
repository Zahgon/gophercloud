package certificates

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// CreateOptsBuilder allows extensions to add additional parameters
// to the Create request.
type CreateOptsBuilder interface {
	ToCertificateCreateMap() (map[string]any, error)
}

// CreateOpts represents options used to create a certificate.
type CreateOpts struct {
	ClusterUUID string `json:"cluster_uuid,omitempty" xor:"BayUUID"`
	BayUUID     string `json:"bay_uuid,omitempty" xor:"ClusterUUID"`
	CSR         string `json:"csr" required:"true"`
}

// ToCertificateCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToCertificateCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get makes a request against the API to get details for a certificate.
func Get(ctx context.Context, client *gophercloud.ServiceClient, clusterID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Create requests the creation of a new certificate.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Update will rotate the CA certificate for a cluster
func Update(ctx context.Context, client *gophercloud.ServiceClient, clusterID string) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
