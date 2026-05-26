package objects

import (
	"context"
	"io"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ErrTempURLKeyNotFound is an error indicating that the Temp URL key was
// neigther set nor resolved from a container or account metadata.
type ErrTempURLKeyNotFound struct{ gophercloud.ErrMissingInput }

func (e ErrTempURLKeyNotFound) Error() string { _ = "STUB: not implemented"; return "" }

// ErrTempURLDigestNotValid is an error indicating that the requested
// cryptographic hash function is not supported.
type ErrTempURLDigestNotValid struct {
	gophercloud.ErrMissingInput
	Digest string
}

func (e ErrTempURLDigestNotValid) Error() string { _ = "STUB: not implemented"; return "" }

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToObjectListParams() (string, error)
}

// ListOpts is a structure that holds parameters for listing objects.
type ListOpts struct {
	// Full has been removed from the Gophercloud API. Gophercloud will now
	// always request the "full" (json) listing, because simplified listing
	// (plaintext) returns false results when names contain end-of-line
	// characters.

	Limit     int    `q:"limit"`
	Marker    string `q:"marker"`
	EndMarker string `q:"end_marker"`
	Format    string `q:"format"`
	Prefix    string `q:"prefix"`
	Delimiter string `q:"delimiter"`
	Path      string `q:"path"`
	Versions  bool   `q:"versions"`
}

// ToObjectListParams formats a ListOpts into a query string.
func (opts ListOpts) ToObjectListParams() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List is a function that retrieves all objects in a container. It also returns
// the details for the container. To extract only the object information or names,
// pass the ListResult response to the ExtractInfo or ExtractNames function,
// respectively.
func List(c *gophercloud.ServiceClient, containerName string, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// DownloadOptsBuilder allows extensions to add additional parameters to the
// Download request.
type DownloadOptsBuilder interface {
	ToObjectDownloadParams() (map[string]string, string, error)
}

// DownloadOpts is a structure that holds parameters for downloading an object.
type DownloadOpts struct {
	IfMatch           string    `h:"If-Match"`
	IfModifiedSince   time.Time `h:"If-Modified-Since"`
	IfNoneMatch       string    `h:"If-None-Match"`
	IfUnmodifiedSince time.Time `h:"If-Unmodified-Since"`
	Newest            bool      `h:"X-Newest"`
	Range             string    `h:"Range"`
	Expires           string    `q:"expires"`
	MultipartManifest string    `q:"multipart-manifest"`
	Signature         string    `q:"signature"`
	ObjectVersionID   string    `q:"version-id"`
}

// ToObjectDownloadParams formats a DownloadOpts into a query string and map of
// headers.
func (opts DownloadOpts) ToObjectDownloadParams() (map[string]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Download is a function that retrieves the content and metadata for an object.
// To extract just the content, call the DownloadResult method ExtractContent,
// after checking DownloadResult's Err field.
func Download(ctx context.Context, c *gophercloud.ServiceClient, containerName, objectName string, opts DownloadOptsBuilder) (r DownloadResult) {
	_ = "STUB: not implemented"
	return *new(DownloadResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToObjectCreateParams() (io.Reader, map[string]string, string, error)
}

// CreateOpts is a structure that holds parameters for creating an object.
type CreateOpts struct {
	Content            io.Reader
	Metadata           map[string]string
	NoETag             bool
	CacheControl       string `h:"Cache-Control"`
	ContentDisposition string `h:"Content-Disposition"`
	ContentEncoding    string `h:"Content-Encoding"`
	ContentLength      int64  `h:"Content-Length"`
	ContentType        string `h:"Content-Type"`
	CopyFrom           string `h:"X-Copy-From"`
	DeleteAfter        int64  `h:"X-Delete-After"`
	DeleteAt           int64  `h:"X-Delete-At"`
	DetectContentType  string `h:"X-Detect-Content-Type"`
	ETag               string `h:"ETag"`
	IfNoneMatch        string `h:"If-None-Match"`
	ObjectManifest     string `h:"X-Object-Manifest"`
	TransferEncoding   string `h:"Transfer-Encoding"`
	Expires            string `q:"expires"`
	MultipartManifest  string `q:"multipart-manifest"`
	Signature          string `q:"signature"`
}

// ToObjectCreateParams formats a CreateOpts into a query string and map of
// headers.
func (opts CreateOpts) ToObjectCreateParams() (io.Reader, map[string]string, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil, "", nil
}

// When we're dealing with big files an io.ReadSeeker allows us to efficiently calculate
// the md5 sum. An io.Reader is only readable once which means we have to copy the entire
// file content into memory first.

// io.Copy into md5 is very efficient as it's done in small chunks.

// Create is a function that creates a new object or replaces an existing
// object.
func Create(ctx context.Context, c *gophercloud.ServiceClient, containerName, objectName string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// CopyOptsBuilder allows extensions to add additional parameters to the
// Copy request.
type CopyOptsBuilder interface {
	ToObjectCopyMap() (map[string]string, error)
}

// CopyOptsQueryBuilder allows extensions to add additional query parameters to
// the Copy request.
type CopyOptsQueryBuilder interface {
	ToObjectCopyQuery() (string, error)
}

// CopyOpts is a structure that holds parameters for copying one object to
// another.
type CopyOpts struct {
	Metadata           map[string]string
	ContentDisposition string `h:"Content-Disposition"`
	ContentEncoding    string `h:"Content-Encoding"`
	ContentType        string `h:"Content-Type"`

	// Destination is where the object should be copied to, in the form:
	// `/container/object`.
	Destination string `h:"Destination" required:"true"`

	ObjectVersionID string `q:"version-id"`
}

// ToObjectCopyMap formats a CopyOpts into a map of headers.
func (opts CopyOpts) ToObjectCopyMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToObjectCopyQuery formats a CopyOpts into a query.
func (opts CopyOpts) ToObjectCopyQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Copy is a function that copies one object to another.
func Copy(ctx context.Context, c *gophercloud.ServiceClient, containerName, objectName string, opts CopyOptsBuilder) (r CopyResult) {
	_ = "STUB: not implemented"
	return *new(CopyResult)
}

// URL-encode the container name and the object name
// separately before joining them around the `/` slash
// separator. Note that the destination path is also
// expected to start with a slash.

// DeleteOptsBuilder allows extensions to add additional parameters to the
// Delete request.
type DeleteOptsBuilder interface {
	ToObjectDeleteQuery() (string, error)
}

// DeleteOpts is a structure that holds parameters for deleting an object.
type DeleteOpts struct {
	MultipartManifest string `q:"multipart-manifest"`
	ObjectVersionID   string `q:"version-id"`
}

// ToObjectDeleteQuery formats a DeleteOpts into a query string.
func (opts DeleteOpts) ToObjectDeleteQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Delete is a function that deletes an object.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, containerName, objectName string, opts DeleteOptsBuilder) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// GetOptsBuilder allows extensions to add additional parameters to the
// Get request.
type GetOptsBuilder interface {
	ToObjectGetParams() (map[string]string, string, error)
}

// GetOpts is a structure that holds parameters for getting an object's
// metadata.
type GetOpts struct {
	Newest          bool   `h:"X-Newest"`
	Expires         string `q:"expires"`
	Signature       string `q:"signature"`
	ObjectVersionID string `q:"version-id"`
}

// ToObjectGetParams formats a GetOpts into a query string and a map of headers.
func (opts GetOpts) ToObjectGetParams() (map[string]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Get is a function that retrieves the metadata of an object. To extract just
// the custom metadata, pass the GetResult response to the ExtractMetadata
// function.
func Get(ctx context.Context, c *gophercloud.ServiceClient, containerName, objectName string, opts GetOptsBuilder) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToObjectUpdateMap() (map[string]string, error)
}

// UpdateOpts is a structure that holds parameters for updating, creating, or
// deleting an object's metadata.
type UpdateOpts struct {
	Metadata           map[string]string
	RemoveMetadata     []string
	ContentDisposition *string `h:"Content-Disposition"`
	ContentEncoding    *string `h:"Content-Encoding"`
	ContentType        *string `h:"Content-Type"`
	DeleteAfter        *int64  `h:"X-Delete-After"`
	DeleteAt           *int64  `h:"X-Delete-At"`
	DetectContentType  *bool   `h:"X-Detect-Content-Type"`
}

// ToObjectUpdateMap formats a UpdateOpts into a map of headers.
func (opts UpdateOpts) ToObjectUpdateMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update is a function that creates, updates, or deletes an object's metadata.
func Update(ctx context.Context, c *gophercloud.ServiceClient, containerName, objectName string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// HTTPMethod represents an HTTP method string (e.g. "GET").
type HTTPMethod string

var (
	// GET represents an HTTP "GET" method.
	GET HTTPMethod = "GET"
	// HEAD represents an HTTP "HEAD" method.
	HEAD HTTPMethod = "HEAD"
	// PUT represents an HTTP "PUT" method.
	PUT HTTPMethod = "PUT"
	// POST represents an HTTP "POST" method.
	POST HTTPMethod = "POST"
	// DELETE represents an HTTP "DELETE" method.
	DELETE HTTPMethod = "DELETE"
)

// CreateTempURLOpts are options for creating a temporary URL for an object.
type CreateTempURLOpts struct {
	// (REQUIRED) Method is the HTTP method to allow for users of the temp URL.
	// Valid values are "GET", "HEAD", "PUT", "POST" and "DELETE".
	Method HTTPMethod

	// (REQUIRED) TTL is the number of seconds the temp URL should be active.
	TTL int

	// (Optional) Split is the string on which to split the object URL. Since only
	// the object path is used in the hash, the object URL needs to be parsed. If
	// empty, the default OpenStack URL split point will be used ("/v1/").
	Split string

	// (Optional) Timestamp is the current timestamp used to calculate the Temp URL
	// signature. If not specified, the current UNIX timestamp is used as the base
	// timestamp.
	Timestamp time.Time

	// (Optional) TempURLKey overrides the Swift container or account Temp URL key.
	// TempURLKey must correspond to a target container/account key, otherwise the
	// generated link will be invalid. If not specified, the key is obtained from
	// a Swift container or account.
	TempURLKey string

	// (Optional) Digest specifies the cryptographic hash function used to
	// calculate the signature. Valid values include sha1, sha256, and
	// sha512. If not specified, the default hash function is sha1.
	Digest string
}

// CreateTempURL is a function for creating a temporary URL for an object. It
// allows users to have "GET" or "POST" access to a particular tenant's object
// for a limited amount of time.
func CreateTempURL(ctx context.Context, c *gophercloud.ServiceClient, containerName, objectName string, opts CreateTempURLOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Initialize time if it was not passed as opts

// UNIX time is always UTC

// Initialize the tempURLKey to calculate a signature

// fallback to a container TempURL key

// fallback to an account TempURL key

// BulkDelete is a function that bulk deletes objects.
// In Swift, the maximum number of deletes per request is set by default to 10000.
//
// See:
// * https://github.com/openstack/swift/blob/6d3d4197151f44bf28b51257c1a4c5d33411dcae/etc/proxy-server.conf-sample#L1029-L1034
// * https://github.com/openstack/swift/blob/e8cecf7fcc1630ee83b08f9a73e1e59c07f8d372/swift/common/middleware/bulk.py#L309
func BulkDelete(ctx context.Context, c *gophercloud.ServiceClient, container string, objects []string) (r BulkDeleteResult) {
	_ = "STUB: not implemented"
	return *new(BulkDeleteResult)
}
