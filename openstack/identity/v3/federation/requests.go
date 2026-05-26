package federation

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListMappings enumerates the mappings.
func ListMappings(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateMappingOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateMappingOptsBuilder interface {
	ToMappingCreateMap() (map[string]any, error)
}

// UpdateMappingOpts provides options for creating a mapping.
type CreateMappingOpts struct {
	// The list of rules used to map remote users into local users
	Rules []MappingRule `json:"rules"`
}

// ToMappingCreateMap formats a CreateMappingOpts into a create request.
func (opts CreateMappingOpts) ToMappingCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMapping creates a new Mapping.
func CreateMapping(ctx context.Context, client *gophercloud.ServiceClient, mappingID string, opts CreateMappingOptsBuilder) (r CreateMappingResult) {
	_ = "STUB: not implemented"
	return *new(CreateMappingResult)
}

// GetMapping retrieves details on a single mapping, by ID.
func GetMapping(ctx context.Context, client *gophercloud.ServiceClient, mappingID string) (r GetMappingResult) {
	_ = "STUB: not implemented"
	return *new(GetMappingResult)
}

// UpdateMappingOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateMappingOptsBuilder interface {
	ToMappingUpdateMap() (map[string]any, error)
}

// UpdateMappingOpts provides options for updating a mapping.
type UpdateMappingOpts struct {
	// The list of rules used to map remote users into local users
	Rules []MappingRule `json:"rules"`
}

// ToMappingUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateMappingOpts) ToMappingUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateMapping updates an existing mapping.
func UpdateMapping(ctx context.Context, client *gophercloud.ServiceClient, mappingID string, opts UpdateMappingOptsBuilder) (r UpdateMappingResult) {
	_ = "STUB: not implemented"
	return *new(UpdateMappingResult)
}

// DeleteMapping deletes a mapping.
func DeleteMapping(ctx context.Context, client *gophercloud.ServiceClient, mappingID string) (r DeleteMappingResult) {
	_ = "STUB: not implemented"
	return *new(DeleteMappingResult)
}
