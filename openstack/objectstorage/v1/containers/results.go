package containers

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Container represents a container resource.
type Container struct {
	// The total number of bytes stored in the container.
	Bytes int64 `json:"bytes"`

	// The total number of objects stored in the container.
	Count int64 `json:"count"`

	// The name of the container.
	Name string `json:"name"`
}

// ContainerPage is the page returned by a pager when traversing over a
// collection of containers.
type ContainerPage struct {
	pagination.MarkerPageBase
}

// IsEmpty returns true if a ListResult contains no container names.
func (r ContainerPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// LastMarker returns the last container name in a ListResult.
func (r ContainerPage) LastMarker() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ExtractInfo is a function that takes a ListResult and returns the
// containers' information.
func ExtractInfo(r pagination.Page) ([]Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractNames is a function that takes a ListResult and returns the
// containers' names.
func ExtractNames(page pagination.Page) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetHeader represents the headers returned in the response from a Get request.
type GetHeader struct {
	AcceptRanges     string    `json:"Accept-Ranges"`
	BytesUsed        int64     `json:"X-Container-Bytes-Used,string"`
	ContentLength    int64     `json:"Content-Length,string"`
	ContentType      string    `json:"Content-Type"`
	Date             time.Time `json:"-"`
	ObjectCount      int64     `json:"X-Container-Object-Count,string"`
	Read             []string  `json:"-"`
	TransID          string    `json:"X-Trans-Id"`
	VersionsLocation string    `json:"X-Versions-Location"`
	HistoryLocation  string    `json:"X-History-Location"`
	Write            []string  `json:"-"`
	StoragePolicy    string    `json:"X-Storage-Policy"`
	TempURLKey       string    `json:"X-Container-Meta-Temp-URL-Key"`
	TempURLKey2      string    `json:"X-Container-Meta-Temp-URL-Key-2"`
	Timestamp        float64   `json:"X-Timestamp,string"`
	VersionsEnabled  bool      `json:"-"`
	SyncKey          string    `json:"X-Sync-Key"`
	SyncTo           string    `json:"X-Sync-To"`
}

func (r *GetHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// custom unmarshaller here is required to handle boolean value
// that starts with a capital letter

// GetResult represents the result of a get operation.
type GetResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Get.
func (r GetResult) Extract() (*GetHeader, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractMetadata is a function that takes a GetResult (of type *http.Response)
// and returns the custom metadata associated with the container.
func (r GetResult) ExtractMetadata() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateHeader represents the headers returned in the response from a Create
// request.
type CreateHeader struct {
	ContentLength int64     `json:"Content-Length,string"`
	ContentType   string    `json:"Content-Type"`
	Date          time.Time `json:"-"`
	TransID       string    `json:"X-Trans-Id"`
}

func (r *CreateHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// CreateResult represents the result of a create operation. To extract the
// the headers from the HTTP response, call its Extract method.
type CreateResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Create.
// To extract the headers from the HTTP response, call its Extract method.
func (r CreateResult) Extract() (*CreateHeader, error) { _ = "STUB: not implemented"; return nil, nil }

// UpdateHeader represents the headers returned in the response from a Update
// request.
type UpdateHeader struct {
	ContentLength int64     `json:"Content-Length,string"`
	ContentType   string    `json:"Content-Type"`
	Date          time.Time `json:"-"`
	TransID       string    `json:"X-Trans-Id"`
}

func (r *UpdateHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UpdateResult represents the result of an update operation. To extract the
// the headers from the HTTP response, call its Extract method.
type UpdateResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Update.
func (r UpdateResult) Extract() (*UpdateHeader, error) { _ = "STUB: not implemented"; return nil, nil }

// DeleteHeader represents the headers returned in the response from a Delete
// request.
type DeleteHeader struct {
	ContentLength int64     `json:"Content-Length,string"`
	ContentType   string    `json:"Content-Type"`
	Date          time.Time `json:"-"`
	TransID       string    `json:"X-Trans-Id"`
}

func (r *DeleteHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// DeleteResult represents the result of a delete operation. To extract the
// headers from the HTTP response, call its Extract method.
type DeleteResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Delete.
func (r DeleteResult) Extract() (*DeleteHeader, error) { _ = "STUB: not implemented"; return nil, nil }

type BulkDeleteResponse struct {
	ResponseStatus string     `json:"Response Status"`
	ResponseBody   string     `json:"Response Body"`
	Errors         [][]string `json:"Errors"`
	NumberDeleted  int        `json:"Number Deleted"`
	NumberNotFound int        `json:"Number Not Found"`
}

// BulkDeleteResult represents the result of a bulk delete operation. To extract
// the response object from the HTTP response, call its Extract method.
type BulkDeleteResult struct {
	gophercloud.Result
}

// Extract will return a BulkDeleteResponse struct returned from a BulkDelete
// call.
func (r BulkDeleteResult) Extract() (*BulkDeleteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
