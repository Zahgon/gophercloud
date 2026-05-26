package configurations

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Config represents a configuration group API resource.
type Config struct {
	Created              time.Time `json:"-"`
	Updated              time.Time `json:"-"`
	DatastoreName        string    `json:"datastore_name"`
	DatastoreVersionID   string    `json:"datastore_version_id"`
	DatastoreVersionName string    `json:"datastore_version_name"`
	Description          string
	ID                   string
	Name                 string
	Values               map[string]any
}

func (r *Config) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// ConfigPage contains a page of Config resources in a paginated collection.
type ConfigPage struct {
	pagination.SinglePageBase
}

// IsEmpty indicates whether a ConfigPage is empty.
func (r ConfigPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractConfigs will retrieve a slice of Config structs from a page.
func ExtractConfigs(r pagination.Page) ([]Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type commonResult struct {
	gophercloud.Result
}

// Extract will retrieve a Config resource from an operation result.
func (r commonResult) Extract() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// GetResult represents the result of a Get operation.
type GetResult struct {
	commonResult
}

// CreateResult represents the result of a Create operation.
type CreateResult struct {
	commonResult
}

// UpdateResult represents the result of an Update operation.
type UpdateResult struct {
	gophercloud.ErrResult
}

// ReplaceResult represents the result of a Replace operation.
type ReplaceResult struct {
	gophercloud.ErrResult
}

// DeleteResult represents the result of a Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Param represents a configuration parameter API resource.
type Param struct {
	Max             float64
	Min             float64
	Name            string
	RestartRequired bool `json:"restart_required"`
	Type            string
}

// ParamPage contains a page of Param resources in a paginated collection.
type ParamPage struct {
	pagination.SinglePageBase
}

// IsEmpty indicates whether a ParamPage is empty.
func (r ParamPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractParams will retrieve a slice of Param structs from a page.
func ExtractParams(r pagination.Page) ([]Param, error) { _ = "STUB: not implemented"; return nil, nil }

// ParamResult represents the result of an operation which retrieves details
// about a particular configuration param.
type ParamResult struct {
	gophercloud.Result
}

// Extract will retrieve a param from an operation result.
func (r ParamResult) Extract() (*Param, error) { _ = "STUB: not implemented"; return nil, nil }
