package transfers

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Transfer represents a Volume Transfer record
type Transfer struct {
	ID        string              `json:"id"`
	AuthKey   string              `json:"auth_key"`
	Name      string              `json:"name"`
	VolumeID  string              `json:"volume_id"`
	CreatedAt time.Time           `json:"-"`
	Links     []map[string]string `json:"links"`
}

// UnmarshalJSON is our unmarshalling helper
func (r *Transfer) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Transfer object out of the commonResult object.
func (r commonResult) Extract() (*Transfer, error) { _ = "STUB: not implemented"; return nil, nil }

// ExtractInto converts our response data into a transfer struct
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

// ExtractTransfers extracts and returns Transfers. It is used while iterating over a transfers.List call.
func ExtractTransfers(r pagination.Page) ([]Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractTransfersInto similar to ExtractInto but operates on a `list` of transfers
func ExtractTransfersInto(r pagination.Page, v any) error { _ = "STUB: not implemented"; return nil }

// TransferPage is a pagination.pager that is returned from a call to the List function.
type TransferPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no Transfers.
func (r TransferPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (page TransferPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
