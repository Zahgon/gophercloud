package nodegroups

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Get makes a request to the Magnum API to retrieve a node group
// with the given ID/name belonging to the given cluster.
// Use the Extract method of the returned GetResult to extract the
// node group from the result.
func Get(ctx context.Context, client *gophercloud.ServiceClient, clusterID, nodeGroupID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

type ListOptsBuilder interface {
	ToNodeGroupsListQuery() (string, error)
}

// ListOpts is used to filter and sort the node groups of a cluster
// when using List.
type ListOpts struct {
	// Pagination marker for large data sets. (UUID field from node group).
	Marker int `q:"marker"`
	// Maximum number of resources to return in a single page.
	Limit int `q:"limit"`
	// Column to sort results by. Default: id.
	SortKey string `q:"sort_key"`
	// Direction to sort. "asc" or "desc". Default: asc.
	SortDir string `q:"sort_dir"`
	// List all nodegroups with the specified role.
	Role string `q:"role"`
}

func (opts ListOpts) ToNodeGroupsListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request to the Magnum API to retrieve node groups
// belonging to the given cluster. The request can be modified to
// filter or sort the list using the options available in ListOpts.
//
// Use the AllPages method of the returned Pager to ensure that
// all node groups are returned (for example when using the Limit
// option to limit the number of node groups returned per page).
//
// Not all node group fields are returned in a list request.
// Only the fields UUID, Name, FlavorID, ImageID,
// NodeCount, Role, IsDefault, Status and StackID
// are returned, all other fields are omitted
// and will have their zero value when extracted.
func List(client *gophercloud.ServiceClient, clusterID string, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

type CreateOptsBuilder interface {
	ToNodeGroupCreateMap() (map[string]any, error)
}

// CreateOpts is used to set available fields upon node group creation.
//
// If unset, some fields have defaults or will inherit from the cluster value.
type CreateOpts struct {
	Name             string `json:"name" required:"true"`
	DockerVolumeSize *int   `json:"docker_volume_size,omitempty"`
	// Labels will default to the cluster labels if unset.
	Labels       map[string]string `json:"labels,omitempty"`
	NodeCount    *int              `json:"node_count,omitempty"`
	MinNodeCount int               `json:"min_node_count,omitempty"`
	// MaxNodeCount can be left unset for no maximum node count.
	MaxNodeCount *int `json:"max_node_count,omitempty"`
	// Role defaults to "worker" if unset.
	Role string `json:"role,omitempty"`
	// Node image ID. Defaults to cluster template image if unset.
	ImageID string `json:"image_id,omitempty"`
	// Node machine flavor ID. Defaults to cluster minion flavor if unset.
	FlavorID    string `json:"flavor_id,omitempty"`
	MergeLabels *bool  `json:"merge_labels,omitempty"`
}

func (opts CreateOpts) ToNodeGroupCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create makes a request to the Magnum API to create a node group
// for the the given cluster.
// Use the Extract method of the returned CreateResult to extract the
// returned node group.
func Create(ctx context.Context, client *gophercloud.ServiceClient, clusterID string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

type UpdateOptsBuilder interface {
	ToResourceUpdateMap() (map[string]any, error)
}

type UpdateOp string

const (
	AddOp     UpdateOp = "add"
	RemoveOp  UpdateOp = "remove"
	ReplaceOp UpdateOp = "replace"
)

// UpdateOpts is used to define the action taken when updating a node group.
//
// Valid Ops are "add", "remove", "replace"
// Valid Paths are "/min_node_count" and "/max_node_count"
type UpdateOpts struct {
	Op    UpdateOp `json:"op" required:"true"`
	Path  string   `json:"path" required:"true"`
	Value any      `json:"value,omitempty"`
}

func (opts UpdateOpts) ToResourceUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update makes a request to the Magnum API to update a field of
// the given node group belonging to the given cluster. More than
// one UpdateOpts can be passed at a time.
// Use the Extract method of the returned UpdateResult to extract the
// updated node group from the result.
func Update[T UpdateOptsBuilder](ctx context.Context, client *gophercloud.ServiceClient, clusterID string, nodeGroupID string, opts []T) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete makes a request to the Magnum API to delete a node group.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, clusterID, nodeGroupID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
