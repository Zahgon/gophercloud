package containers

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ContainerType represents the valid types of containers.
type ContainerType string

const (
	GenericContainer     ContainerType = "generic"
	RSAContainer         ContainerType = "rsa"
	CertificateContainer ContainerType = "certificate"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToContainerListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Limit is the amount of containers to retrieve.
	Limit int `q:"limit"`

	// Name is the name of the container
	Name string `q:"name"`

	// Offset is the index within the list to retrieve.
	Offset int `q:"offset"`
}

// ToContainerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToContainerListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List retrieves a list of containers.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details of a container.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToContainerCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a container.
type CreateOpts struct {
	// Type represents the type of container.
	Type ContainerType `json:"type" required:"true"`

	// Name is the name of the container.
	Name string `json:"name"`

	// SecretRefs is a list of secret refs for the container.
	SecretRefs []SecretRef `json:"secret_refs,omitempty"`
}

// ToContainerCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToContainerCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new container.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a container.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListConsumersOptsBuilder allows extensions to add additional parameters to
// the ListConsumers request
type ListConsumersOptsBuilder interface {
	ToContainerListConsumersQuery() (string, error)
}

// ListConsumersOpts provides options to filter the List results.
type ListConsumersOpts struct {
	// Limit is the amount of consumers to retrieve.
	Limit int `q:"limit"`

	// Offset is the index within the list to retrieve.
	Offset int `q:"offset"`
}

// ToContainerListConsumersQuery formats a ListConsumersOpts into a query
// string.
func (opts ListConsumersOpts) ToContainerListConsumersQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ListConsumers retrieves a list of consumers from a container.
func ListConsumers(client *gophercloud.ServiceClient, containerID string, opts ListConsumersOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateConsumerOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateConsumerOptsBuilder interface {
	ToContainerConsumerCreateMap() (map[string]any, error)
}

// CreateConsumerOpts provides options used to create a container.
type CreateConsumerOpts struct {
	// Name is the name of the consumer.
	Name string `json:"name"`

	// URL is the URL to the consumer resource.
	URL string `json:"URL"`
}

// ToContainerConsumerCreateMap formats a CreateConsumerOpts into a create
// request.
func (opts CreateConsumerOpts) ToContainerConsumerCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateConsumer creates a new consumer.
func CreateConsumer(ctx context.Context, client *gophercloud.ServiceClient, containerID string, opts CreateConsumerOptsBuilder) (r CreateConsumerResult) {
	_ = "STUB: not implemented"
	return *new(CreateConsumerResult)
}

// DeleteConsumerOptsBuilder allows extensions to add additional parameters to
// the Delete request.
type DeleteConsumerOptsBuilder interface {
	ToContainerConsumerDeleteMap() (map[string]any, error)
}

// DeleteConsumerOpts represents options used for deleting a consumer.
type DeleteConsumerOpts struct {
	// Name is the name of the consumer.
	Name string `json:"name"`

	// URL is the URL to the consumer resource.
	URL string `json:"URL"`
}

// ToContainerConsumerDeleteMap formats a DeleteConsumerOpts into a create
// request.
func (opts DeleteConsumerOpts) ToContainerConsumerDeleteMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteConsumer deletes a consumer.
func DeleteConsumer(ctx context.Context, client *gophercloud.ServiceClient, containerID string, opts DeleteConsumerOptsBuilder) (r DeleteConsumerResult) {
	_ = "STUB: not implemented"
	return *new(DeleteConsumerResult)
}

// SecretRefBuilder allows extensions to add additional parameters to the
// Create request.
type SecretRefBuilder interface {
	ToContainerSecretRefMap() (map[string]any, error)
}

// ToContainerSecretRefMap formats a SecretRefBuilder into a create
// request.
func (opts SecretRef) ToContainerSecretRefMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecret creates a new consumer.
func CreateSecretRef(ctx context.Context, client *gophercloud.ServiceClient, containerID string, opts SecretRefBuilder) (r CreateSecretRefResult) {
	_ = "STUB: not implemented"
	return *new(CreateSecretRefResult)
}

// DeleteSecret deletes a consumer.
func DeleteSecretRef(ctx context.Context, client *gophercloud.ServiceClient, containerID string, opts SecretRefBuilder) (r DeleteSecretRefResult) {
	_ = "STUB: not implemented"
	return *new(DeleteSecretRefResult)
}
