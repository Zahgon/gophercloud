package clusters

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToClusterCreateMap() (map[string]any, error)
}

// CreateOpts params
type CreateOpts struct {
	ClusterTemplateID string            `json:"cluster_template_id" required:"true"`
	CreateTimeout     *int              `json:"create_timeout"`
	DiscoveryURL      string            `json:"discovery_url,omitempty"`
	DockerVolumeSize  *int              `json:"docker_volume_size,omitempty"`
	FlavorID          string            `json:"flavor_id,omitempty"`
	Keypair           string            `json:"keypair,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	MasterCount       *int              `json:"master_count,omitempty"`
	MasterFlavorID    string            `json:"master_flavor_id,omitempty"`
	Name              string            `json:"name"`
	NodeCount         *int              `json:"node_count,omitempty"`
	FloatingIPEnabled *bool             `json:"floating_ip_enabled,omitempty"`
	MasterLBEnabled   *bool             `json:"master_lb_enabled,omitempty"`
	FixedNetwork      string            `json:"fixed_network,omitempty"`
	FixedSubnet       string            `json:"fixed_subnet,omitempty"`
	MergeLabels       *bool             `json:"merge_labels,omitempty"`
}

// ToClusterCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToClusterCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new cluster.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get retrieves a specific clusters based on its unique ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Delete deletes the specified cluster ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToClustersListQuery() (string, error)
}

// ListOpts allows the sorting of paginated collections through
// the API. SortKey allows you to sort by a particular cluster attribute.
// SortDir sets the direction, and is either `asc' or `desc'.
// Marker and Limit are used for pagination.
type ListOpts struct {
	Marker  string `q:"marker"`
	Limit   int    `q:"limit"`
	SortKey string `q:"sort_key"`
	SortDir string `q:"sort_dir"`
}

// ToClustersListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToClustersListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// clusters. It accepts a ListOptsBuilder, which allows you to sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListDetail returns a Pager which allows you to iterate over a collection of
// clusters with detailed information.
// It accepts a ListOptsBuilder, which allows you to sort the returned
// collection for greater efficiency.
func ListDetail(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

type UpdateOp string

const (
	AddOp     UpdateOp = "add"
	RemoveOp  UpdateOp = "remove"
	ReplaceOp UpdateOp = "replace"
)

type UpdateOpts struct {
	Op    UpdateOp `json:"op" required:"true"`
	Path  string   `json:"path" required:"true"`
	Value any      `json:"value,omitempty"`
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToClustersUpdateMap() (map[string]any, error)
}

// ToClusterUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateOpts) ToClustersUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update implements cluster updated request.
func Update[T UpdateOptsBuilder](ctx context.Context, client *gophercloud.ServiceClient, id string, opts []T) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

type UpgradeOpts struct {
	ClusterTemplate string `json:"cluster_template" required:"true"`
	MaxBatchSize    *int   `json:"max_batch_size,omitempty"`
	NodeGroup       string `json:"nodegroup,omitempty"`
}

// UpgradeOptsBuilder allows extensions to add additional parameters to the
// Upgrade request.
type UpgradeOptsBuilder interface {
	ToClustersUpgradeMap() (map[string]any, error)
}

// ToClustersUpgradeMap constructs a request body from UpgradeOpts.
func (opts UpgradeOpts) ToClustersUpgradeMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Upgrade implements cluster upgrade request.
func Upgrade(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpgradeOptsBuilder) (r UpgradeResult) {
	_ = "STUB: not implemented"
	return *new(UpgradeResult)
}

// ResizeOptsBuilder allows extensions to add additional parameters to the
// Resize request.
type ResizeOptsBuilder interface {
	ToClusterResizeMap() (map[string]any, error)
}

// ResizeOpts params
type ResizeOpts struct {
	NodeCount     *int     `json:"node_count" required:"true"`
	NodesToRemove []string `json:"nodes_to_remove,omitempty"`
	NodeGroup     string   `json:"nodegroup,omitempty"`
}

// ToClusterResizeMap constructs a request body from ResizeOpts.
func (opts ResizeOpts) ToClusterResizeMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resize an existing cluster node count.
func Resize(ctx context.Context, client *gophercloud.ServiceClient, id string, opts ResizeOptsBuilder) (r ResizeResult) {
	_ = "STUB: not implemented"
	return *new(ResizeResult)
}
