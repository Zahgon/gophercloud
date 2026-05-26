package qos

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type CreateOptsBuilder interface {
	ToQoSCreateMap() (map[string]any, error)
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToQoSListQuery() (string, error)
}

type QoSConsumer string

const (
	ConsumerFront QoSConsumer = "front-end"
	ConsumerBack  QoSConsumer = "back-end"
	ConsumerBoth  QoSConsumer = "both"
)

// CreateOpts contains options for creating a QoS specification.
// This object is passed to the qos.Create function.
type CreateOpts struct {
	// The name of the QoS spec
	Name string `json:"name"`
	// The consumer of the QoS spec. Possible values are
	// both, front-end, back-end.
	Consumer QoSConsumer `json:"consumer,omitempty"`
	// Specs is a collection of miscellaneous key/values used to set
	// specifications for the QoS
	Specs map[string]string `json:"-"`
}

// ToQoSCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToQoSCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create will create a new QoS based on the values in CreateOpts. To extract
// the QoS object from the response, call the Extract method on the
// CreateResult.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// DeleteOptsBuilder allows extensions to add additional parameters to the
// Delete request.
type DeleteOptsBuilder interface {
	ToQoSDeleteQuery() (string, error)
}

// DeleteOpts contains options for deleting a QoS. This object is passed to
// the qos.Delete function.
type DeleteOpts struct {
	// Delete a QoS specification even if it is in-use
	Force bool `q:"force"`
}

// ToQoSDeleteQuery formats a DeleteOpts into a query string.
func (opts DeleteOpts) ToQoSDeleteQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Delete will delete the existing QoS with the provided ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string, opts DeleteOptsBuilder) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

type ListOpts struct {
	// Sort is Comma-separated list of sort keys and optional sort
	// directions in the form of < key > [: < direction > ]. A valid
	//direction is asc (ascending) or desc (descending).
	Sort string `q:"sort"`

	// Marker and Limit control paging.
	// Marker instructs List where to start listing from.
	Marker string `q:"marker"`

	// Limit instructs List to refrain from sending excessively large lists of
	// QoS.
	Limit int `q:"limit"`
}

// ToQoSListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToQoSListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List instructs OpenStack to provide a list of QoS.
// You may provide criteria by which List curtails its results for easier
// processing.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details of a single qos. Use Extract to convert its
// result into a QoS.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateQosSpecsOptsBuilder allows extensions to add additional parameters to the
// CreateQosSpecs requests.
type CreateQosSpecsOptsBuilder interface {
	ToQosSpecsCreateMap() (map[string]any, error)
}

// UpdateOpts contains options for creating a QoS specification.
// This object is passed to the qos.Update function.
type UpdateOpts struct {
	// The consumer of the QoS spec. Possible values are
	// both, front-end, back-end.
	Consumer QoSConsumer `json:"consumer,omitempty"`
	// Specs is a collection of miscellaneous key/values used to set
	// specifications for the QoS
	Specs map[string]string `json:"-"`
}

type UpdateOptsBuilder interface {
	ToQoSUpdateMap() (map[string]any, error)
}

// ToQoSUpdateMap assembles a request body based on the contents of a
// UpdateOpts.
func (opts UpdateOpts) ToQoSUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update will update an existing QoS based on the values in UpdateOpts.
// To extract the QoS object from the response, call the Extract method
// on the UpdateResult.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r updateResult) {
	_ = "STUB: not implemented"
	return *new(updateResult)
}

// DeleteKeysOptsBuilder allows extensions to add additional parameters to the
// CreateExtraSpecs requests.
type DeleteKeysOptsBuilder interface {
	ToDeleteKeysCreateMap() (map[string]any, error)
}

// DeleteKeysOpts is a string slice that contains keys to be deleted.
type DeleteKeysOpts []string

// ToDeleteKeysCreateMap assembles a body for a Create request based on
// the contents of ExtraSpecsOpts.
func (opts DeleteKeysOpts) ToDeleteKeysCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteKeys will delete the keys/specs from the specified QoS
func DeleteKeys(ctx context.Context, client *gophercloud.ServiceClient, qosID string, opts DeleteKeysOptsBuilder) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// AssociateOpitsBuilder allows extensions to define volume type id
// to the associate query
type AssociateOptsBuilder interface {
	ToQosAssociateQuery() (string, error)
}

// AssociateOpts contains options for associating a QoS with a
// volume type
type AssociateOpts struct {
	VolumeTypeID string `q:"vol_type_id" required:"true"`
}

// ToQosAssociateQuery formats an AssociateOpts into a query string
func (opts AssociateOpts) ToQosAssociateQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Associate will associate a qos with a volute type
func Associate(ctx context.Context, client *gophercloud.ServiceClient, qosID string, opts AssociateOptsBuilder) (r AssociateResult) {
	_ = "STUB: not implemented"
	return *new(AssociateResult)
}

// DisassociateOpitsBuilder allows extensions to define volume type id
// to the disassociate query
type DisassociateOptsBuilder interface {
	ToQosDisassociateQuery() (string, error)
}

// DisassociateOpts contains options for disassociating a QoS from a
// volume type
type DisassociateOpts struct {
	VolumeTypeID string `q:"vol_type_id" required:"true"`
}

// ToQosDisassociateQuery formats a DisassociateOpts into a query string
func (opts DisassociateOpts) ToQosDisassociateQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Disassociate will disassociate a qos from a volute type
func Disassociate(ctx context.Context, client *gophercloud.ServiceClient, qosID string, opts DisassociateOptsBuilder) (r DisassociateResult) {
	_ = "STUB: not implemented"
	return *new(DisassociateResult)
}

// DisassociateAll will disassociate a qos from all volute types
func DisassociateAll(ctx context.Context, client *gophercloud.ServiceClient, qosID string) (r DisassociateAllResult) {
	_ = "STUB: not implemented"
	return *new(DisassociateAllResult)
}

// ListAssociations retrieves the associations of a QoS.
func ListAssociations(client *gophercloud.ServiceClient, qosID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
