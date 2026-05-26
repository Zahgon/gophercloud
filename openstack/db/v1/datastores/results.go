package datastores

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Version represents a version API resource. Multiple versions belong to a Datastore.
type Version struct {
	ID    string
	Links []gophercloud.Link
	Name  string
}

// Datastore represents a Datastore API resource.
type Datastore struct {
	DefaultVersion string `json:"default_version"`
	ID             string
	Links          []gophercloud.Link
	Name           string
	Versions       []Version
}

// DatastorePartial is a meta structure which is used in various API responses.
// It is a lightweight and truncated version of a full Datastore resource,
// offering details of the Version, Type and VersionID only.
type DatastorePartial struct {
	Version   string
	Type      string
	VersionID string `json:"version_id"`
}

// GetResult represents the result of a Get operation.
type GetResult struct {
	gophercloud.Result
}

// GetVersionResult represents the result of getting a version.
type GetVersionResult struct {
	gophercloud.Result
}

// DatastorePage represents a page of datastore resources.
type DatastorePage struct {
	pagination.SinglePageBase
}

// IsEmpty indicates whether a Datastore collection is empty.
func (r DatastorePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractDatastores retrieves a slice of datastore structs from a paginated
// collection.
func ExtractDatastores(r pagination.Page) ([]Datastore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract retrieves a single Datastore struct from an operation result.
func (r GetResult) Extract() (*Datastore, error) { _ = "STUB: not implemented"; return nil, nil }

// VersionPage represents a page of version resources.
type VersionPage struct {
	pagination.SinglePageBase
}

// IsEmpty indicates whether a collection of version resources is empty.
func (r VersionPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractVersions retrieves a slice of versions from a paginated collection.
func ExtractVersions(r pagination.Page) ([]Version, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract retrieves a single Version struct from an operation result.
func (r GetVersionResult) Extract() (*Version, error) { _ = "STUB: not implemented"; return nil, nil }
