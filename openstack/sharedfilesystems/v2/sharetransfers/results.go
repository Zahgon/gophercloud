package sharetransfers

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

const (
	invalidMarker = "-1"
)

// Transfer represents a Share Transfer record.
type Transfer struct {
	ID                   string              `json:"id"`
	Accepted             bool                `json:"accepted"`
	AuthKey              string              `json:"auth_key"`
	Name                 string              `json:"name"`
	SourceProjectID      string              `json:"source_project_id"`
	DestinationProjectID string              `json:"destination_project_id"`
	ResourceID           string              `json:"resource_id"`
	ResourceType         string              `json:"resource_type"`
	CreatedAt            time.Time           `json:"-"`
	ExpiresAt            time.Time           `json:"-"`
	Links                []map[string]string `json:"links"`
}

// UnmarshalJSON is our unmarshalling helper.
func (r *Transfer) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Transfer object out of the commonResult object.
func (r commonResult) Extract() (*Transfer, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractInto converts our response data into a transfer struct.
func (r commonResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

// AcceptResult contains the response body and error from an Accept request.
type AcceptResult struct {
	gophercloud.ErrResult
}

// ExtractTransfers extracts and returns Transfers. It is used while iterating over a transfers.List call.
func ExtractTransfers(r pagination.Page) ([]Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractTransfersInto similar to ExtractInto but operates on a `list` of transfers
func ExtractTransfersInto(r pagination.Page, v any) error { _ = "STUB: not implemented"; return nil }

// TransferPage is a pagination.pager that is returned from a call to the List function.
type TransferPage struct {
	pagination.MarkerPageBase
}

// NextPageURL generates the URL for the page of results after this one.
func (r TransferPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LastMarker returns the last offset in a ListResult.
func (r TransferPage) LastMarker() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Limit is not present, only one page required

// IsEmpty satisifies the IsEmpty method of the Page interface.
func (r TransferPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }
