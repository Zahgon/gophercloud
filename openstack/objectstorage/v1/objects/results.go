package objects

import (
	"io"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Object is a structure that holds information related to a storage object.
type Object struct {
	// Bytes is the total number of bytes that comprise the object.
	Bytes int64 `json:"bytes"`

	// ContentType is the content type of the object.
	ContentType string `json:"content_type"`

	// Hash represents the MD5 checksum value of the object's content.
	Hash string `json:"hash"`

	// LastModified is the time the object was last modified.
	LastModified time.Time `json:"-"`

	// Name is the unique name for the object.
	Name string `json:"name"`

	// Subdir denotes if the result contains a subdir.
	Subdir string `json:"subdir"`

	// IsLatest indicates whether the object version is the latest one.
	IsLatest bool `json:"is_latest"`

	// VersionID contains a version ID of the object, when container
	// versioning is enabled.
	VersionID string `json:"version_id"`
}

func (r *Object) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// ObjectPage is a single page of objects that is returned from a call to the
// List function.
type ObjectPage struct {
	pagination.MarkerPageBase
}

// IsEmpty returns true if a ListResult contains no object names.
func (r ObjectPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// LastMarker returns the last object name in a ListResult.
func (r ObjectPage) LastMarker() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// ExtractInfo is a function that takes a page of objects and returns their
		// full information.
		nil
}

func ExtractInfo(r pagination.Page) ([]Object, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractNames is a function that takes a page of objects and returns only
// their names.
func ExtractNames(r pagination.Page) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// DownloadHeader represents the headers returned in the response from a
// Download request.
type DownloadHeader struct {
	AcceptRanges       string    `json:"Accept-Ranges"`
	ContentDisposition string    `json:"Content-Disposition"`
	ContentEncoding    string    `json:"Content-Encoding"`
	ContentLength      int64     `json:"Content-Length,string"`
	ContentType        string    `json:"Content-Type"`
	Date               time.Time `json:"-"`
	DeleteAt           time.Time `json:"-"`
	ETag               string    `json:"Etag"`
	LastModified       time.Time `json:"-"`
	ObjectManifest     string    `json:"X-Object-Manifest"`
	StaticLargeObject  bool      `json:"-"`
	TransID            string    `json:"X-Trans-Id"`
	ObjectVersionID    string    `json:"X-Object-Version-Id"`
}

func (r *DownloadHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// DownloadResult is a *http.Response that is returned from a call to the
// Download function.
type DownloadResult struct {
	gophercloud.HeaderResult
	Body io.ReadCloser
}

// Extract will return a struct of headers returned from a call to Download.
func (r DownloadResult) Extract() (*DownloadHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractContent is a function that takes a DownloadResult's io.Reader body
// and reads all available data into a slice of bytes. Please be aware that due
// the nature of io.Reader is forward-only - meaning that it can only be read
// once and not rewound. You can recreate a reader from the output of this
// function by using bytes.NewReader(downloadBytes)
func (r *DownloadResult) ExtractContent() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetHeader represents the headers returned in the response from a Get request.
type GetHeader struct {
	ContentDisposition string    `json:"Content-Disposition"`
	ContentEncoding    string    `json:"Content-Encoding"`
	ContentLength      int64     `json:"Content-Length,string"`
	ContentType        string    `json:"Content-Type"`
	Date               time.Time `json:"-"`
	DeleteAt           time.Time `json:"-"`
	ETag               string    `json:"Etag"`
	LastModified       time.Time `json:"-"`
	ObjectManifest     string    `json:"X-Object-Manifest"`
	StaticLargeObject  bool      `json:"-"`
	TransID            string    `json:"X-Trans-Id"`
	ObjectVersionID    string    `json:"X-Object-Version-Id"`
}

func (r *GetHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// GetResult is a *http.Response that is returned from a call to the Get
// function.
type GetResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Get.
func (r GetResult) Extract() (*GetHeader, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractMetadata is a function that takes a GetResult (of type *http.Response)
// and returns the custom metadata associated with the object.
func (r GetResult) ExtractMetadata() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateHeader represents the headers returned in the response from a
// Create request.
type CreateHeader struct {
	ContentLength   int64     `json:"Content-Length,string"`
	ContentType     string    `json:"Content-Type"`
	Date            time.Time `json:"-"`
	ETag            string    `json:"Etag"`
	LastModified    time.Time `json:"-"`
	TransID         string    `json:"X-Trans-Id"`
	ObjectVersionID string    `json:"X-Object-Version-Id"`
}

func (r *CreateHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// CreateResult represents the result of a create operation.
type CreateResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Create.
func (r CreateResult) Extract() (*CreateHeader, error) { _ = "STUB: not implemented"; return nil, nil }

// UpdateHeader represents the headers returned in the response from a
// Update request.
type UpdateHeader struct {
	ContentLength   int64     `json:"Content-Length,string"`
	ContentType     string    `json:"Content-Type"`
	Date            time.Time `json:"-"`
	TransID         string    `json:"X-Trans-Id"`
	ObjectVersionID string    `json:"X-Object-Version-Id"`
}

func (r *UpdateHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UpdateResult represents the result of an update operation.
type UpdateResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Update.
func (r UpdateResult) Extract() (*UpdateHeader, error) { _ = "STUB: not implemented"; return nil, nil }

// DeleteHeader represents the headers returned in the response from a
// Delete request.
type DeleteHeader struct {
	ContentLength          int64     `json:"Content-Length,string"`
	ContentType            string    `json:"Content-Type"`
	Date                   time.Time `json:"-"`
	TransID                string    `json:"X-Trans-Id"`
	ObjectVersionID        string    `json:"X-Object-Version-Id"`
	ObjectCurrentVersionID string    `json:"X-Object-Current-Version-Id"`
}

func (r *DeleteHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Delete.
func (r DeleteResult) Extract() (*DeleteHeader, error) { _ = "STUB: not implemented"; return nil, nil }

// CopyHeader represents the headers returned in the response from a
// Copy request.
type CopyHeader struct {
	ContentLength          int64     `json:"Content-Length,string"`
	ContentType            string    `json:"Content-Type"`
	CopiedFrom             string    `json:"X-Copied-From"`
	CopiedFromLastModified time.Time `json:"-"`
	Date                   time.Time `json:"-"`
	ETag                   string    `json:"Etag"`
	LastModified           time.Time `json:"-"`
	TransID                string    `json:"X-Trans-Id"`
	ObjectVersionID        string    `json:"X-Object-Version-Id"`
}

func (r *CopyHeader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// CopyResult represents the result of a copy operation.
type CopyResult struct {
	gophercloud.HeaderResult
}

// Extract will return a struct of headers returned from a call to Copy.
func (r CopyResult) Extract() (*CopyHeader, error) { _ = "STUB: not implemented"; return nil, nil }

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

// extractLastMarker is a function that takes a page of objects and returns the
// marker for the page. This can either be a subdir or the last object's name.
func extractLastMarker(r pagination.Page) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// If a delimiter was requested, check if a subdir exists.
		nil
}
