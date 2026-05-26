package containers

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Container represents a container in the key manager service.
type Container struct {
	// Consumers are the consumers of the container.
	Consumers []ConsumerRef `json:"consumers"`

	// ContainerRef is the URL to the container
	ContainerRef string `json:"container_ref"`

	// Created is the date the container was created.
	Created time.Time `json:"-"`

	// CreatorID is the creator of the container.
	CreatorID string `json:"creator_id"`

	// Name is the name of the container.
	Name string `json:"name"`

	// SecretRefs are the secret references of the container.
	SecretRefs []SecretRef `json:"secret_refs"`

	// Status is the status of the container.
	Status string `json:"status"`

	// Type is the type of container.
	Type string `json:"type"`

	// Updated is the date the container was updated.
	Updated time.Time `json:"-"`
}

func (r *Container) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// ConsumerRef represents a consumer reference in a container.
type ConsumerRef struct {
	// Name is the name of the consumer.
	Name string `json:"name"`

	// URL is the URL to the consumer resource.
	URL string `json:"url"`
}

// SecretRef is a reference to a secret.
type SecretRef struct {
	SecretRef string `json:"secret_ref"`
	Name      string `json:"name"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as a Container.
func (r commonResult) Extract() (*Container, error) { _ = "STUB: not implemented"; return nil, nil }

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a container.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a container.
type CreateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ContainerPage is a single page of container results.
type ContainerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Container contains any results.
func (r ContainerPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r ContainerPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractContainers returns a slice of Containers contained in a single page of
// results.
func ExtractContainers(r pagination.Page) ([]Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consumer represents a consumer in a container.
type Consumer struct {
	// Created is the date the container was created.
	Created time.Time `json:"-"`

	// Name is the name of the container.
	Name string `json:"name"`

	// Status is the status of the container.
	Status string `json:"status"`

	// Updated is the date the container was updated.
	Updated time.Time `json:"-"`

	// URL is the url to the consumer.
	URL string `json:"url"`
}

func (r *Consumer) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// CreateConsumerResult is the response from a CreateConsumer operation.
// Call its Extract method to interpret it as a container.
type CreateConsumerResult struct {
	// This is not a typo: the API returns a Container, not a Consumer
	commonResult
}

// DeleteConsumerResult is the response from a DeleteConsumer operation.
// Call its Extract to interpret it as a container.
type DeleteConsumerResult struct {
	// This is not a typo: the API returns a Container, not a Consumer
	commonResult
}

// ConsumerPage is a single page of consumer results.
type ConsumerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of consumers contains any results.
func (r ConsumerPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r ConsumerPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractConsumers returns a slice of Consumers contained in a single page of
// results.
func ExtractConsumers(r pagination.Page) ([]Consumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract interprets any CreateSecretRefResult as a Container
func (r CreateSecretRefResult) Extract() (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecretRefResult is the response from a CreateSecretRef operation.
// Call its Extract method to interpret it as a container.
type CreateSecretRefResult struct {
	// This is not a typo.
	commonResult
}

// DeleteSecretRefResult is the response from a DeleteSecretRef operation.
type DeleteSecretRefResult struct {
	gophercloud.ErrResult
}
