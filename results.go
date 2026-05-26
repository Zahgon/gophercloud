package gophercloud

import (
	"net/http"
	"time"
)

/*
Result is an internal type to be used by individual resource packages, but its
methods will be available on a wide variety of user-facing embedding types.

It acts as a base struct that other Result types, returned from request
functions, can embed for convenience. All Results capture basic information
from the HTTP transaction that was performed, including the response body,
HTTP headers, and any errors that happened.

Generally, each Result type will have an Extract method that can be used to
further interpret the result's payload in a specific context. Extensions or
providers can then provide additional extraction functions to pull out
provider- or extension-specific information as well.
*/
type Result struct {
	// Body is the payload of the HTTP response from the server. In most cases,
	// this will be the deserialized JSON structure.
	Body any

	// StatusCode is the HTTP status code of the original response. Will be
	// one of the OkCodes defined on the gophercloud.RequestOpts that was
	// used in the request.
	StatusCode int

	// Header contains the HTTP header structure from the original response.
	Header http.Header

	// Err is an error that occurred during the operation. It's deferred until
	// extraction to make it easier to chain the Extract call.
	Err error
}

// ExtractInto allows users to provide an object into which `Extract` will extract
// the `Result.Body`. This would be useful for OpenStack providers that have
// different fields in the response object than OpenStack proper.
func (r Result) ExtractInto(to any) error { _ = "STUB: not implemented"; return nil }

func (r Result) extractIntoPtr(to any, label string) error { _ = "STUB: not implemented"; return nil }

// For each iteration of the slice, we create a new struct.
// This is to work around a bug where elements of a slice
// are reused and not overwritten when the same copy of the
// struct is used:
//
// https://github.com/golang/go/issues/21092
// https://github.com/golang/go/issues/24155
// https://play.golang.org/p/NHo3ywlPZli

// This is needed for structs with an UnmarshalJSON method.
// Technically this is just unmarshalling the response into
// a struct that is never used, but it's good enough to
// trigger the UnmarshalJSON method.

// Unmarshal is used rather than NewDecoder to also work
// around the above-mentioned bug.

// "to" should now be properly modeled to receive the
// JSON response body and unmarshal into all the correct
// fields of the struct or composed extension struct
// at the end of this method.

// jtopjian: This was put into place to resolve the issue
// described at
// https://github.com/gophercloud/gophercloud/issues/1963
//
// This probably isn't the best fix, but it appears to
// be resolving the issue, so I'm going to implement it
// for now.
//
// For future readers, this entire case statement could
// use a review.

// ExtractIntoStructPtr will unmarshal the Result (r) into the provided
// any (to).
//
// NOTE: For internal use only
//
// `to` must be a pointer to an underlying struct type
//
// If provided, `label` will be filtered out of the response
// body prior to `r` being unmarshalled into `to`.
func (r Result) ExtractIntoStructPtr(to any, label string) error {
	_ = "STUB: not implemented"
	return nil
}

// ExtractIntoSlicePtr will unmarshal the Result (r) into the provided
// any (to).
//
// NOTE: For internal use only
//
// `to` must be a pointer to an underlying slice type
//
// If provided, `label` will be filtered out of the response
// body prior to `r` being unmarshalled into `to`.
func (r Result) ExtractIntoSlicePtr(to any, label string) error {
	_ = "STUB: not implemented"
	return nil
}

// PrettyPrintJSON creates a string containing the full response body as
// pretty-printed JSON. It's useful for capturing test fixtures and for
// debugging extraction bugs. If you include its output in an issue related to
// a buggy extraction function, we will all love you forever.
func (r Result) PrettyPrintJSON() string { _ = "STUB: not implemented"; return "" }

// ErrResult is an internal type to be used by individual resource packages, but
// its methods will be available on a wide variety of user-facing embedding
// types.
//
// It represents results that only contain a potential error and
// nothing else. Usually, if the operation executed successfully, the Err field
// will be nil; otherwise it will be stocked with a relevant error. Use the
// ExtractErr method
// to cleanly pull it out.
type ErrResult struct {
	Result
}

// ExtractErr is a function that extracts error information, or nil, from a result.
func (r ErrResult) ExtractErr() error {
	_ = "STUB: not implemented"

	/*
	   HeaderResult is an internal type to be used by individual resource packages, but
	   its methods will be available on a wide variety of user-facing embedding types.

	   It represents a result that only contains an error (possibly nil) and an
	   http.Header. This is used, for example, by the objectstorage packages in
	   openstack, because most of the operations don't return response bodies, but do
	   have relevant information in headers.
	*/return nil
}

type HeaderResult struct {
	Result
}

// ExtractInto allows users to provide an object into which `Extract` will
// extract the http.Header headers of the result.
func (r HeaderResult) ExtractInto(to any) error { _ = "STUB: not implemented"; return nil }

// RFC3339Milli describes a common time format used by some API responses.
const RFC3339Milli = "2006-01-02T15:04:05.999999Z"

type JSONRFC3339Milli time.Time

func (jt *JSONRFC3339Milli) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

const RFC3339MilliNoZ = "2006-01-02T15:04:05.999999"

type JSONRFC3339MilliNoZ time.Time

func (jt *JSONRFC3339MilliNoZ) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type JSONRFC1123 time.Time

func (jt *JSONRFC1123) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type JSONUnix time.Time

func (jt *JSONUnix) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// RFC3339NoZ is the time format used in Heat (Orchestration).
const RFC3339NoZ = "2006-01-02T15:04:05"

type JSONRFC3339NoZ time.Time

func (jt *JSONRFC3339NoZ) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// RFC3339ZNoT is the time format used in Zun (Containers Service).
const RFC3339ZNoT = "2006-01-02 15:04:05-07:00"

type JSONRFC3339ZNoT time.Time

func (jt *JSONRFC3339ZNoT) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// RFC3339ZNoTNoZ is another time format used in Zun (Containers Service).
const RFC3339ZNoTNoZ = "2006-01-02 15:04:05"

type JSONRFC3339ZNoTNoZ time.Time

func (jt *JSONRFC3339ZNoTNoZ) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

/*
Link is an internal type to be used in packages of collection resources that are
paginated in a certain way.

It's a response substructure common to many paginated collection results that is
used to point to related pages. Usually, the one we care about is the one with
Rel field set to "next".
*/
type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

/*
ExtractNextURL is an internal function useful for packages of collection
resources that are paginated in a certain way.

It attempts to extract the "next" URL from slice of Link structs, or
"" if no such URL is present.
*/
func ExtractNextURL(links []Link) (string, error) { _ = "STUB: not implemented"; return "", nil }
